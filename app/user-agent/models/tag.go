package models

import "go-admin/common/models"

type Tag struct {
	TagId   int    `json:"tagId" gorm:"primaryKey;autoIncrement"` //标签ID
	TagName string `json:"tagName" gorm:"size:128;"`              //标签名称
	Desc    string `json:"desc"  gorm:"size:255"  `               //标签描述
	models.ControlBy
	models.ModelTime
}

func (Tag) TableName() string {
	return "tag"

}
func (e *Tag) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Tag) GetId() interface{} {
	return e.TagId
}
