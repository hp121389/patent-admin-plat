package models

import (
	"go-admin/common/models"
)

type Patent struct {
	//models.Model        //就是自增id
	PatentId int    `json:"PatentId" gorm:"size:128;primaryKey;autoIncrement;comment:专利ID(主键)"`
	TI       string `json:"TI" gorm:"size:128;comment:专利名"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号"`
	AD       string `json:"AD" gorm:"size:128;comment:申请日"`
	PD       string `json:"PD" gorm:"size:128;comment:公开日"`
	CL       string `json:"CL" gorm:"size:128;comment:简介"`
	PA       string `json:"PA" gorm:"size:128;comment:申请单位"`
	AR       string `json:"AR" gorm:"size:128;comment:地址"`
	INN      string `json:"INN" gorm:"size:128;comment:申请人"`
	models.ControlBy
	//嵌入结构体：先写好models然后嵌入，等效于models本体
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
