package models

import "go-admin/common/models"

type PatentTag struct {
	models.Model
	PatentId int `gorm:"foreignKey:PatentId;comment:专利Id" json:"PatentId" `
	TagId    int `gorm:"comment:标签ID"  json:"TagId"`
	models.ControlBy
}

func (e *PatentTag) TableName() string {
	return "patent_tag"
}

func (e *PatentTag) GetTagId() interface{} {
	return e.TagId
}

func (e *PatentTag) GetPatentId() interface{} {
	return e.PatentId
}
