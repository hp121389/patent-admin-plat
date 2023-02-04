package model

import "go-admin/common/models"

type Ticket struct {
	ID         int    `json:"id" gorm:"size:128;primaryKey;autoIncrement;comment:工单ID(主键)"`
	RelaID     int    `json:"relaID" gorm:"size:128;comment:对象ID"`
	Name       string `json:"name" gorm:"size:128;comment:工单名称"`
	Properties string `json:"properties" gorm:"comment:工单详情"`
	Type       string `json:"type" gorm:"size:64;comment:工单类型"`
	Status     string `json:"status" gorm:"size:64;comment:工单状态"`
	OptLogs    string `json:"optLogs" gorm:"comment:操作日志"`
	models.ControlBy
	models.ModelTime

	RelObj interface{} `json:"relObj" gorm:"-"`
}

func (Ticket) TableName() string {
	return "ticket"
}

func (e *Ticket) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Ticket) GetId() interface{} {
	return e.ID
}
