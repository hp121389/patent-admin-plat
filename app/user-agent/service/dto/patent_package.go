package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

//patent-package

type PackageBack struct {
	PatentId  int `form:"patentId" search:"type:exact;column:TagId;table:patent_package" comment:"专利ID"`
	PackageId int `form:"packageId" search:"type:exact;column:TagId;table:patent_package" comment:"专利包ID"`
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

}
