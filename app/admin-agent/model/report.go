package model

import (
	"go-admin/common/models"
)

type Report struct {
	ReportId         int    `json:"reportId" gorm:"size:128;primaryKey;autoIncrement;comment:报告ID(主键)"`
	ReportName       string `json:"reportName" gorm:"comment:报告名称"`
	ReportProperties string `json:"reportProperties" gorm:"comment:报告详情"`
	Type             string `json:"Type" gorm:"size:64;comment:报告类型（侵权/估值）"`
	RejectTag        string `json:"rejectTag" gorm:"size:8;comment:驳回标签(null:未审核/reject/upload)"`
	models.ControlBy
	CreatedAt string `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt string `json:"UpdatedAt" gorm:"comment:最后更新时间"`
	Files     string `json:"files" comment:"报告文件"`
}

func (Report) TableName() string {
	return "report"
}

func (e *Report) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Report) GetId() interface{} {
	return e.ReportId
}
