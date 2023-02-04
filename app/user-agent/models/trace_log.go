package models

import "go-admin/common/models"

type TraceLog struct {
	TraceID int    `json:"traceID" gorm:"size:128;primaryKey;autoIncrement;comment:traceID(主键)"`
	Action  string `json:"action" gorm:"size:128;comment:操作"`
	Desc    string `json:"desc" gorm:"size:1024;comment:描述"`
	Request string `json:"request" gorm:"size:1024;comment:用户请求"`
	models.ControlBy
	models.ModelTime
}

func (e *TraceLog) TableName() string {
	return "trace_log"
}

func (e *TraceLog) Generate() models.ActiveRecord {
	return e
}

func (e *TraceLog) GetId() interface{} {
	return e.TraceID
}
