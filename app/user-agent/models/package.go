package models

import "go-admin/common/models"

type Package struct {
	PackageId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"PackageId"`
	PackageName string `json:"PackageName" gorm:"size:128;comment:专利包"`
	Desc        string `json:"Desc" gorm:"size:128;comment:描述"`
	models.ControlBy
	models.ModelTime
}

type UserPackage struct {
	models.Model
	PackageId int `gorm:"foreignKey:PackageId;comment:PackageId" json:"PackageId" `
	UserId    int `gorm:"comment:用户ID"  json:"UserId"`
	ID        int `gorm:"primaryKey;autoIncrement;comment:编码" json:"Id" `
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

func (e *UserPackage) GetPackageId() interface{} {
	return e.PackageId
}

func (e *UserPackage) TableName() string {
	return "user_package"
}

func (e *UserPackage) GetUserId() interface{} {
	return e.UserId
}
