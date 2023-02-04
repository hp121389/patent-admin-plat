package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

//patent-package

type PackageBack struct {
	PatentId  int       `form:"patentId" search:"type:exact;column:TagId;table:patent_package" comment:"专利ID"`
	PackageId int       `form:"packageId" search:"type:exact;column:TagId;table:patent_package" comment:"专利包ID"`
	PNM       string    `json:"PNM" gorm:"size:128;comment:申请号"`
	Desc      string    `json:"desc" gorm:"size:128;comment:描述"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"comment:最后更新时间"`
}

type PackagePageGetReq struct {
	dto.Pagination `search:"-"`
	PackageBack
	common.ControlBy
	CreatedAt time.Time `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"comment:最后更新时间"`
}

func (d *PackagePageGetReq) GeneratePackagePatent(g *models.PatentPackage) {
	g.PatentId = d.PatentId
	g.PackageId = d.PackageId
	g.ControlBy = d.ControlBy
	g.PNM = d.PNM
	g.Desc = d.Desc
	g.CreatedAt = d.CreatedAt
	g.UpdatedAt = d.UpdatedAt

}

type PatentPackageReq struct {
	PatentId  int    `form:"patentId" search:"type:exact;column:TagId;table:patent_package" comment:"专利ID"`
	PackageId int    `form:"packageId" search:"type:exact;column:TagId;table:patent_package" comment:"专利包ID"`
	PNM       string `json:"PNM" gorm:"size:128;comment:申请号"`
	Desc      string `json:"desc" gorm:"size:128;comment:描述"`
	common.ControlBy
	CreatedAt time.Time `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"comment:最后更新时间"`
}

type IsPatentInPackageResp struct {
	Existed bool `json:"existed"`
}

func (d *PatentPackageReq) GeneratePackagePatent(g *models.PatentPackage) {
	g.PatentId = d.PatentId
	g.PackageId = d.PackageId
	g.ControlBy = d.ControlBy
	g.PNM = d.PNM
	g.Desc = d.Desc
	g.CreatedAt = d.CreatedAt
	g.UpdatedAt = d.UpdatedAt
}
