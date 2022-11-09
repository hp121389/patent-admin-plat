package dto

import (
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"
)

const (
	InfringementType = "侵权报告"
	ValuationType    = "估值报告"
)

type ReportGetPageReq struct {
	ReportId         int    `form:"reportId" search:"type:exact;column:ReportId;table:report" comment:"报告ID"`
	ReportProperties string `form:"reportProperties" search:"type:exact;column:报告详情;table:report" comment:"报告详情""`
	ReportName       string `form:"reportName" search:"type:exact;column:reportName;table:report" comment:"报告名称"`
	Type             string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型（侵权/估值）"`
	models.ControlBy
}

type ReportById struct {
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	common.ControlBy
}
