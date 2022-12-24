package model

import "go-admin/common/models"

type Dept struct {
	DeptId           int    `json:"deptId" gorm:"size:128;primaryKey;autoIncrement;comment:团队ID(主键)"`
	DeptName         string `json:"deptName" gorm:"comment:团队名称"`
	DeptProperties   string `json:"deptProperties" gorm:"comment:团队详情"`
	ResearchInterest string `json:"researchInterest" gorm:"comment:研究方向"`
	CreatedAt        string `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt        string `json:"UpdatedAt" gorm:"comment:最后更新时间"`
	DeptStatus       string `json:"deptStatus" gorm:"comment:团队状态"`
	models.ControlBy
}

func (Dept) TableName() string {
	return "dept"
}

func (e *Dept) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Dept) GetId() interface{} {
	return e.DeptId
}
