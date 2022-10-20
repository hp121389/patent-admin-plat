package dto

import (
	"go-admin/app/patent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TagInsertReq struct {
	TagId   int    `json:"tagId" comment:"标签ID"`   //标签ID
	TagName string `json:"tagName" comment:"标签名;"` //标签名称
	Desc    string `json:"desc"    comment:"标签描述"` //标签描述
	common.ControlBy
}

type TagUpdateReq struct {
	TagId   int    `json:"tagId" comment:"标签ID"`   //标签ID
	TagName string `json:"tagName" comment:"标签名;"` //标签名称
	Desc    string `json:"desc"    comment:"标签描述"` //标签描述
	common.ControlBy
}

type TagById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *TagInsertReq) Generate(model *models.Tag) {
	if s.TagId != 0 {
		model.TagId = s.TagId
	}
	model.TagName = s.TagName
	model.Desc = s.TagName
}

func (s *TagUpdateReq) Generate(model *models.Tag) {
	if s.TagId != 0 {
		model.TagId = s.TagId
	}
	model.TagName = s.TagName
	model.Desc = s.TagName
}

func (s *TagInsertReq) GetId() interface{} {
	return s.TagId
}

func (s *TagUpdateReq) GetId() interface{} {
	return s.TagId
}

type TagGetReq struct {
	Id int `uri:"id"`
}

func (s *TagGetReq) GetId() interface{} {
	return s.Id
}

type TagsByIdsForRelationshipPatents struct {
	dto.ObjectOfTagId
}

func (s *TagsByIdsForRelationshipPatents) GetTagId() []int {
	s.TagIds = append(s.TagIds, s.TagId)
	return s.TagIds
}

func (s *TagsByIdsForRelationshipPatents) GetNeedSearch() interface{} {
	return *s
}
