package dtos

import (
	"go-admin/common/models"
	"mime/multipart"
)

type PatentReport struct {
	Id       int `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	PatentId int `json:"patentId" gorm:"size:128;comment:专利ID"`
	ReportId int `json:"reportId" gorm:"size:128;comment:报告ID"`
	models.ControlBy
	file multipart.FileHeader
}
