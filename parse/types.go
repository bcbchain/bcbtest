package parse

type ThirdSuites struct {
	OrderKey []string
	Value    map[string][]CaseInfo
}

type SecondSuites struct {
	OrderKey []string
	Value    map[string]ThirdSuites
}

type FirstSuites struct {
	OrderKey []string
	Value    map[string]SecondSuites
}

type CaseInfo struct {
	Desc string
	Exp  string
}

type TestKind struct {
	KindID   int64
	ReportID int64
	KindName string
	Suites   FirstSuites
	IsReport bool
}

type TestKinds []*TestKind

func NewTestKind(kindID int64, reportID int64, kindName string, suites FirstSuites, isReport bool) *TestKind {
	return &TestKind{
		KindID:   kindID,
		ReportID: reportID,
		KindName: kindName,
		Suites:   suites,
		IsReport: isReport,
	}
}

func NewTSuite() ThirdSuites {
	return ThirdSuites{
		OrderKey: make([]string, 0),
		Value:    make(map[string][]CaseInfo, 0),
	}
}

func NewSSuite() SecondSuites {
	return SecondSuites{
		OrderKey: make([]string, 0),
		Value:    make(map[string]ThirdSuites, 0),
	}
}

func NewFSuite() FirstSuites {
	return FirstSuites{
		OrderKey: make([]string, 0),
		Value:    make(map[string]SecondSuites, 0),
	}
}
