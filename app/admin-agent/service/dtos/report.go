package dtos

import (
	"encoding/json"
	"fmt"
	"go-admin/app/admin-agent/model"
	"go-admin/app/user-agent/my_config"
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"
	"path"
	"strings"
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
	CreatedAt string   `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt string   `json:"updatedAt" gorm:"comment:最后更新时间"`
	FilesOpt  string   `json:"filesOpt" comment:"文件操作"`
	Files     []string `json:"files" comment:"报告文件"`
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

func (s *ReportGetPageReq) GenerateNoneFile(model *model.Report) {
	if s.ReportId != 0 {
		model.ReportId = s.ReportId
	}
	model.ReportName = s.ReportName
	model.RejectTag = s.RejectTag
	model.Type = s.Type
	model.ReportProperties = s.ReportProperties
	model.UpdatedAt = s.UpdatedAt

	if s.RejectTag == RejectTag {
		model.Files = "rejected clean"
	} else {
		model.Files = "null"
	}

}

type InnerFile struct {
	FileName string `json:"FileName"`
	FilePath string `json:"FilePath"`
}

//newInnerFiles这个函数生成了文件这个结构体？（文件+路径）

func newInnerFiles(files ...string) []*InnerFile {
	res := make([]*InnerFile, 0, len(files))
	for _, f := range files {
		tmp := strings.Split(f, "/")
		fmt.Println(tmp)
		fn := strings.Join(strings.Split(tmp[len(tmp)-1], ".")[1:], ".")
		fmt.Println(fn)

		res = append(res, &InnerFile{
			FileName: fn,
			FilePath: path.Join(my_config.CurrentPatentConfig.FileUrl, f),
		})
	}
	return res
}

//生成结构体并添加文件or在原有文件切片后添加文件

func (s *ReportGetPageReq) GenerateAndAddFiles(model *model.Report) {
	s.Generate(model)
	if len(model.Files) == 0 {
		innerFiles := newInnerFiles(s.Files...)
		fbs, _ := json.Marshal(innerFiles)
		// returns the JSON encoding of v 输入v,遍历v,返回byte[]
		model.Files = string(fbs)
	} else { // 长度大于0，append,不是大于1
		files := make([]*InnerFile, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)
		// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
		innerFiles := newInnerFiles(s.Files...)
		// 以下步骤相同
		innerFiles = append(innerFiles, files...)
		fbs, _ := json.Marshal(innerFiles)
		model.Files = string(fbs)
	}
}

//生成结构体,接住删除 *部分文件* 后的files

func (s *ReportGetPageReq) GenerateAndDeleteFiles(model *model.Report) {
	s.Generate(model)
	if len(model.Files) != 0 {
		files := make([]*InnerFile, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)
		// 把值存在&files里
		needToDel := make(map[string]struct{})
		//??????如何得知哪些文件需要删除？？？？needToDel？？
		for _, df := range s.Files {
			// 遍历files，把files的元素映射入map一一对应
			needToDel[df] = struct{}{}
		}

		slow := 0
		for _, f := range files {
			// 此处files是unmarshal来的，原始切片
			// 判断 key 是否在 map 里 if _, ok := map[key];
			// ok 是 false 则 slow++
			if _, ok := needToDel[f.FilePath]; !ok {
				files[slow] = f
				slow++
			}
		}
		files = files[:slow] //截取从头到slow的切片

		fbs, _ := json.Marshal(files)
		model.Files = string(fbs)
	}
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
