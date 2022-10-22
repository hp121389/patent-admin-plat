package dto

import (
	"go-admin/app/user-agent/models"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PackageGetPageReq struct {
	dto.Pagination `search:"-"`
	PackageId      int    `form:"PackageId" search:"type:exact;column:package_id;table:package" comment:"专利包ID"`
	PackageName    string `form:"PackageName" search:"type:contains;column:package_name;table:package" comment:"专利包名"`
	Desc           string `form:"Desc" search:"type:contains;column:desc;table:package" comment:"描述"`
}

func (m *PackageGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PackageInsertReq struct {
	PackageId   int    `json:"PackageId" comment:"专利包ID"` // 专利包ID
	PackageName string `json:"PackageName" comment:"专利包名" vd:"len($)>0"`
	Desc        string `json:"Desc" comment:"描述"`
	common.ControlBy
}

func (s *PackageInsertReq) GenerateList(model *models.Package) {
	if s.PackageId != 0 {
		model.PackageId = s.PackageId
	}
	model.PackageName = s.PackageName
	model.Desc = s.Desc
}

func (s *PackageInsertReq) GetId() interface{} {
	return s.PackageId
}

type PackageUpdateReq struct {
	PackageId   int    `json:"PackageId" comment:"专利包ID"` // 专利包ID
	PackageName string `json:"PackageName" comment:"专利包名"`
	Desc        string `json:"Desc" comment:"描述"`
	common.ControlBy
}

func (s *PackageUpdateReq) Generate(model *models.Package) {
	if s.PackageId != 0 {
		model.PackageId = s.PackageId
	}
	model.PackageName = s.PackageName
	model.Desc = s.Desc
}

func (s *PackageUpdateReq) GetId() interface{} {
	return s.PackageId
}

type PackageById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *PackageById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

type PackagesByIdsForRelationshipUsers struct {
	dto.ObjectOfPackageId
}

func (s *PackagesByIdsForRelationshipUsers) GetPackageId() []int {

	s.PackageIds = append(s.PackageIds, s.PackageId)
	return s.PackageIds

}

type UserPackageGetPageReq struct {
	dto.Pagination `search:"-"`
	UserId         int `form:"UserId" search:"type:exact;column:user_id;table:user_package" comment:"用户ID"`
	PackageId      int `form:"PackageId" search:"type:exact;column:package_id;table:user_package" comment:"专利包ID"`
	UserPackageOrder
}

type UserPackageOrder struct {
	PackageIdOrder string `search:"type:order;column:package_id;table:user_package" form:"PackageIdOrder"`
}

func (m *UserPackageGetPageReq) GetNeedSearch() interface{} {
	return *m
}

func (d *UserPackageGetPageReq) GetUserId() interface{} {
	return d.UserId
}

func (d *UserPackageGetPageReq) GetPackageId() interface{} {
	return d.PackageId
}

type UserPackageInsertReq struct {
	UserId    int `form:"UserId" search:"type:exact;column:user_id;table:user_package" comment:"用户ID"`
	PackageId int `form:"PackageId" search:"type:exact;column:package_id;table:user_package" comment:"专利包ID"`
	common.ControlBy
}

func (s *UserPackageInsertReq) GenerateUserPackage(g *models.UserPackage) {
	g.PackageId = s.PackageId
	g.UserId = s.UserId

}

type UserPackageObject struct {
	UserId    int `form:"UserId" search:"type:exact;column:user_id;table:user_package" comment:"用户ID"`
	PackageId int `uri:"package_id"`
	common.ControlBy
}

func (d *UserPackageObject) GetPackageId() interface{} {
	return d.PackageId
}

func NewUserPackageInsert(userId, pId int) *UserPackageInsertReq {
	return &UserPackageInsertReq{
		UserId:    userId,
		PackageId: pId,
	}
}
