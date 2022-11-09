package dto

import (
	"encoding/json"
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PatentGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int    `form:"patentId" search:"type:exact;column:PatentId;table:patent" comment:"专利ID"`
	TI             string `form:"TI" search:"type:exact;column:TI;table:patent" comment:"专利名"`
	PNM            string `form:"PNM" search:"type:exact;column:PNN;table:patent" comment:"申请号"`
	AD             string `form:"AD" search:"type:exact;column:AD;table:patent" comment:"申请日"`
	PD             string `form:"PD" search:"type:exact;column:PD;table:patent" comment:"公开日"`
	CL             string `form:"CL" search:"type:exact;column:CL;table:patent" comment:"简介"`
	PA             string `form:"PA" search:"type:exact;column:PA;table:patent" comment:"申请单位"`
	AR             string `form:"AR" search:"type:exact;column:AR;table:patent" comment:"地址"`
	PINN           string `form:"PINN" search:"type:exact;column:PINN;table:patent" comment:"申请人"`
	CLS            string `json:"CLS" gorm:"size:128;comment:法律状态"`
	PatentOrder
}

type PatentOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:patent" form:"createdAtOrder"`
}

func (m *PatentGetPageReq) GetNeedSearch() interface{} {
	return *m
}
func (m *PatentGetPageReq) GetPatentId() interface{} {
	return m.PatentId
}

type PatentReq struct {
	PatentId int    `json:"patentId" gorm:"size:128;comment:专利ID"`
	TI       string `json:"TI" gorm:"size:128;comment:专利名"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号" ` //vd:"len($)>0"
	AD       string `json:"AD" gorm:"size:128;comment:申请日"`
	PD       string `json:"PD" gorm:"size:128;comment:公开日"`
	CL       string `json:"CL" gorm:"comment:简介"`
	PA       string `json:"PA" gorm:"size:128;comment:申请单位"`
	AR       string `json:"AR" gorm:"size:128;comment:地址"`
	PINN     string `json:"PINN" gorm:"size:128;comment:申请人"`
	CLS      string `json:"CLS" gorm:"size:128;comment:法律状态"`
	common.ControlBy
}

func (s *PatentReq) GenerateList(model *models.Patent) {
	if s.PatentId != 0 {
		model.PatentId = s.PatentId
	}
	model.PNM = s.PNM
	model.ControlBy = s.ControlBy
	pbs, _ := json.Marshal(s)
	model.PatentProperties = string(pbs)
}

type PatentById struct {
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}

type PatentsIds struct {
	PatentId  int   `json:"patent_Id"`
	PatentIds []int `json:"patent_Ids"`
}

func (s *PatentsIds) GetPatentId() []int {
	s.PatentIds = append(s.PatentIds, s.PatentId)
	return s.PatentIds
}

type PatentBriefInfo struct {
	PatentId int    `json:"patentId" gorm:"size:128;comment:专利ID"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号" vd:"len($)>0"`
}
