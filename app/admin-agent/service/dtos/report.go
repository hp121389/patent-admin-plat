package dtos

import (
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"
)

const (
	InfringementType = "infringement"
	ValuationType    = "valuation"
	RejectTag        = "reject"
	UploadTag        = "upload"
)

type ReportGetPageReq struct {
	ReportId         int    `form:"reportId" search:"type:exact;column:ReportId;table:report" comment:"报告ID"`
	ReportProperties string `form:"reportProperties" search:"type:exact;column:报告详情;table:report" comment:"报告详情""`
	ReportName       string `form:"reportName" search:"type:exact;column:reportName;table:report" comment:"报告名称"`
	Type             string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型（侵权/估值）"`
	ReportReject
	models.ControlBy
}

type ReportReject struct {
	RejectTag string `form:"rejectTag" gorm:"size:4;comment:驳回标签"`
}

type ReportUpload struct {
	ReportId   int    `form:"reportId" search:"type:exact;column:ReportId;table:report" comment:"报告ID"`
	Type       string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型（侵权/估值）"`
	RejectTag  string `form:"rejectTag" search:"size:4;comment:驳回标签"`
	UploadFile string `form:"rejectTag" search:"comment:驳回标签"`
	models.ControlBy
}

type ReportById struct {
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	common.ControlBy
}

type PatentById struct {
	PatentId int `json:"reportId" gorm:"size:128;comment:报告ID"`
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
