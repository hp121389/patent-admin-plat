package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

//patent-package

type PackageBack struct {
	PatentId  int    `form:"patentId" search:"type:exact;column:TagId;table:patent_package" comment:"专利ID"`
	PackageId int    `form:"packageId" search:"type:exact;column:TagId;table:patent_package" comment:"专利包ID"`
	PNM       string `json:"PNM" gorm:"size:128;comment:申请号"`
}

type PackagePageGetReq struct {
	dto.Pagination `search:"-"`
	PackageBack
	PatentTagOrder
	common.ControlBy
}

func (d *PackagePageGetReq) GeneratePackagePatent(g *models.PatentPackage) {
	g.PatentId = d.PatentId
	g.PackageId = d.PackageId
	g.ControlBy = d.ControlBy
	g.PNM = d.PNM
}

type PatentPackageReq struct {
	PatentId  int    `form:"patentId" search:"type:exact;column:TagId;table:patent_package" comment:"专利ID"`
	PackageId int    `form:"packageId" search:"type:exact;column:TagId;table:patent_package" comment:"专利包ID"`
	PNM       string `json:"PNM" gorm:"size:128;comment:申请号"`
	common.ControlBy
}

type IsPatentInPackageResp struct {
	Existed bool `json:"existed"`
}

func (d *PatentPackageReq) GeneratePackagePatent(g *models.PatentPackage) {
	g.PatentId = d.PatentId
	g.PackageId = d.PackageId
	g.ControlBy = d.ControlBy
	g.PNM = d.PNM
}
