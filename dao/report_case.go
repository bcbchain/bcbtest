package dao

type ReportCase struct {
	Id              int
	CaseID          int64
	ReportID        int64
	KindID          int64
	SuiteID         int64
	Result          string
	IsExpect        uint8
	CheckItemResult string
}

func (m *ReportCase) TableName() string {
	return "report_case"
}

type ReportCaseHis struct {
	Id              int
	CaseID          int64
	ReportID        int64
	KindID          int64
	SuiteID         int64
	Result          string
	IsExpect        uint8
	CheckItemResult string
}

func (m *ReportCaseHis) TableName() string {
	return "report_case_his"
}

func TrunkCaseReport() {
	db.Exec("TRUNCATE TABLE report_case;")
}

func GetCaseReports() []*ReportCase {
	reportCases := make([]*ReportCase, 0)
	db.Order("case_id asc").Find(&reportCases)
	return reportCases
}

func GetCaseReport(caseID int64) *ReportCase {
	reportCase := new(ReportCase)
	db.Where("case_id = ?", caseID).Find(&reportCase)
	return reportCase
}

func GetSuiteReports(suiteID int64) []*ReportCase {
	reportCases := make([]*ReportCase, 0)
	db.Where("suite_id = ?", suiteID).Order("case_id asc").Find(&reportCases)
	return reportCases
}

func MigrateCaseReport() {
	for _, caseReport := range GetCaseReports() {
		db.Create(&ReportCaseHis{
			CaseID:          caseReport.CaseID,
			ReportID:        caseReport.ReportID,
			KindID:          caseReport.KindID,
			SuiteID:         caseReport.SuiteID,
			Result:          caseReport.Result,
			IsExpect:        caseReport.IsExpect,
			CheckItemResult: caseReport.CheckItemResult,
		})
	}
}
