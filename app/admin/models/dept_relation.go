package models

import (
	"go-admin/common/models"
)

type DeptRelation struct {
	Id     int `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	UserId int `json:"userId" gorm:"size:128;comment:成员ID"`
	DeptId int `json:"deptId" gorm:"size:128;comment:部门ID"`
	models.ModelTime
	models.ControlBy
}

func (e *DeptRelation) TableName() string {
	return "dept_rela"
}

func (e *DeptRelation) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *DeptRelation) GetId() interface{} {
	return e.Id
}
