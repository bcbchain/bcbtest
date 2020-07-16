package parse

import (
	"bcbtest/config"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"sort"
)

var (
	orderSheetKey []int
	testKind      *TestKind
	testKinds     = &TestKinds{}
	cases         = make([]CaseInfo, 0)
	tSuites       = NewTSuite()
	sSuites       = NewSSuite()
	fSuites       = NewFSuite()
	thirdDesc     = ""
	secondDesc    = ""
	firstDesc     = ""
)

func Init() {
	cfg := config.GetConfig()
	f, err := excelize.OpenFile(cfg.ExcelFile)
	if err != nil {
		panic(err)
	}

	//db := dao.GetClient()
	//db.Delete(&dao.TestCaseKind{})
	//dao.TrunkTestKind()
	//dao.TrunkTestSuite()
	//dao.TrunkTestCase()

	// 遍历所有sheet
	for index, _ := range f.GetSheetMap() {
		orderSheetKey = append(orderSheetKey, index)
	}

	// 对sheetKey排序
	sort.Ints(orderSheetKey)

	for _, sheetKey := range orderSheetKey {
		//kindID := GetKindIDSeq()
		//db.Create(&dao.TestCaseKind{
		//	KindID:   kindID,
		//	KindName: f.GetSheetName(sheetKey),
		//})
		rows := f.GetRows(f.GetSheetName(sheetKey))
		depth := getDepth(rows)

		switch depth {
		case 5:
			fSuites = threeLevelParse(rows)
		case 4:
			fSuites = twoLevelParse(rows)
		default:
			fSuites = NewFSuite()
		}
		fmt.Println(f.GetSheetName(sheetKey))
		testKind = NewTestKind(GetKindIDSeq(), GetReportIDSeq(), f.GetSheetName(sheetKey), fSuites, true)
		*testKinds = append(*testKinds, testKind)

		//db.Create(&dao.TestCaseKind{
		//	KindID:   testKind.KindID,
		//	KindName: testKind.KindName,
		//})
		//for _, firstV := range testKind.Suites.orderKey {
		//	for _, secondV := range testKind.Suites.Value[firstV].orderKey {
		//		for _, thirdV := range testKind.Suites.Value[firstV].Value[secondV].orderKey {
		//			suiteID := dao.GetSuiteIDSeq()
		//			db.Create(&dao.TestSuite{
		//				SuiteID:     suiteID,
		//				KindID:      testKind.KindID,
		//				FirstSuite:  firstV,
		//				SecondSuite: secondV,
		//				ThirdSuite:  thirdV,
		//			})
		//			for _, cases := range testKind.Suites.Value[firstV].Value[secondV].Value[thirdV] {
		//				db.Create(&dao.TestCase{
		//					CaseID:       dao.GetCaseIDSeq(),
		//					SuiteID:      suiteID,
		//					Desc:         cases.Desc,
		//					RepeatCount:  0,
		//					ExpectedCode: 0,
		//				})
		//			}
		//		}
		//	}
		//}

	}
}

func getDepth(rows [][]string) int {
	depth := len(rows[0])
	for i := len(rows[0]) - 1; i > 0; i-- {
		if rows[0][i] != "" {
			break
		}
		depth--
	}
	return depth
}

func threeLevelParse(rows [][]string) FirstSuites {
	rows = rows[1:]
	fSuites = NewFSuite()
	for _, row := range rows {
		// 这行数据中包含第一级集合
		if row[0] != "" && row[1] != "" && row[2] != "" && row[3] != "" {
			firstDesc = row[0]
			secondDesc = row[1]
			thirdDesc = row[2]

			cases = make([]CaseInfo, 0)
			sSuites = NewSSuite()
			tSuites = NewTSuite()
			//fSuites = NewFSuite(fSuitesKey)

			cInfo := CaseInfo{
				Desc: row[3],
				Exp:  row[4],
			}
			cases = append(cases, cInfo)
			tSuites.Value[row[2]] = cases
			tSuites.OrderKey = append(tSuites.OrderKey, row[2])

			sSuites.Value[row[1]] = tSuites
			sSuites.OrderKey = append(sSuites.OrderKey, row[1])

			fSuites.Value[row[0]] = sSuites
			fSuites.OrderKey = append(fSuites.OrderKey, row[0])
		}

		if row[0] == "" && row[1] == "" && row[2] == "" && row[3] != "" {
			cInfo := CaseInfo{
				Desc: row[3],
				Exp:  row[4],
			}
			cases = append(cases, cInfo)
			tSuites.Value[thirdDesc] = cases
			sSuites.Value[secondDesc] = tSuites
			fSuites.Value[firstDesc] = sSuites
		}

		if row[0] == "" && row[1] == "" && row[2] != "" && row[3] != "" {
			thirdDesc = row[2]
			cases = make([]CaseInfo, 0)

			cInfo := CaseInfo{
				Desc: row[3],
				Exp:  row[4],
			}
			cases = append(cases, cInfo)

			//t1 := NewTSuite(tSuitesKey)
			tSuites.Value[row[2]] = cases
			tSuites.OrderKey = append(tSuites.OrderKey, row[2])

			sSuites.Value[secondDesc] = tSuites
			fSuites.Value[firstDesc] = sSuites
		}

		if row[0] == "" && row[1] != "" && row[2] != "" && row[3] != "" {
			//sSuitesKey = 0
			thirdDesc = row[2]
			secondDesc = row[1]
			cases = make([]CaseInfo, 0)
			tSuites = NewTSuite()

			cInfo := CaseInfo{
				Desc: row[3],
				Exp:  row[4],
			}
			cases = append(cases, cInfo)

			tSuites.Value[row[2]] = cases
			tSuites.OrderKey = append(tSuites.OrderKey, row[2])

			sSuites.Value[secondDesc] = tSuites
			sSuites.OrderKey = append(sSuites.OrderKey, row[1])

			fSuites.Value[firstDesc] = sSuites
		}

		if row[0] == "" && row[1] == "" && row[2] == "" && row[3] == "" {
			continue
		}
	}
	return fSuites
}

func twoLevelParse(rows [][]string) FirstSuites {
	rows = rows[1:]
	fSuites = NewFSuite()
	for _, row := range rows {
		// 这行数据中包含第一级集合
		if row[0] != "" && row[1] != "" && row[2] != "" {
			firstDesc = row[0]
			secondDesc = row[1]
			thirdDesc = row[1]

			cases = make([]CaseInfo, 0)
			sSuites = NewSSuite()
			tSuites = NewTSuite()

			cInfo := CaseInfo{
				Desc: row[2],
				Exp:  row[3],
			}
			cases = append(cases, cInfo)
			tSuites.Value[row[1]] = cases
			tSuites.OrderKey = append(tSuites.OrderKey, row[1])
			sSuites.Value[row[1]] = tSuites
			sSuites.OrderKey = append(sSuites.OrderKey, row[1])
			fSuites.Value[row[0]] = sSuites
			fSuites.OrderKey = append(fSuites.OrderKey, row[0])
		}

		if row[0] == "" && row[1] == "" && row[2] != "" {
			cInfo := CaseInfo{
				Desc: row[2],
				Exp:  row[3],
			}
			cases = append(cases, cInfo)
			tSuites.Value[thirdDesc] = cases
			sSuites.Value[secondDesc] = tSuites
			fSuites.Value[firstDesc] = sSuites
		}

		if row[0] == "" && row[1] != "" && row[2] != "" {
			thirdDesc = row[1]
			secondDesc = row[1]
			cases = make([]CaseInfo, 0)

			cInfo := CaseInfo{
				Desc: row[2],
				Exp:  row[3],
			}
			cases = append(cases, cInfo)

			tSuites = NewTSuite()
			tSuites.Value[row[1]] = cases
			tSuites.OrderKey = append(tSuites.OrderKey, row[1])
			sSuites.Value[row[1]] = tSuites
			sSuites.OrderKey = append(sSuites.OrderKey, row[1])
			fSuites.Value[firstDesc] = sSuites
		}
		//fmt.Println(fSuites)

		//if row[0] == "" && row[1] != "" && row[2] != "" {
		//	thirdDesc = row[1]
		//	secondDesc = row[1]
		//	cases = make([]CaseInfo, 0)
		//	tSuites = make(thirdSuites, 0)
		//
		//	cInfo := CaseInfo{
		//		Desc: row[2],
		//		Exp:  row[3],
		//	}
		//	cases = append(cases, cInfo)
		//
		//	tSuites[row[1]] = cases
		//	sSuites[secondDesc] = tSuites
		//	fSuites[firstDesc] = sSuites
		//}

		if row[0] == "" && row[1] == "" && row[2] == "" {
			continue
		}
		fmt.Println(fSuites)
	}
	return fSuites
}

func GetTestKinds() *TestKinds {
	if testKinds == nil {
		log.Print("must InitExcel first.")
	}
	return testKinds
}
