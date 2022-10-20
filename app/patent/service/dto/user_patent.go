package dto

import (
	"go-admin/app/patent/models"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

const (
	ClaimType = "认领"
	FocusType = "关注"
)

type UserPatentGetPageReq struct {
	dto.Pagination `search:"-"`
	UserId         int    `form:"UserId" search:"type:exact;column:UserId;table:user_patent" comment:"用户ID" `
	PatentId       int    `form:"PatentId" search:"type:exact;column:TagId;table:user_patent" comment:"专利ID" `
	Type           string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
	PatentTagOrder
}

type UserPatentOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:user_patent" form:"createdAtOrder"`
}

func (d *UserPatentGetPageReq) GetNeedSearch() interface{} {
	return *d
}

func (d *UserPatentGetPageReq) GetUserId() interface{} {
	return d.UserId
}

func (d *UserPatentGetPageReq) GetPatentId() interface{} {
	return d.PatentId
}

//type UserPatentInsertReq struct {
//	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
//	UserId   int    `json:"UserId" gorm:"size:128;comment:用户ID"`
//	PatentId int    `json:"PatentId" gorm:"size:128;comment:专利ID"`
//	common.ControlBy
//}
//
//func (s *UserPatentInsertReq) GenerateUserPatent(g *models.UserPatent) {
//	g.PatentId = s.PatentId
//	g.UserId = s.UserId
//	g.Type = s.Type
//}

type UserPatentObject struct {
	UserId   int    `json:"UserId" gorm:"size:128;comment:用户ID"`
	PatentId int    `uri:"patent_id"`
	Type     string `uri:"type"` //路由对大小写敏感
	common.ControlBy
}

func (d *UserPatentObject) GetPatentId() interface{} {
	return d.PatentId
}

func (d *UserPatentObject) GetType() interface{} {
	return d.Type
}

func (d *UserPatentObject) GenerateUserPatent(g *models.UserPatent) {
	g.PatentId = d.PatentId
	g.UserId = d.UserId
	g.Type = d.Type
}

func NewUserPatentClaim(userId, patentId, createdBy int) *UserPatentObject {
	return &UserPatentObject{
		UserId:   userId,
		PatentId: patentId,
		Type:     ClaimType,
		ControlBy: common.ControlBy{
			CreateBy: createdBy,
		},
	}
}

func NewUserPatentFocus(userId, patentId, createdBy int) *UserPatentObject {
	return &UserPatentObject{
		UserId:   userId,
		PatentId: patentId,
		Type:     FocusType,
		ControlBy: common.ControlBy{
			CreateBy: createdBy,
		},
	}
}

type UpDateUserPatentObject struct {
	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
	UserId   int    `json:"UserId" gorm:"size:128;comment:用户ID"`
	PatentId int    `json:"PatentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}

func (s *UpDateUserPatentObject) GenerateUserPatent(g *models.UserPatent) {
	g.PatentId = s.PatentId
	g.UserId = s.UserId
	g.Type = s.Type

}
