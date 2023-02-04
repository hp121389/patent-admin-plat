package dto

import (
	"encoding/json"
	"go-admin/app/user-agent/models"
	cDto "go-admin/common/dto"
	common "go-admin/common/models"
)

const PatentPriceBase = 8000

type PatentReq struct {
	PatentId int    `json:"patentId" gorm:"size:128;comment:专利ID"`
	TI       string `json:"TI" gorm:"size:128;comment:专利名"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号" ` //vd:"len($)>0"
	AD       string `json:"AD" gorm:"size:128;comment:申请日"`
	PD       string `json:"PD" gorm:"size:128;comment:公开日"`
	CL       string `json:"CL" gorm:"comment:简介"`
	PA       string `json:"PA" gorm:"size:128;comment:申请单位"`
	AR       string `json:"AR" gorm:"size:128;comment:地址"`
	PINN     string `json:"PINN" gorm:"size:128;comment:主发明人"`
	CLS      string `json:"CLS" gorm:"size:128;comment:法律状态"`
	INN      string `json:"INN" gorm:"size:128;comment:发明人"`
	IDX      int    `json:"IDX" gorm:"size:128;comment:价值指数"`
	Desc     string `json:"desc" gorm:"size:128;comment:描述"`
	common.ControlBy
}

func (s *PatentReq) GenerateList(model *models.Patent) {
	if s.PatentId != 0 {
		model.PatentId = s.PatentId
	}
	model.PNM = s.PNM
	model.ControlBy = s.ControlBy
	pbs, _ := json.Marshal(s)            //把s（json）转化为byte[]
	model.PatentProperties = string(pbs) //把byte[]转化为string
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

type PatentDescReq struct {
	PNM       string `json:"PNM"`
	UserId    int    `json:"userId"`
	PackageID int    `json:"packageId"`
	Desc      string `json:"desc"`

	common.ControlBy
}

func (r *PatentDescReq) GenerateUserPatent(model *models.UserPatent) {
	model.PNM = r.PNM
	model.UserId = r.UserId
	model.Desc = r.Desc
	model.CreateBy = r.CreateBy
	model.UpdateBy = r.UpdateBy
}

func (r *PatentDescReq) GeneratePatentPackage(model *models.PatentPackage) {
	model.PNM = r.PNM
	model.PackageId = r.PackageID
	model.Desc = r.Desc
}

type PatentPagesReq struct {
	cDto.Pagination
}

type FindPatentPagesReq struct {
	cDto.Pagination
	Query string `json:"query"`
}
