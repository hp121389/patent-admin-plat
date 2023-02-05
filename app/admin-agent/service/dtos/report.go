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
	ReportTypeNovelty = "查新报告"
	ReportTypeTort    = "侵权报告"
	ReportTypeEval    = "估值报告"
)

type ReportPagesReq struct {
	cDto.Pagination
	Type   string `json:"type"`
	UserID int    `json:"userID"`
	Query  string `json:"query"`
}

func (s *ReportPagesReq) GetConditions() string {
	switch {
	case len(s.Type) != 0:
		return fmt.Sprintf("type = %s", s.Type)
	default:
		return ""
	}
}

type ReportReq struct {
	ReportId         int                 `json:"-"`
	ReportName       string              `json:"reportName" gorm:"comment:报告名称"`
	ReportProperties Properties          `json:"reportProperties" gorm:"comment:报告详情"`
	Type             string              `json:"reportType" gorm:"size:64;comment:报告类型（侵权/估值）"`
	FilesOpt         string              `json:"filesOpt" comment:"文件操作"`
	Files            []apis.FileResponse `json:"files" comment:"报告文件"`
}

func (s *ReportReq) Generate(model *model.Report) {
	model.ReportName = s.ReportName
	model.ReportProperties = s.ReportProperties.String()
	model.Type = s.Type
}

func (s *ReportReq) GenerateAndAddFiles(model *model.Report) {
	s.Generate(model)
	if len(model.Files) == 0 {
		fbs, _ := json.Marshal(s.Files)
		model.Files = string(fbs)
	} else {
		files := make([]apis.FileResponse, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)
		files = append(files, s.Files...)
		fbs, _ := json.Marshal(files)
		model.Files = string(fbs)
	}
}

func (s *ReportReq) GenerateAndDeleteFiles(model *model.Report) {
	s.Generate(model)
	if len(model.Files) != 0 {
		files := make([]apis.FileResponse, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)

		needToDel := make(map[string]struct{})
		for _, df := range s.Files {
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

func (s ReportReq) GenUpdateLogs() []string {
	logs := make([]string, 0)
	if len(s.Type) != 0 {
		logs = append(logs, fmt.Sprintf("修改报告类型为%s", s.Type))
	}
	if len(s.ReportName) != 0 {
		logs = append(logs, fmt.Sprintf("修改报告名称为%s; ", s.ReportName))
	}
	if len(s.ReportProperties) != 0 {
		logs = append(logs, "修改报告信息")
	}

	filenames := make([]string, 0, len(s.Files))
	for _, f := range s.Files {
		filenames = append(filenames, f.Name)
	}
	filenamesStr := strings.Join(filenames, ",")
	switch s.FilesOpt {
	case dto.FilesAdd:
		logs = append(logs, fmt.Sprintf("上传文件: %s", filenamesStr))
	case dto.FilesDelete:
		logs = append(logs, fmt.Sprintf("删除文件: %s", filenamesStr))
	}

	return logs
}

type Properties map[string]interface{}

func (p Properties) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}
