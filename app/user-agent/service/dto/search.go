package dto

import (
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
)

type SimpleSearchReq struct {
	dto.Pagination

	Query string `json:"query"`
	DB    string `json:"DB"`

	UserId int `json:"-"`
}

type StoredQueryReq struct {
	dto.Pagination

	QueryID int `json:"queryID"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	Query string `json:"expression"`
	DB    string `json:"DB"`

	UserId int `json:"-"`
}

type StoredQueryFindReq struct {
	dto.Pagination

	Query string `json:"query"`

	UserId int `json:"-"`
}

func (r *StoredQueryReq) Generate(m *models.StoredQuery) {
	if r.QueryID != 0 {
		m.QueryID = r.QueryID
	}
	m.Name = r.Name
	m.Desc = r.Desc
	m.Expression = r.Query
	m.DB = r.DB
	m.CreateBy = r.UserId
	m.UpdateBy = r.UserId
}

func (r *StoredQueryReq) GetID() int {
	return r.QueryID
}
