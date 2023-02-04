package dtos

import (
	"fmt"
	"go-admin/app/admin-agent/model"
	"go-admin/common/dto"
)

const (
	TicketTypeReport = "report"
	TicketTypeCommon = "common"
)

const (
	TicketStatusOpen     = "open"
	TicketStatusClosed   = "closed"
	TicketStatusFinished = "finished"
)

type TicketPagesReq struct {
	dto.Pagination
	Type   string `json:"type"`
	Status string `json:"status"`
	UserID int    `json:"userID"`
	Query  string `json:"query"`
}

type TicketListReq struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	UserID int    `json:"userID"`
	Query  string `json:"query"`
}

type TicketRelObj interface {
	GenUpdateLogs() []string
}

type TicketReq[T TicketRelObj] struct {
	TicketDBReq
	RelObj T `json:"relObj"`
}

func (tr *TicketReq[T]) GenOptLogsWhenUpdate() {
	msg := ""
	index := 1
	if len(tr.Name) != 0 {
		msg += fmt.Sprintf("%d. 修改工单名为: %s", index, tr.Name)
		index++
	}
	if len(tr.Type) != 0 {
		msg += fmt.Sprintf("%d. 修改工单类型为: %s", index, tr.Type)
		index++
	}
	if len(tr.Properties) != 0 {
		msg += fmt.Sprintf("%d. 修改工单信息", index)
		index++
	}
	for _, relLog := range tr.RelObj.GenUpdateLogs() {
		msg += fmt.Sprintf("%d. %s", index, relLog)
		index++
	}
	tr.OptMsg = msg
}

func NewReportTicketReq() TicketReq[ReportReq] {
	tr := TicketReq[ReportReq]{}
	return tr
}

type TicketDBReq struct {
	RelaID     int        `json:"relaID"`
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
	UserID     int        `json:"userID"`
	OptMsg     string     `json:"optMsg"`
}

func NewTicketDBReq(uid int, optMsg string) *TicketDBReq {
	return &TicketDBReq{
		UserID: uid,
		OptMsg: optMsg,
	}
}

func (r *TicketDBReq) Generate(model *model.Ticket) {
	if r.RelaID != 0 {
		model.RelaID = r.RelaID
	}
	if len(r.Name) != 0 {
		model.Name = r.Name
	}
	if len(r.Properties) != 0 {
		model.Properties = r.Properties.String()
	}
	if len(r.Type) != 0 {
		model.Type = r.Type
	}
}

func (r *TicketDBReq) GenOptLogsWhenUpdate() {
	msg := ""
	index := 0
	if len(r.Name) != 0 {
		msg += fmt.Sprintf("%d. 修改工单名为: %s", index, r.Name)
		index++
	}
	if len(r.Type) != 0 {
		msg += fmt.Sprintf("%d. 修改工单类型为: %s", index, r.Type)
		index++
	}
	if len(r.Properties) != 0 {
		msg += fmt.Sprintf("%d. 修改工单信息", index)
		index++
	}
	r.OptMsg = msg
}
