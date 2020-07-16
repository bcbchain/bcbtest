package dao

type TestCase struct {
	Id           int
	CaseID       int64
	KindID       int64
	SuiteID      int64
	Desc         string
	RepeatCount  int64
	ExpectDesc   string
	ExpectedCode int64
}

func (m *TestCase) TableName() string {
	return "test_case"
}

func TrunkTestCase() {
	db.Exec("TRUNCATE TABLE test_case;")
}

func GetSuiteCases(suiteID int64) []*TestCase {
	testCases := make([]*TestCase, 0)
	db.Where("suite_id = ?", suiteID).Order("case_id asc").Find(&testCases)
	return testCases
}

func GetCase(caseID int64) *TestCase {
	testCase := new(TestCase)
	db.Where("case_id = ?", caseID).Find(&testCase)
	return testCase
}

func GetAllCase() []*TestCase {
	testCases := make([]*TestCase, 0)
	db.Order("case_id asc").Find(&testCases)
	return testCases
}
