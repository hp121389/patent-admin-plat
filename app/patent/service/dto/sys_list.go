package dto

import (
	"go-admin/app/patent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

//查询必须写form字段

type SysListGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int    `form:"PatentId" search:"type:exact;column:PatentId;table:sys_list" comment:"专利ID"`
	TI             string `form:"TI" search:"type:exact;column:TI;table:sys_list" comment:"专利名"`
	PNM            string `form:"PNM" search:"type:exact;column:PNN;table:sys_list" comment:"申请号"`
	AD             string `form:"AD" search:"type:exact;column:AD;table:sys_list" comment:"申请日"`
	PD             string `form:"PD" search:"type:exact;column:PD;table:sys_list" comment:"公开日"`
	CL             string `form:"CL" search:"type:exact;column:CL;table:sys_list" comment:"简介"`
	PA             string `form:"PA" search:"type:exact;column:PA;table:sys_list" comment:"申请单位"`
	AR             string `form:"AR" search:"type:exact;column:AR;table:sys_list" comment:"地址"`
	INN            string `form:"INN" search:"type:exact;column:INN;table:sys_list" comment:"申请人"`
	SysListOrder
}

type SysListUpdateReq struct {
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

func (s SysListUpdateReq) GetPatentId() interface{} {
	return s.PatentId
}

type SysListOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_list" form:"createdAtOrder"`
}

func (m *SysListGetPageReq) GetNeedSearch() interface{} {
	return *m
}

func (s *SysListUpdateReq) GenerateList(model *models.SysList) {
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

type SysListControl struct {
	PatentId      int       `uri:"Id" comment:"主键"` // 主键
	Username      string    `json:"username" comment:"用户名"`
	Status        string    `json:"status" comment:"状态"`
	Ipaddr        string    `json:"ipaddr" comment:"ip地址"`
	LoginLocation string    `json:"loginLocation" comment:"归属地"`
	Browser       string    `json:"browser" comment:"浏览器"`
	Os            string    `json:"os" comment:"系统"`
	Platform      string    `json:"platform" comment:"固件"`
	LoginTime     time.Time `json:"loginTime" comment:"登录时间"`
	Remark        string    `json:"remark" comment:"备注"`
	Msg           string    `json:"msg" comment:"信息"`
}

type SysListGetReq struct {
	PatentId int `uri:"patent_id"`
}

func (s *SysListGetReq) GetPatentId() interface{} {
	return s.PatentId
}

// SysLoginLogDeleteReq 功能删除请求参数

type SysListDeleteReq struct {
	PatentIds []int `json:"patent_ids"`
}

func (s *SysListDeleteReq) GetPatentId() interface{} {
	return s.PatentIds
}

type SysListInsertReq struct {
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

func (s *SysListInsertReq) GenerateList(model *models.SysList) {
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

func (s *SysListInsertReq) GetPatentId() interface{} {
	return s.PatentId
}

type SysListById struct {
	dto.ObjectByPatentId
	common.ControlBy
}

func (s *SysListById) GetPatentId() interface{} {
	return s.PAId
}

func (s *SysListById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysList{}, nil
}

//type SysListByName struct {
//	dto.ObjectByPatentName
//	common.ControlBy
//}
//
//func (s *SysListByName) GetPatentTI() interface{} {
//	return s.TI
//}
//
//func (s *SysListByName) GenerateM() (common.ActiveRecord, error) {
//	return &models.SysList{}, nil
//}
