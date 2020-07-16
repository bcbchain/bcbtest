package service

import "bcbtest/dao"

type ReportCase struct {
	Id     int64  `json:"id"`
	Desc   string `json:"desc"`
	Exp    string `json:"exp"`
	Result string `json:"result"`
}

//
//type ReportTestSuite struct {
//	FirstSuite string `json:"first_suite"`
//	SecondSuite string `json:"second_suite"`
//	ThirdSuite string `json:"third_suite"`
//	Cases []ReportCase `json:"cases"`
//}
//
//type ReportTestKind struct {
//	Name string `json:"kind_name"`
//	TestSuites []ReportTestSuite `json:"test_suites"`
//}

type ReportThirdTestSuite struct {
	Name  string        `json:"third_suite_name"`
	Cases []*ReportCase `json:"third_suite_data"`
}

type ReportSecondTestSuite struct {
	Name            string                  `json:"second_suite_name"`
	ThirdTestSuites []*ReportThirdTestSuite `json:"second_suite_data"`
}

type ReportFirstTestSuite struct {
	Name             string                   `json:"first_suite_name"`
	SecondTestSuites []*ReportSecondTestSuite `json:"first_suite_data"`
}

type ReportTestKind struct {
	Name            string                  `json:"kind_name"`
	FirstTestSuites []*ReportFirstTestSuite `json:"kind_data"`
	//TestSuite FirstTestSuite `json:"test_suite"`
}

//func GetReports() []ReportTestKind {
//	var sRTestKinds []ReportTestKind
//	for _, testKind := range dao.GetTestKinds() {
//		sRTestKind := ReportTestKind{
//			Name: testKind.KindName,
//		}
//		var sRTestSuites []ReportTestSuite
//		for _, testSuite := range dao.GetTestSuites(testKind.KindID) {
//			sRTestSuite := ReportTestSuite{
//				FirstSuite:  testSuite.FirstSuite,
//				SecondSuite: testSuite.SecondSuite,
//				ThirdSuite:  testSuite.ThirdSuite,
//			}
//			var sRCaseInfos []ReportCase
//			for _, testCase := range dao.GetSuiteReports(testSuite.SuiteID) {
//				sRCaseInfo := ReportCase{
//					Id:   testCase.CaseID,
//					Desc: dao.GetCase(testCase.CaseID).Desc,
//					Exp:  dao.GetCase(testCase.CaseID).ExpectDesc,
//					Result: dao.GetCaseReport(testCase.CaseID).Result,
//				}
//				sRCaseInfos = append(sRCaseInfos, sRCaseInfo)
//			}
//			sRTestSuite.Cases = sRCaseInfos
//			sRTestSuites = append(sRTestSuites, sRTestSuite)
//		}
//		sRTestKind.TestSuites = sRTestSuites
//		sRTestKinds = append(sRTestKinds, sRTestKind)
//	}
//	return sRTestKinds
//
//}

func GetReports() []ReportTestKind {
	var rTestKinds []ReportTestKind
	for _, testKind := range dao.GetTestKinds() {
		rTestKind := ReportTestKind{
			Name: testKind.KindName,
		}
		var rFTestSuites []*ReportFirstTestSuite
		for _, firstSuiteName := range dao.GetFirstSuiteNames(testKind.KindID) {
			var rFTestSuite = &ReportFirstTestSuite{
				Name: firstSuiteName,
			}
			for _, secondSuiteName := range dao.GetSecondSuiteNames(firstSuiteName) {
				var rSTestSuite = &ReportSecondTestSuite{
					Name: secondSuiteName,
				}
				for _, thirdTestSuite := range dao.GetThirdSuiteNames(secondSuiteName) {
					var rTTestSuite = &ReportThirdTestSuite{
						Name: thirdTestSuite.ThirdSuite,
					}
					//for _, testCase := range dao.GetSuiteCases(thirdTestSuite.SuiteID) {
					for _, testCase := range dao.GetSuiteReports(thirdTestSuite.SuiteID) {
						rCaseInfo := &ReportCase{
							Id:     testCase.CaseID,
							Desc:   dao.GetCase(testCase.CaseID).Desc,
							Exp:    dao.GetCase(testCase.CaseID).ExpectDesc,
							Result: testCase.Result,
						}
						rTTestSuite.Cases = append(rTTestSuite.Cases, rCaseInfo)
					}
					rSTestSuite.ThirdTestSuites = append(rSTestSuite.ThirdTestSuites, rTTestSuite)
				}
				rFTestSuite.SecondTestSuites = append(rFTestSuite.SecondTestSuites, rSTestSuite)
			}
			rFTestSuites = append(rFTestSuites, rFTestSuite)
			//sTestKind.FirstTestSuites = append(sTestKind.FirstTestSuites, fTestSuite)
			//sTestKinds = append(sTestKinds, sTestKind)
			rTestKind.FirstTestSuites = rFTestSuites
		}
		//sTestKind.FirstTestSuites = *fir
		rTestKinds = append(rTestKinds, rTestKind)
	}
	return rTestKinds
}
