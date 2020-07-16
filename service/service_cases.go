package service

import (
	"bcbtest/dao"
	"fmt"
)

type CaseInfo struct {
	Id   int64  `json:"id"`
	Desc string `json:"desc"`
	Exp  string `json:"exp"`
}

type TestSuite struct {
	FirstSuite  string     `json:"first_suite"`
	SecondSuite string     `json:"second_suite"`
	ThirdSuite  string     `json:"third_suite"`
	Cases       []CaseInfo `json:"cases"`
}

type TestKind struct {
	Name       string      `json:"kind_name"`
	TestSuites []TestSuite `json:"test_suites"`
}

func ExecBCBCase() {
	// 将之前的测试报告信息迁到历史表
	dao.MigrateCaseReport()

	// 情况测试用例报告表
	dao.TrunkCaseReport()

	db := dao.GetClient()

	// 新增测试报告
	reportID := dao.GetReportID() + 1
	db.Create(&dao.Report{
		ReportID:     reportID,
		NodeVersion:  "CentOS-7.4",
		ChainVersion: "2.2.1.19452",
		IsRunning:    1,
	})

	for _, caseInfo := range dao.GetAllCase() {
		// TODO 执行用例

		// 生成用例测试报告
		db.Create(&dao.ReportCase{
			CaseID:          caseInfo.CaseID,
			ReportID:        reportID,
			KindID:          caseInfo.KindID,
			SuiteID:         caseInfo.SuiteID,
			Result:          "成功",
			IsExpect:        0,
			CheckItemResult: "检查项1:0;检查项2:0;检查项3:0;检查项4:0;",
		})
	}
}

func GetTestKinds() []TestKind {
	var sTestKinds []TestKind
	for _, testKind := range dao.GetTestKinds() {
		sTestKind := TestKind{
			Name: testKind.KindName,
		}
		var sTestSuites []TestSuite
		for _, testSuite := range dao.GetTestSuites(testKind.KindID) {
			sTestSuite := TestSuite{
				FirstSuite:  testSuite.FirstSuite,
				SecondSuite: testSuite.SecondSuite,
				ThirdSuite:  testSuite.ThirdSuite,
			}
			var sCaseInfos []CaseInfo
			for _, testCase := range dao.GetSuiteCases(testSuite.SuiteID) {
				fmt.Println(testCase)
				sCaseInfo := CaseInfo{
					Id:   testCase.CaseID,
					Desc: testCase.Desc,
					Exp:  testCase.ExpectDesc,
				}
				sCaseInfos = append(sCaseInfos, sCaseInfo)
			}
			sTestSuite.Cases = sCaseInfos
			sTestSuites = append(sTestSuites, sTestSuite)
		}
		sTestKind.TestSuites = sTestSuites
		sTestKinds = append(sTestKinds, sTestKind)
	}
	return sTestKinds
}
