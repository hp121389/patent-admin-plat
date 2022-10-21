package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PatentTagGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int `uri:"patent_id"`
	TagId          int `form:"TagId" search:"type:exact;column:TagId;table:patent_tag" comment:"标签ID"`
	PatentTagOrder
}

type TagPageGetReq struct {
	dto.Pagination `search:"-"`
	PatentId       int `form:"PatentId" search:"type:exact;column:TagId;table:patent_tag" comment:"专利ID"`
	TagId          int `uri:"tag_id"`
	PatentTagOrder
}

func (m *TagPageGetReq) GetPatentId() interface{} {
	return m.PatentId
}

func (m *TagPageGetReq) GetTagId() interface{} {
	return m.TagId
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
	TagId    int `json:"TagId" gorm:"size:128;comment:标签ID"`
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`
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

type TagUpdateReqByPatent struct {
	TagId    int `json:"TagId" gorm:"size:128;comment:标签ID"`
	PatentId int `uri:"patent_id"`
	common.ControlBy
}

type PatentUpdateReqByTag struct {
	TagId    int `uri:"tag_id"`
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}

type PatentTagObject struct {
	TagId    int `uri:"tag_id"`
	PatentId int `uri:"patent_id"`
	common.ControlBy
}

func (d *PatentTagObject) GetPatentId() interface{} {
	return d.PatentId
}

func (d *PatentTagObject) GetTagId() interface{} {
	return d.TagId
}
