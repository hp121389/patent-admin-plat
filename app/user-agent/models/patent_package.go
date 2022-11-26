package models

import (
	"go-admin/common/models"
	"time"
)

//----------------------------------------patent-package----------------------------------------

type PatentPackage struct {
	models.Model
	PatentId  int    `gorm:"foreignKey:PatentId;comment:专利Id" json:"PatentId" `
	PackageId int    `gorm:"comment:专利包ID"  json:"PackageId"`
	PNM       string `json:"PNM" gorm:"size:128;comment:申请号"`
	models.ControlBy
	CreatedAt time.Time `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (e *PatentPackage) TableName() string {
	return "patent_package"
}
