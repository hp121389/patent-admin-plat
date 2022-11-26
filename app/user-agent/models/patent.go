package models

import (
	"go-admin/common/models"
	"time"
)

type Patent struct {
	PatentId         int    `json:"patentId" gorm:"size:128;primaryKey;autoIncrement;comment:专利ID(主键)"`
	PNM              string `json:"PNM" gorm:"size:128;comment:申请号"`
	PatentProperties string `json:"patentProperties" gorm:"comment:专利详情"`
	models.ControlBy
	CreatedAt time.Time `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"comment:最后更新时间"`
}

func (Patent) TableName() string {
	return "user-agent"
}

func (e *Patent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Patent) GetId() interface{} {
	return e.PatentId
}
