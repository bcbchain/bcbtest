package dao

type TestSuite struct {
	Id          int //对应数据表的自增id
	SuiteID     int64
	KindID      int64
	FirstSuite  string
	SecondSuite string
	ThirdSuite  string
}

func (m *TestSuite) TableName() string {
	return "test_suite"
}

func TrunkTestSuite() {
	db.Exec("TRUNCATE TABLE test_suite;")
}

func GetTestSuites(kindID int64) []*TestSuite {
	testSuites := make([]*TestSuite, 0)
	db.Where("kind_id = ?", kindID).Order("suite_id asc").Find(&testSuites)
	return testSuites
}

func GetTestSuite(suiteID int64) *TestSuite {
	testSuite := new(TestSuite)
	db.Where("suite_id = ?", suiteID).Find(&testSuite)
	return testSuite
}

func GetFirstSuiteNames(kindID int64) []string {
	var firstSuiteNames []string
	testSuites := make([]*TestSuite, 0)
	db.Select("distinct first_suite").Where("kind_id = ?", kindID).Order("suite_id asc").Find(&testSuites)
	for _, testSuite := range testSuites {
		firstSuiteNames = append(firstSuiteNames, testSuite.FirstSuite)
	}
	return firstSuiteNames
}

func GetSecondSuiteNames(firstSuiteName string) []string {
	var SecondSuiteNames []string
	testSuites := make([]*TestSuite, 0)
	db.Select("distinct second_suite").Where("first_suite = ?", firstSuiteName).Order("suite_id asc").Find(&testSuites)
	for _, testSuite := range testSuites {
		SecondSuiteNames = append(SecondSuiteNames, testSuite.SecondSuite)
	}
	return SecondSuiteNames
}

func GetThirdSuiteNames(secondSuiteName string) []*TestSuite {
	testSuites := make([]*TestSuite, 0)
	db.Select("suite_id, third_suite").Where("second_suite = ?", secondSuiteName).Order("suite_id asc").Find(&testSuites)
	return testSuites
}
