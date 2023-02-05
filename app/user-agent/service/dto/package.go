package dto

import (
	"encoding/json"
	"go-admin/app/other/apis"
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

const (
	FilesAdd    = "add"
	FilesDelete = "del"
)

type PackageListReq struct {
	UserId int `json:"-" form:"desc" search:"type:order;column:created_at;table:package"`
}

type PackageFindReq struct {
	UserId int    `json:"-" form:"desc" search:"type:order;column:created_at;table:package"`
	Query  string `json:"query"`
}

type PackageInsertReq struct {
	PackageId   int    `json:"packageId" comment:"专利包ID"` // 专利包ID
	PackageName string `json:"packageName" comment:"专利包名" vd:"len($)>0"`
	Desc        string `json:"desc" comment:"描述"`
	common.ControlBy
}

func (s *PackageInsertReq) Generate(model *models.Package) {
	if s.PackageId != 0 {
		model.PackageId = s.PackageId
	}
	model.PackageName = s.PackageName
	model.Desc = s.Desc
	model.ControlBy = s.ControlBy
}

func (s *PackageInsertReq) GetId() interface{} {
	return s.PackageId
}

type PackageUpdateReq struct {
	PackageId   int                 `json:"packageId" comment:"专利包ID"` // 专利包ID
	PackageName string              `json:"packageName" comment:"专利包名"`
	Desc        string              `json:"desc" comment:"描述"`
	FilesOpt    string              `json:"filesOpt" comment:"文件操作"`
	Files       []apis.FileResponse `json:"files" comment:"专利包附件"`
	common.ControlBy
}

func (s *PackageUpdateReq) Generate(model *models.Package) {
	if s.PackageId != 0 {
		model.PackageId = s.PackageId
	}
	model.PackageName = s.PackageName
	model.Desc = s.Desc
}

func (s *PackageUpdateReq) GenerateAndAddFiles(model *models.Package) {
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

func (s *PackageUpdateReq) GenerateAndDeleteFiles(model *models.Package) {
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

func (s *PackageUpdateReq) GetId() interface{} {
	return s.PackageId
}

type PackageById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *PackageById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}
