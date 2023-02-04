package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
)

type TraceReq struct {
	UserID  int    `json:"userID"`
	Action  string `json:"action"`
	Desc    string `json:"desc"`
	Request string `json:"request"`
}

func (r *TraceReq) Generate(m *models.TraceLog) {
	m.CreateBy = r.UserID
	m.Action = r.Action
	m.Desc = r.Desc
	m.Request = r.Request
}

type TracePageReq struct {
	dto.Pagination

	UserID int    `json:"userID"`
	Action string `json:"action"`
}
