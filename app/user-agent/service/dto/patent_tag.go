package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

//patent-tag

type PatentTagGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int `form:"patentId" search:"type:exact;column:TagId;table:patent_tag" comment:"专利ID"`
	TagId          int `form:"tagId" search:"type:exact;column:TagId;table:patent_tag" comment:"标签ID"`
	PatentTagOrder
	common.ControlBy
}

type PatentTagOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:patent_tag" form:"CreatedAtOrder"`
}

func (m *PatentTagGetPageReq) GetNeedSearch() interface{} {
	return *m
}

func (m *PatentTagGetPageReq) GetPatentId() interface{} {
	return m.PatentId
}

func (m *PatentTagGetPageReq) GetTagId() interface{} {
	return m.TagId
}

type PatentTagInsertReq struct {
	TagId    int `json:"tagId" gorm:"size:128;comment:标签ID"`
	PatentId int `json:"patentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}

func (d *PatentTagInsertReq) GeneratePatentTag(g *models.PatentTag) {
	g.PatentId = d.PatentId
	g.TagId = d.TagId
}

func (d *PatentTagInsertReq) GetPatentId() interface{} {
	return d.PatentId
}

func (d *PatentTagInsertReq) GetTagId() interface{} {
	return d.TagId
}

