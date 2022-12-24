package model

import (
	"go-admin/common/models"
)

type ReportRelation struct {
	Id       int    `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	PatentId int    `json:"patentId" gorm:"size:128;comment:专利ID"`
	ReportId int    `json:"reportId" gorm:"size:128;comment:报告ID"`
	UserId   int    `json:"userId" gorm:"size:128;comment:用户ID"`
	Type     string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型"`
	models.ControlBy
	CreatedAt string `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt string `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (ReportRelation) TableName() string {
	return "report_rela"
}

func (e *ReportRelation) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ReportRelation) GetId() interface{} {
	return e.Id
}
