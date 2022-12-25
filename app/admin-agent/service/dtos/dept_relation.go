package dtos

import (
	"go-admin/app/admin-agent/model"
	"go-admin/common/models"
)

type DeptRelaReq struct {
	Id     int `json:"Id" gorm:"primaryKey;autoIncrement;comment:主键"`
	UserId int `json:"userId" gorm:"size:128;comment:成员ID"`
	DeptId int `json:"deptId" gorm:"size:128;comment:部门ID"`

	MemType       string `form:"MemType" search:"type:exact;column:MemType;table:dept" comment:"成员类型"`
	CreatedAt     string `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     string `json:"updatedAt" gorm:"comment:最后更新时间"`
	MemStatus     string `json:"memStatus" gorm:"comment:成员状态"`
	ExamineStatus string `json:"ExamineStatus" gorm:"comment:审核状态"`
	models.ControlBy
}

func (s *DeptRelaReq) GenerateRela(model *model.DeptRelation) {
	if s.DeptId != 0 {
		model.DeptId = s.DeptId
	}

	model.UserId = s.UserId
	model.MemType = s.MemType
	model.ExamineStatus = s.ExamineStatus
	model.MemStatus = s.MemStatus
	model.CreatedAt = s.CreatedAt
	model.UpdatedAt = s.UpdatedAt
	model.CreateBy = s.CreateBy
	model.UpdateBy = s.UpdateBy
}
