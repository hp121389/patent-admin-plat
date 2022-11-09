package models

import "go-admin/common/models"

//----------------------------------------patent-package----------------------------------------

type PatentPackage struct {
	models.Model
	PatentId  int    `gorm:"foreignKey:PatentId;comment:专利Id" json:"PatentId" `
	PackageId int    `gorm:"comment:专利包ID"  json:"PackageId"`
	PNM       string `json:"PNM" gorm:"size:128;comment:申请号"`
	models.ControlBy
}

func (e *PatentPackage) TableName() string {
	return "patent_package"
}
