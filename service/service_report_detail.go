package service

import (
	"bcbtest/dao"
	"strconv"
	"strings"
)

type ReportDetailCase struct {
	Id         int64      `json:"id"`
	Desc       string     `json:"desc"`
	Exp        string     `json:"exp"`
	Result     string     `json:"result"`
	CheckItems CheckItems `json:"check_item"`
}

type ReportDetailThirdTestSuite struct {
	Name  string              `json:"third_suite_name"`
	Cases []*ReportDetailCase `json:"third_suite_data"`
}

type ReportDetailSecondTestSuite struct {
	Name            string                        `json:"second_suite_name"`
	ThirdTestSuites []*ReportDetailThirdTestSuite `json:"second_suite_data"`
}

type ReportDetailFirstTestSuite struct {
	Name             string                         `json:"first_suite_name"`
	SecondTestSuites []*ReportDetailSecondTestSuite `json:"first_suite_data"`
}

type ReportDetailTestKind struct {
	Name            string                        `json:"kind_name"`
	FirstTestSuites []*ReportDetailFirstTestSuite `json:"kind_data"`
	//TestSuite FirstTestSuite `json:"test_suite"`
}

//type ReportDetailTestSuite struct {
//	FirstSuite string `json:"first_suite"`
//	SecondSuite string `json:"second_suite"`
//	ThirdSuite string `json:"third_suite"`
//	Cases []ReportDetailCase `json:"report_cases"`
//}
//
//type ReportDetailTestKind struct {
//	Name string `json:"kind_name"`
//	TestSuites []ReportDetailTestSuite `json:"test_suites"`
//}

//func GetReportsDetail() []ReportDetailTestKind {
//	var sRDTestKinds []ReportDetailTestKind
//	for _, testKind := range dao.GetTestKinds() {
//		sRDTestKind := ReportDetailTestKind{
//			Name: testKind.KindName,
//		}
//		var sRDTestSuites []ReportDetailTestSuite
//		for _, testSuite := range dao.GetTestSuites(testKind.KindID) {
//			sRDTestSuite := ReportDetailTestSuite{
//				FirstSuite:  testSuite.FirstSuite,
//				SecondSuite: testSuite.SecondSuite,
//				ThirdSuite:  testSuite.ThirdSuite,
//			}
//			var sRDCaseInfos []ReportDetailCase
//			for _, testCase := range dao.GetSuiteReports(testSuite.SuiteID) {
//
//				sRDCaseInfo := ReportDetailCase{
//					Id:   testCase.CaseID,
//					Desc: dao.GetCase(testCase.CaseID).Desc,
//					Exp:  dao.GetCase(testCase.CaseID).ExpectDesc,
//					Result: dao.GetCaseReport(testCase.CaseID).Result,
//					CheckItems: *str2CheckItem(dao.GetCaseReport(testCase.CaseID).CheckItemResult),
//				}
//				sRDCaseInfos = append(sRDCaseInfos, sRDCaseInfo)
//			}
//			sRDTestSuite.Cases = sRDCaseInfos
//			sRDTestSuites = append(sRDTestSuites, sRDTestSuite)
//		}
//		sRDTestKind.TestSuites = sRDTestSuites
//		sRDTestKinds = append(sRDTestKinds, sRDTestKind)
//	}
//	return sRDTestKinds
//}

func GetReportsDetail() []ReportDetailTestKind {
	var rDTestKinds []ReportDetailTestKind
	for _, testKind := range dao.GetTestKinds() {
		rDTestKind := ReportDetailTestKind{
			Name: testKind.KindName,
		}
		var rDFTestSuites []*ReportDetailFirstTestSuite
		for _, firstSuiteName := range dao.GetFirstSuiteNames(testKind.KindID) {
			var rDFTestSuite = &ReportDetailFirstTestSuite{
				Name: firstSuiteName,
			}
			for _, secondSuiteName := range dao.GetSecondSuiteNames(firstSuiteName) {
				var rDSTestSuite = &ReportDetailSecondTestSuite{
					Name: secondSuiteName,
				}
				for _, thirdTestSuite := range dao.GetThirdSuiteNames(secondSuiteName) {
					var rDTTestSuite = &ReportDetailThirdTestSuite{
						Name: thirdTestSuite.ThirdSuite,
					}
					//for _, testCase := range dao.GetSuiteCases(thirdTestSuite.SuiteID) {
					for _, testCase := range dao.GetSuiteReports(thirdTestSuite.SuiteID) {
						rDCaseInfo := &ReportDetailCase{
							Id:         testCase.CaseID,
							Desc:       dao.GetCase(testCase.CaseID).Desc,
							Exp:        dao.GetCase(testCase.CaseID).ExpectDesc,
							Result:     testCase.Result,
							CheckItems: *str2CheckItem(testCase.CheckItemResult),
						}
						rDTTestSuite.Cases = append(rDTTestSuite.Cases, rDCaseInfo)
					}
					rDSTestSuite.ThirdTestSuites = append(rDSTestSuite.ThirdTestSuites, rDTTestSuite)
				}
				rDFTestSuite.SecondTestSuites = append(rDFTestSuite.SecondTestSuites, rDSTestSuite)
			}
			rDFTestSuites = append(rDFTestSuites, rDFTestSuite)
			//sTestKind.FirstTestSuites = append(sTestKind.FirstTestSuites, fTestSuite)
			//sTestKinds = append(sTestKinds, sTestKind)
			rDTestKind.FirstTestSuites = rDFTestSuites
		}
		//sTestKind.FirstTestSuites = *fir
		rDTestKinds = append(rDTestKinds, rDTestKind)
	}
	return rDTestKinds
}

func str2CheckItem(strCheckItem string) *CheckItems {
	var cis = &CheckItems{}
	checkItemsKV := strings.Split(strCheckItem, ";")
	for _, checkItemKV := range checkItemsKV {
		tempKV := strings.Split(checkItemKV, ":")
		if len(tempKV) == 2 {
			ciRes, err := strconv.ParseUint(tempKV[1], 8, 8)
			if err != nil {
				ciRes = 1
			}
			var ci = &CheckItem{
				Desc:   tempKV[0],
				Result: uint8(ciRes),
			}
			*cis = append(*cis, ci)
		}
	}
	return cis
}
