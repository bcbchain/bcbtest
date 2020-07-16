package service

import (
	"bcbtest/dao"
)

type CaseInfoP struct {
	Id   int64  `json:"id"`
	Desc string `json:"desc"`
	Exp  string `json:"exp"`
}

type ThirdTestSuite struct {
	Name  string       `json:"third_suite_name"`
	Cases []*CaseInfoP `json:"third_suite_data"`
}

type SecondTestSuite struct {
	Name            string            `json:"second_suite_name"`
	ThirdTestSuites []*ThirdTestSuite `json:"second_suite_data"`
}

type FirstTestSuite struct {
	Name             string             `json:"first_suite_name"`
	SecondTestSuites []*SecondTestSuite `json:"first_suite_data"`
}

type TestSuiteP struct {
	//Name string
	FirstTestSuites []*FirstTestSuite `json:"data"`
}

type TestKindP struct {
	Name            string            `json:"kind_name"`
	FirstTestSuites []*FirstTestSuite `json:"kind_data"`
	//TestSuite FirstTestSuite `json:"test_suite"`
}

func GetCases() []TestKindP {
	var sTestKinds []TestKindP
	for _, testKind := range dao.GetTestKinds() {
		sTestKind := TestKindP{
			Name: testKind.KindName,
		}
		var fTestSuites []*FirstTestSuite
		for _, firstSuiteName := range dao.GetFirstSuiteNames(testKind.KindID) {
			var fTestSuite = &FirstTestSuite{
				Name: firstSuiteName,
			}
			for _, secondSuiteName := range dao.GetSecondSuiteNames(firstSuiteName) {
				var sTestSuite = &SecondTestSuite{
					Name: secondSuiteName,
				}
				for _, thirdTestSuite := range dao.GetThirdSuiteNames(secondSuiteName) {
					var tTestSuite = &ThirdTestSuite{
						Name: thirdTestSuite.ThirdSuite,
					}
					for _, testCase := range dao.GetSuiteCases(thirdTestSuite.SuiteID) {
						caseInfo := &CaseInfoP{
							Id:   testCase.CaseID,
							Desc: testCase.Desc,
							Exp:  testCase.ExpectDesc,
						}
						tTestSuite.Cases = append(tTestSuite.Cases, caseInfo)
					}
					sTestSuite.ThirdTestSuites = append(sTestSuite.ThirdTestSuites, tTestSuite)
				}
				fTestSuite.SecondTestSuites = append(fTestSuite.SecondTestSuites, sTestSuite)
			}
			fTestSuites = append(fTestSuites, fTestSuite)
			//sTestKind.FirstTestSuites = append(sTestKind.FirstTestSuites, fTestSuite)
			//sTestKinds = append(sTestKinds, sTestKind)
			sTestKind.FirstTestSuites = fTestSuites
		}
		//sTestKind.FirstTestSuites = *fir
		sTestKinds = append(sTestKinds, sTestKind)
	}
	return sTestKinds
}
