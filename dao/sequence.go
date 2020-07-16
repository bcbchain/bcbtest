package dao

import "sync/atomic"

var (
	kindID   int64 = 10000
	suiteID  int64 = 20000
	caseID   int64 = 30000
	reportID int64 = 40000
)

func GetKindIDSeq() int64 {
	atomic.AddInt64(&kindID, 1)
	return kindID
}

func GetSuiteIDSeq() int64 {
	atomic.AddInt64(&suiteID, 1)
	return suiteID
}

func GetCaseIDSeq() int64 {
	atomic.AddInt64(&caseID, 1)
	return caseID
}

func GetReportIDSeq() int64 {
	atomic.AddInt64(&reportID, 1)
	return reportID
}
