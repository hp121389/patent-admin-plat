package model

import (
	"go-admin/common/models"
)

type PatentReport struct {
	Id       int `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	PatentId int `json:"patentId" gorm:"size:128;comment:专利ID"`
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	models.ControlBy
	models.ModelTime
}

func (PatentReport) TableName() string {
	return "patent_report"
}

func (e *PatentReport) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PatentReport) GetId() interface{} {
	return e.Id
}
