package dao

import "time"

type Report struct {
	Id           int
	ReportID     int64
	startTime    time.Time
	NodeVersion  string
	ChainVersion string
	IsRunning    uint8
}

func (m *Report) TableName() string {
	return "report"
}

func GetReportID() int64 {
	report := new(Report)
	db.Order("report_id desc").First(&report)
	return report.ReportID
}
