package dtos

import (
	"go-admin/app/admin-agent/model"
)

type ReportRelaReq struct {
	TicketId int `json:"ticketId" gorm:"size:128;comment:工单ID"`
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	UserId   int `json:"userId" gorm:"size:128;comment:用户ID"`
}

func (s *ReportRelaReq) Generate(model *model.ReportRelation) {
	if s.ReportId != 0 {
		model.ReportId = s.ReportId
	}
	if s.TicketId != 0 {
		model.TicketId = s.TicketId
	}
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
}
