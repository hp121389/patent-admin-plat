package models

import "go-admin/common/models"

type SysUserPatent struct {
	models.Model
	PatentId int    `gorm:"foreignKey:PatentId;comment:PatentId" json:"PatentId" `
	UserId   int    `gorm:"comment:用户ID"  json:"userId"`
	ID       int    `gorm:"primaryKey;autoIncrement;comment:编码" json:"ID" `
	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
	models.ControlBy
}

func (e *SysUserPatent) TableName() string {
	return "user_patent"
}

func (e *SysUserPatent) GetUserId() interface{} {
	return e.UserId
}

func (e *SysUserPatent) GetPatentId() interface{} {
	return e.PatentId
}
