package dto

import (
	"go-admin/app/patent/models"
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
	PNM string `json:"PNM" gorm:"size:128;comment:申请号"`
	AD  string `json:"AD" gorm:"size:128;comment:申请日"`
	PD  string `json:"PD" gorm:"size:128;comment:公开日"`
	CL  string `json:"CL" gorm:"size:128;comment:简介"`
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
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号"`
	AD       string `json:"AD" gorm:"size:128;comment:申请日"`
	PD       string `json:"PD" gorm:"size:128;comment:公开日"`
	CL       string `json:"CL" gorm:"size:128;comment:简介"`
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
