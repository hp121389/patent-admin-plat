package models

import "go-admin/common/models"

type UserTag struct {
	UserTagId int `json:"UserTagId" gorm:"primaryKey;autoIncrement"` //用户-标签关系ID
	UserId    int `json:"UserId" gorm:""`                            //用户ID
	TagId     int `json:"TagId"  gorm:""  `                          //标签ID
	models.ControlBy
}

func (UserTag) TableName() string {
	return "user_tag"
}

func (e *UserTag) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserTag) GetId() interface{} {
	return e.TagId
}
