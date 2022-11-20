package models

import "go-admin/common/models"

type Package struct {
	PackageId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"packageId"`
	PackageName string `json:"packageName" gorm:"size:128;comment:专利包"`
	Desc        string `json:"desc" gorm:"size:128;comment:描述"`
	Files       string `json:"files" gorm:"comment:专利包附件"`
	models.ControlBy
	models.ModelTime
}

func (e *Package) TableName() string {
	return "package"
}

func (e *Package) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Package) GetId() interface{} {
	return e.PackageId
}
