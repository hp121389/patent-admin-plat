package models

import (
	"go-admin/common/models"
)

type Patent struct {
	//models.Model        //就是自增id
	PatentId         int    `json:"PatentId" gorm:"size:128;primaryKey;autoIncrement;comment:专利ID(主键)"`
	PNM              string `json:"PNM" gorm:"size:128;comment:申请号"`
	PatentProperties string `json:"patentProperties" gorm:"comment:专利详情"`
	models.ControlBy
	//嵌入结构体：先写好models然后嵌入，等效于models本体
}

func (Patent) TableName() string {
	return "patent"
}

func (e *Patent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Patent) GetId() interface{} {
	return e.PatentId
}

//user-patent

type UserPatent struct {
	models.Model
	PatentId int    `gorm:"foreignKey:PatentId;comment:PatentId" json:"PatentId" `
	UserId   int    `gorm:"comment:用户ID"  json:"UserId"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号"`
	Type     string `json:"Type" gorm:"size:64;comment:关系类型（关注/认领）"`
	models.ControlBy
}

func (e *UserPatent) TableName() string {
	return "user_patent"
}

//patent-tag

type PatentTag struct {
	models.Model
	PatentId int `gorm:"foreignKey:PatentId;comment:专利Id" json:"PatentId" `
	TagId    int `gorm:"comment:标签ID"  json:"TagId"`
	models.ControlBy
}

func (e *PatentTag) TableName() string {
	return "patent_tag"
}

//patent-package

type PatentPackage struct {
	models.Model
	PatentId  int    `gorm:"foreignKey:PatentId;comment:专利Id" json:"PatentId" `
	PackageId int    `gorm:"comment:专利包ID"  json:"PackageId"`
	PNM       string `json:"PNM" gorm:"size:128;comment:申请号"`
	models.ControlBy
}

func (e *PatentPackage) TableName() string {
	return "patent_package"
}
