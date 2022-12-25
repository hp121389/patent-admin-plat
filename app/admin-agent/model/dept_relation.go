package model

import (
	"go-admin/common/models"
)

type DeptRelation struct {
	Id            int    `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	UserId        int    `json:"userId" gorm:"size:128;comment:成员ID"`
	DeptId        int    `json:"deptId" gorm:"size:128;comment:部门ID"`
	MemType       string `form:"memType" search:"type:exact;column:MemType;table:dept" comment:"成员类型"`
	CreatedAt     string `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     string `json:"updatedAt" gorm:"comment:最后更新时间"`
	MemStatus     string `json:"memStatus" gorm:"comment:成员状态"`
	ExamineStatus string `json:"ExamineStatus" gorm:"comment:审核状态"`
	models.ControlBy
}

func (DeptRelation) TableName() string {
	return "dept_rela"
}

func (e *DeptRelation) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *DeptRelation) GetId() interface{} {
	return e.Id
}
