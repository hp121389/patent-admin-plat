package dto

import "go-admin/common/dto"

type PatentTagGetPageReq struct {
	dto.Pagination `search:"-"`
	PatentId       int `form:"PatentId" search:"type:exact;column:PatentId;table:patent_tag" comment:"专利ID"`
	TagId          int `form:"PatentId" search:"type:exact;column:TagId;table:patent_tag" comment:"标签ID"`
	PatentTagOrder
}

type PatentTagOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:patent_tag" form:"createdAtOrder"`
}

func (m *PatentTagGetPageReq) GetNeedSearch() interface{} {
	return *m
}
