package models

import (
	"go-admin/common/models"
)

type Patent2 struct {
	PatentId int    `form:"patentId" search:"type:exact;column:PatentId;table:patent" comment:"专利ID"`
	TI       string `form:"TI" search:"type:exact;column:TI;table:patent" comment:"专利名"`
	PNM      string `form:"PNM" search:"type:exact;column:PNN;table:patent" comment:"申请号"`
	AD       string `form:"AD" search:"type:exact;column:AD;table:patent" comment:"申请日"`
	PD       string `form:"PD" search:"type:exact;column:PD;table:patent" comment:"公开日"`
	CL       string `form:"CL" search:"type:exact;column:CL;table:patent" comment:"简介"`
	PA       string `form:"PA" search:"type:exact;column:PA;table:patent" comment:"申请单位"`
	AR       string `form:"AR" search:"type:exact;column:AR;table:patent" comment:"地址"`
	PINN     string `form:"PINN" search:"type:exact;column:PINN;table:patent" comment:"申请人"`
	CLS      string `json:"CLS" gorm:"size:128;comment:法律状态"`
	CLAIMS   string `json:"Claims" gorm:"size:128;comment:权利申请书"`
	models.ControlBy
}

func (Patent2) TableName() string {
	return "patent"
}

func (e *Patent2) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Patent2) GetId() interface{} {
	return e.PatentId
}
