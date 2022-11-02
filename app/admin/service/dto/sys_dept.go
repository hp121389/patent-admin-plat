package dto

import (
	"go-admin/app/admin/models"
	common "go-admin/common/models"
)

type SysDeptInsertReq struct {
	DeptId   int    `uri:"sysId" comment:"部门Id"`     //部门Id
	DeptName string `json:"deptName" comment:"部门名称"` //部门名称
	DeptDesc string `json:"deptDesc" comment:"部门描述"` //部门描述
	common.ControlBy
}

func (s *SysDeptInsertReq) Generate(model *models.SysDept) {
	if s.DeptId != 0 {
		model.DeptId = s.DeptId
	}
	model.DeptName = s.DeptName
	model.DeptDesc = s.DeptDesc
}
