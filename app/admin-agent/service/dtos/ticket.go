package dtos

import (
	"encoding/json"
	"fmt"
	"go-admin/app/admin-agent/model"
	"go-admin/app/other/apis"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
	"strings"
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
	cDto.Pagination
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
	switch tr.FilesOpt {
	case dto.FilesAdd:
		msg += fmt.Sprintf("%d. 上传工单文件%s", index, tr.Files)
		index++
	case dto.FilesDelete:
		msg += fmt.Sprintf("%d. 删除工单文件%s", index, tr.Files)
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
	RelaID     int                 `json:"relaID"`
	Name       string              `json:"name"`
	Properties Properties          `json:"properties"`
	Type       string              `json:"type"`
	UserID     int                 `json:"userID"`
	OptMsg     string              `json:"optMsg"`
	FilesOpt   string              `json:"filesOpt" comment:"文件操作"`
	Files      []apis.FileResponse `json:"files" comment:"报告文件"`
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

func (r *TicketDBReq) GenerateAndAddFiles(model *model.Ticket) {
	r.Generate(model)
	if len(model.Files) == 0 {
		fbs, _ := json.Marshal(r.Files)
		model.Files = string(fbs)
	} else {
		files := make([]apis.FileResponse, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)
		files = append(files, r.Files...)
		fbs, _ := json.Marshal(files)
		model.Files = string(fbs)
	}
}

func (r *TicketDBReq) GenerateAndDeleteFiles(model *model.Ticket) {
	r.Generate(model)
	if len(model.Files) != 0 {
		files := make([]apis.FileResponse, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)

		needToDel := make(map[string]struct{})
		for _, df := range r.Files {
			needToDel[df.FullPath] = struct{}{}
		}

		slow := 0
		for _, f := range files {
			if _, ok := needToDel[f.FullPath]; !ok {
				files[slow] = f
				slow++
			}
		}
		files = files[:slow]
		fbs, _ := json.Marshal(files)
		model.Files = string(fbs)
	}
}

func (r *TicketDBReq) GenOptLogsWhenUpdate() {
	msg := ""
	index := 1
	if len(r.Name) != 0 {
		msg += fmt.Sprintf("%d. 修改工单名为: %s; ", index, r.Name)
		index++
	}
	if len(r.Type) != 0 {
		msg += fmt.Sprintf("%d. 修改工单类型为: %s; ", index, r.Type)
		index++
	}
	if len(r.Properties) != 0 {
		msg += fmt.Sprintf("%d. 修改工单信息; ", index)
		index++
	}

	filenames := make([]string, 0, len(r.Files))
	for _, f := range r.Files {
		filenames = append(filenames, f.Name)
	}
	filenamesStr := strings.Join(filenames, ",")
	switch r.FilesOpt {
	case dto.FilesAdd:
		msg += fmt.Sprintf("%d. 上传工单文件: %s; ", index, filenamesStr)
		index++
	case dto.FilesDelete:
		msg += fmt.Sprintf("%d. 删除工单文件: %s; ", index, filenamesStr)
		index++
	}
	r.OptMsg = msg
}
