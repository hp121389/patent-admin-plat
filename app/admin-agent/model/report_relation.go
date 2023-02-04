package model

import (
	"go-admin/common/models"
)

type ReportRelation struct {
	Id       int `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	TicketId int `json:"ticketId" gorm:"size:128;comment:工单ID"`
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	UserId   int `json:"userId" gorm:"size:128;comment:用户ID"`
	models.ModelTime
	models.ControlBy
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
