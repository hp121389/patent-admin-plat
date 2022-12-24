package dtos

import (
	"go-admin/app/admin-agent/model"
	"go-admin/common/models"
)

type ReportRelaReq struct {
	Id       int `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	PatentId int `json:"patentId" gorm:"size:128;comment:专利ID"`
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	UserId   int `json:"userId" gorm:"size:128;comment:用户ID"`
	models.ControlBy
	CreatedAt string `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt string `json:"updatedAt" gorm:"comment:最后更新时间"`
	Type      string `form:"Type" search:"type:exact;column:Type;table:report" comment:"报告类型"`
}

func (s *ReportRelaReq) GenerateRela(model *model.ReportRelation) {
	if s.ReportId != 0 {
		model.ReportId = s.ReportId
	}
	if s.PatentId != 0 {
		model.PatentId = s.PatentId
	}
	model.UserId = s.UserId
	model.Type = s.Type
	model.CreatedAt = s.CreatedAt
	model.UpdatedAt = s.UpdatedAt
	model.CreateBy = s.CreateBy
	model.UpdateBy = s.UpdateBy
}
