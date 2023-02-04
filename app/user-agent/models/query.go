package models

import "go-admin/common/models"

type StoredQuery struct {
	QueryID    int    `json:"queryID" gorm:"size:128;primaryKey;autoIncrement;comment:检索表达式ID(主键)"`
	Name       string `json:"name" gorm:"size:128;comment:名称"`
	Desc       string `json:"desc" gorm:"size:1024;comment:描述"`
	Expression string `json:"expression" gorm:"comment:检索表达式"`
	DB         string `json:"DB" gorm:"comment:检索数据库"`
	models.ControlBy
	models.ModelTime
}

func (e *StoredQuery) TableName() string {
	return "stored_query"
}

func (e *StoredQuery) Generate() models.ActiveRecord {
	return e
}

func (e *StoredQuery) GetId() interface{} {
	return e.QueryID
}
