package models

import "go-admin/common/models"

type SysDept struct {
	DeptId   int    `json:"sysId" gorm:"primaryKey;autoIncrement"` //部门Id
	DeptName string `json:"deptName" gorm:"size:128"`              //部门名称
	DeptDesc string `json:"deptDesc" gorm:"size:255"`              //部门描述
	models.ControlBy
}

func (SysDept) TableName() string {
	return "sys_dept"
}

func (e *SysDept) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDept) GetId() interface{} {
	return e.DeptId
}
