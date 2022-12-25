package dtos

import (
	"go-admin/app/admin-agent/model"
	"go-admin/common/models"
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

func (s *DeptReq) GenerateDept(model *model.Dept) {
	if s.DeptId != 0 {
		model.DeptId = s.DeptId
	}
	model.DeptName = s.DeptName
	model.CreatedAt = UpdateTime()
	model.DeptProperties = s.DeptProperties
	model.ResearchInterest = s.ResearchInterest
	model.DeptStatus = s.DeptStatus
}

type DeptReq struct {
	DeptId           int    `json:"deptId" gorm:"size:128;primaryKey;autoIncrement;comment:团队ID(主键)"`
	DeptStatus       string `json:"deptStatus" gorm:"comment:团队状态"`
	DeptName         string `json:"deptName" gorm:"comment:团队名称"`
	DeptProperties   string `json:"deptProperties" gorm:"comment:团队详情"`
	ResearchInterest string `json:"researchInterest" gorm:"comment:研究方向"`
	CreatedAt        string `json:"CreatedAt" gorm:"comment:创建时间"`
	UpdatedAt        string `json:"UpdatedAt" gorm:"comment:最后更新时间"`
	models.ControlBy
}
