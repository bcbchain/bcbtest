package service

import "bcbtest/dao"

type SingleReportCase struct {
	Id          int64
	KindName    string
	FirstSuite  string
	SecondSuite string
	ThirdSuite  string
	Desc        string
	Exp         string
	Result      string
	CheckItems  CheckItems
}

func GetReport(caseID int64) SingleReportCase {
	caseReport := dao.GetCaseReport(caseID)
	ts := dao.GetTestSuite(caseReport.SuiteID)
	tc := dao.GetCase(caseID)
	singleReportCase := SingleReportCase{
		Id:          caseReport.CaseID,
		KindName:    dao.GetTestKind(caseReport.KindID).KindName,
		FirstSuite:  ts.FirstSuite,
		SecondSuite: ts.SecondSuite,
		ThirdSuite:  ts.ThirdSuite,
		Desc:        tc.Desc,
		Exp:         tc.ExpectDesc,
		Result:      caseReport.Result,
		CheckItems:  *str2CheckItem(caseReport.CheckItemResult),
	}
	return singleReportCase
}
