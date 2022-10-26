package models

import "go-admin/common/models"

//----------------------------------------user-patent----------------------------------------

type UserPatent struct {
	models.Model
	PatentId int    `gorm:"foreignKey:PatentId;comment:PatentId" json:"PatentId" `
	UserId   int    `gorm:"comment:用户ID"  json:"UserId"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号"`
	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
	models.ControlBy
}

func (e *UserPatent) TableName() string {
	return "user_patent"
}
