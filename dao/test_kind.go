package dao

import (
	"bcbtest/parse"
)

type TestCaseKind struct {
	Id       int
	KindID   int64
	KindName string
}

func (m *TestCaseKind) TableName() string {
	return "test_case_kind"
}

func TrunkTestKind() {
	db.Exec("TRUNCATE TABLE test_case_kind;")
}

func LoadTestKinds(testKinds parse.TestKinds) {
	TrunkTestKind()
	TrunkTestSuite()
	TrunkTestCase()

	for _, testKind := range testKinds {
		db.Create(&TestCaseKind{
			KindID:   testKind.KindID,
			KindName: testKind.KindName,
		})
		for _, firstV := range testKind.Suites.OrderKey {
			for _, secondV := range testKind.Suites.Value[firstV].OrderKey {
				for _, thirdV := range testKind.Suites.Value[firstV].Value[secondV].OrderKey {
					suiteID := GetSuiteIDSeq()
					db.Create(&TestSuite{
						SuiteID:     suiteID,
						KindID:      testKind.KindID,
						FirstSuite:  firstV,
						SecondSuite: secondV,
						ThirdSuite:  thirdV,
					})
					for _, cases := range testKind.Suites.Value[firstV].Value[secondV].Value[thirdV] {
						db.Create(&TestCase{
							CaseID:       GetCaseIDSeq(),
							KindID:       testKind.KindID,
							SuiteID:      suiteID,
							Desc:         cases.Desc,
							RepeatCount:  0,
							ExpectedCode: 0,
							ExpectDesc:   cases.Exp,
						})
					}
				}
			}
		}
	}
}

func GetTestKinds() []*TestCaseKind {
	testCaseKinds := make([]*TestCaseKind, 0)
	db.Order("kind_id asc").Find(&testCaseKinds)
	return testCaseKinds
}

func GetTestKind(kindID int64) *TestCaseKind {
	testCaseKind := new(TestCaseKind)
	db.Where("kind_id = ?", kindID).Find(&testCaseKind)
	return testCaseKind
}
