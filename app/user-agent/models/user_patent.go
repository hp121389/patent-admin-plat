package models

import (
	"go-admin/common/models"
)

//----------------------------------------user-patent----------------------------------------

type UserPatent struct {
	models.Model
	PatentId int    `gorm:"foreignKey:PatentId;comment:PatentId" json:"patentId" `
	UserId   int    `gorm:"comment:用户ID"  json:"userId"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号"`
	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`

	CreatedAt string `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt string `json:"updatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}

func (e *UserPatent) TableName() string {
	return "user_patent"
}
func (e *UserPatent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserPatent) GetId() interface{} {
	return e.UserId
}
