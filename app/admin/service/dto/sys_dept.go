package dto

import (
	"go-admin/app/admin/models"
)

const (
	Member        = "组员"
	NotMember     = "非组员"
	Online        = "存在"
	Offline       = "暂时下线"
	ApplyMember   = "申请成为组员"
	AlreadyMember = "成功成为组员"
	ApplyEXIT     = "申请退出团队"
	AlreadyEXIT   = "成功退出团队"

	//ApplyLeader   = "申请成为组长"
	//AlreadyLeader = "成功成为组长"
	//AlreadyLeader = "成功成为组长"
	//Leader = "组长"
)

func (s *DeptReq) GenerateDept(model *models.SysDept) {
	if s.DeptId != 0 {
		model.DeptId = s.DeptId
	}
	model.DeptName = s.DeptName
	model.DeptProperties = s.DeptProperties
}

type DeptReq struct {
	DeptId         int    `json:"deptId" gorm:"size:128;primaryKey;autoIncrement;comment:团队ID(主键)"`
	DeptName       string `json:"deptName" gorm:"comment:团队名称"`
	DeptProperties string `json:"deptProperties" gorm:"comment:团队详情"`
}
