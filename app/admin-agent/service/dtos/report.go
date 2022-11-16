package dtos

import (
	"go-admin/app/admin-agent/model"
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"
	"time"
)

const (
	InfringementType = "infringement"
	ValuationType    = "valuation"
	RejectTag        = "已驳回"
	UploadTag        = "已上传"
	ProcessTag       = "处理中"
)

type ReportGetPageReq struct {
	ReportId         int    `form:"reportId" search:"type:exact;column:ReportId;table:report" comment:"报告ID"`
	ReportProperties string `form:"reportProperties" search:"type:exact;column:报告详情;table:report" comment:"报告详情""`
	ReportName       string `form:"reportName" search:"type:exact;column:reportName;table:report" comment:"报告名称"`
	Type             string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型（侵权/估值）"`
	ReportReject
	models.ControlBy
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (s *ReportGetPageReq) Generate(model *model.Report) {
	if s.ReportId != 0 {
		model.ReportId = s.ReportId
	}
	model.ReportName = s.ReportName
	model.RejectTag = s.RejectTag
	model.Type = s.Type
	model.ReportProperties = s.ReportProperties
	model.CreatedAt = s.CreatedAt
	model.UpdatedAt = s.UpdatedAt
}

type ReportReject struct {
	RejectTag string `form:"rejectTag" gorm:"size:4;comment:驳回标签"`
}

type ReportUpload struct {
	ReportId   int    `form:"reportId" search:"type:exact;column:ReportId;table:report" comment:"报告ID"`
	Type       string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型（侵权/估值）"`
	RejectTag  string `form:"rejectTag" search:"size:4;comment:驳回标签"`
	UploadFile string `form:"uploadFile" search:"comment:上传文件"`
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
