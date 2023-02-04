package dtos

import (
	"encoding/json"
	"fmt"
	"go-admin/app/admin-agent/model"
	"go-admin/app/user-agent/my_config"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
	"path"
	"strings"
)

const (
	RejectTag  = "已驳回"
	UploadTag  = "已上传"
	ProcessTag = "处理中"
	ApplyTag   = "未审核"
	CancelTag  = "已撤销"
	OKTag      = "已完成"
)

const (
	ReportTypeNovelty = "查新报告"
	ReportTypeTort    = "侵权报告"
	ReportTypeEval    = "估值报告"
)

type innerFile struct {
	FileName string `json:"FileName"`
	FilePath string `json:"FilePath"`
}

//newInnerFiles这个函数生成了文件这个结构体？（文件+路径）

func newInnerFiles(files ...string) []*innerFile {
	res := make([]*innerFile, 0, len(files))
	for _, f := range files {
		tmp := strings.Split(f, "/")
		fmt.Println(tmp)
		fn := strings.Join(strings.Split(tmp[len(tmp)-1], ".")[1:], ".")
		fmt.Println(fn)

		res = append(res, &innerFile{
			FileName: fn,
			FilePath: path.Join(my_config.CurrentPatentConfig.FileUrl, f),
		})
	}
	return res
}

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
	ReportId         int        `json:"-"`
	ReportName       string     `json:"reportName" gorm:"comment:报告名称"`
	ReportProperties Properties `json:"reportProperties" gorm:"comment:报告详情"`
	Type             string     `json:"reportType" gorm:"size:64;comment:报告类型（侵权/估值）"`
	FilesOpt         string     `json:"filesOpt" comment:"文件操作"`
	Files            []string   `json:"files" comment:"报告文件"`
}

func (s *ReportReq) Generate(model *model.Report) {
	model.ReportName = s.ReportName
	model.ReportProperties = s.ReportProperties.String()
	model.Type = s.Type
}

func (s *ReportReq) GenerateAndAddFiles(model *model.Report) {
	s.Generate(model)
	if len(model.Files) == 0 {
		innerFiles := newInnerFiles(s.Files...)
		fbs, _ := json.Marshal(innerFiles)
		model.Files = string(fbs)
	} else {
		files := make([]*innerFile, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)
		innerFiles := newInnerFiles(s.Files...)
		innerFiles = append(innerFiles, files...)
		fbs, _ := json.Marshal(innerFiles)
		model.Files = string(fbs)
	}
}

func (s *ReportReq) GenerateAndDeleteFiles(model *model.Report) {
	s.Generate(model)
	if len(model.Files) != 0 {
		files := make([]*innerFile, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)

		needToDel := make(map[string]struct{})
		for _, df := range s.Files {
			needToDel[df] = struct{}{}
		}

		slow := 0
		for _, f := range files {
			if _, ok := needToDel[f.FilePath]; !ok {
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
	switch s.FilesOpt {
	case dto.FilesAdd:
		logs = append(logs, fmt.Sprintf("上传文件%s", s.Files))
	case dto.FilesDelete:
		logs = append(logs, fmt.Sprintf("删除文件%s", s.Files))
	}
	return logs
}

type Properties map[string]interface{}

func (p Properties) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}
