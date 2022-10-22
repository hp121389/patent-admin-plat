package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

//查询必须写form字段

type PatentGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int    `form:"PatentId" search:"type:exact;column:PatentId;table:patent" comment:"专利ID"`
	TI             string `form:"TI" search:"type:exact;column:TI;table:patent" comment:"专利名"`
	PNM            string `form:"PNM" search:"type:exact;column:PNN;table:patent" comment:"申请号"`
	AD             string `form:"AD" search:"type:exact;column:AD;table:patent" comment:"申请日"`
	PD             string `form:"PD" search:"type:exact;column:PD;table:patent" comment:"公开日"`
	CL             string `form:"CL" search:"type:exact;column:CL;table:patent" comment:"简介"`
	PA             string `form:"PA" search:"type:exact;column:PA;table:patent" comment:"申请单位"`
	AR             string `form:"AR" search:"type:exact;column:AR;table:patent" comment:"地址"`
	INN            string `form:"INN" search:"type:exact;column:INN;table:patent" comment:"申请人"`
	PatentOrder
}

type PatentUpdateReq struct {
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`

	TI  string `json:"TI" gorm:"size:128;comment:专利名"`
	PNM string `json:"PNM" gorm:"size:128;comment:申请号" vd:"len($)>0"`
	AD  string `json:"AD" gorm:"size:128;comment:申请日"`
	PD  string `json:"PD" gorm:"size:128;comment:公开日"`
	CL  string `json:"CL" gorm:"comment:简介"`
	PA  string `json:"PA" gorm:"size:128;comment:申请单位"`
	AR  string `json:"AR" gorm:"size:128;comment:地址"`
	INN string `json:"INN" gorm:"size:128;comment:申请人"`
	common.ControlBy
}

func (s PatentUpdateReq) GetPatentId() interface{} {
	return s.PatentId
}

type PatentOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:patent" form:"createdAtOrder"`
}

func (m *PatentGetPageReq) GetNeedSearch() interface{} {
	return *m
}
func (s *PatentGetPageReq) GetPatentId() interface{} {

	return s.PatentId
}

func (s *PatentUpdateReq) GenerateList(model *models.Patent) {
	if s.PatentId != 0 {
		model.PatentId = s.PatentId
	}
	model.TI = s.TI
	model.CL = s.CL
	model.AR = s.AR
	model.PNM = s.PNM
	model.AD = s.AD
	model.PD = s.PD
	model.INN = s.INN
	model.PA = s.PA
}

type PatentGetReq struct {
	PatentId int `uri:"patent_id"`
}

func (s *PatentGetReq) GetPatentId() interface{} {
	return s.PatentId
}

// PatentDeleteReq 功能删除请求参数

type PatentDeleteReq struct {
	PatentId int `json:"PatentIds"`
}

func (s *PatentDeleteReq) GetPatentId() interface{} {
	return s.PatentId
}

type PatentInsertReq struct {
	PatentId int    `json:"PatentId" gorm:"size:128;comment:专利ID"`
	TI       string `json:"TI" gorm:"size:128;comment:专利名"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号" vd:"len($)>0"`
	AD       string `json:"AD" gorm:"size:128;comment:申请日"`
	PD       string `json:"PD" gorm:"size:128;comment:公开日"`
	CL       string `json:"CL" gorm:"comment:简介"`
	PA       string `json:"PA" gorm:"size:128;comment:申请单位"`
	AR       string `json:"AR" gorm:"size:128;comment:地址"`
	INN      string `json:"INN" gorm:"size:128;comment:申请人"`
	common.ControlBy
}

func (s *PatentInsertReq) GenerateList(model *models.Patent) {
	if s.PatentId != 0 {
		model.PatentId = s.PatentId
	}
	model.TI = s.TI
	model.CL = s.CL
	model.AR = s.AR
	model.PNM = s.PNM
	model.AD = s.AD
	model.PD = s.PD
	model.INN = s.INN
	model.PA = s.PA
	model.CreateBy = s.CreateBy
}

func (s *PatentInsertReq) GetPatentId() interface{} {
	return s.PatentId
}

type PatentById struct {
	dto.ObjectByPatentId
	common.ControlBy
}

func (s *PatentById) GetPatentId() interface{} {
	return s.PatentId
}

func (s *PatentById) GenerateM() (common.ActiveRecord, error) {
	return &models.Patent{}, nil
}

type PatentsByIdsForRelationshipUsers struct {
	dto.ObjectOfPatentId
}

func (s *PatentsByIdsForRelationshipUsers) GetPatentId() []int {

	s.PatentIds = append(s.PatentIds, s.PatentId)
	return s.PatentIds

}

func (s *PatentsByIdsForRelationshipUsers) GetNeedSearch() interface{} {
	return *s
}

type PatentsByIdsForRelationshipTags struct {
	dto.ObjectOfPatentId
}

func (s *PatentsByIdsForRelationshipTags) GetNeedSearch() interface{} {
	return *s
}

func (s *PatentsByIdsForRelationshipTags) GetPatentId() []int {

	s.PatentIds = append(s.PatentIds, s.PatentId)
	return s.PatentIds
}

//user-patent
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

type UserPatentObject struct {
	UserId   int    `json:"UserId" gorm:"size:128;comment:用户ID"`
	PatentId int    `form:"PatentId" search:"type:exact;column:TagId;table:user_patent" comment:"专利ID" `
	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
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

func NewUserPatentClaim(userId, patentId, createdBy, updatedBy int) *UserPatentObject {
	return &UserPatentObject{
		UserId:   userId,
		PatentId: patentId,
		Type:     ClaimType,
		ControlBy: common.ControlBy{
			CreateBy: createdBy,
			UpdateBy: updatedBy,
		},
	}
}

func NewUserPatentFocus(userId, patentId, createdBy, updatedBy int) *UserPatentObject {
	return &UserPatentObject{
		UserId:   userId,
		PatentId: patentId,
		Type:     FocusType,
		ControlBy: common.ControlBy{
			CreateBy: createdBy,
			UpdateBy: updatedBy,
		},
	}
}

//patent-tag

type PatentTagGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int `form:"PatentId" search:"type:exact;column:TagId;table:patent_tag" comment:"专利ID"`
	TagId          int `form:"TagId" search:"type:exact;column:TagId;table:patent_tag" comment:"标签ID"`
	PatentTagOrder
}

type TagPageGetReq struct {
	dto.Pagination `search:"-"`
	PatentId       int `form:"PatentId" search:"type:exact;column:TagId;table:patent_tag" comment:"专利ID"`
	TagId          int `form:"TagId" search:"type:exact;column:TagId;table:patent_tag" comment:"标签ID"`
	PatentTagOrder
}

func (m *TagPageGetReq) GetPatentId() interface{} {
	return m.PatentId
}

func (m *TagPageGetReq) GetTagId() interface{} {
	return m.TagId
}

type PatentTagOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:patent_tag" form:"CreatedAtOrder"`
}

func (m *PatentTagGetPageReq) GetNeedSearch() interface{} {
	return *m
}

func (m *PatentTagGetPageReq) GetPatentId() interface{} {
	return m.PatentId
}

func (m *PatentTagGetPageReq) GetTagId() interface{} {
	return m.TagId
}

type PatentTagInsertReq struct {
	TagId    int `json:"TagId" gorm:"size:128;comment:标签ID"`
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}

func (d *PatentTagInsertReq) GeneratePatentTag(g *models.PatentTag) {
	g.PatentId = d.PatentId
	g.TagId = d.TagId
}

func (d *PatentTagInsertReq) GetPatentId() interface{} {
	return d.PatentId
}

func (d *PatentTagInsertReq) GetTagId() interface{} {
	return d.TagId
}

type TagUpdateReqByPatent struct {
	TagId    int `json:"TagId" gorm:"size:128;comment:标签ID"`
	PatentId int `uri:"patent_id"`
	common.ControlBy
}

type PatentUpdateReqByTag struct {
	TagId    int `uri:"tag_id"`
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}
