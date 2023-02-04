package models

import (
	"go-admin/common/models"
)

type Patent struct {
	PatentId         int    `json:"patentId" gorm:"size:128;primaryKey;autoIncrement;comment:专利ID(主键)"`
	PNM              string `json:"PNM" gorm:"size:128;comment:申请号"`
	PatentProperties string `json:"patentProperties" gorm:"comment:专利详情"`
	Desc             string `json:"desc" gorm:"-:all"`
	Price            int    `json:"price,omitempty" gorm:"-:all"`
	models.ControlBy
	models.ModelTime
}

func (Patent) TableName() string {
	return "patent"
}

func (e *Patent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Patent) GetId() interface{} {
	return e.PatentId
}
