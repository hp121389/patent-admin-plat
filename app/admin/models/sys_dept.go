package models

import "go-admin/common/models"

type SysDept struct {
	DeptId         int    `json:"deptId" gorm:"size:128;primaryKey;autoIncrement;comment:团队ID(主键)"`
	DeptName       string `json:"deptName" gorm:"comment:团队名称"`
	DeptProperties string `json:"deptProperties" gorm:"comment:团队详情"`
	Status         string `json:"status" gorm:"comment:团队状态"`
	models.ModelTime
	models.ControlBy

	Members int `json:"members" gorm:"-"`
}

func (SysDept) TableName() string {
	return "dept"
}

func (e *SysDept) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDept) GetId() interface{} {
	return e.DeptId
}
