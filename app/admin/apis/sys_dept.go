package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"strconv"
)

type Dept struct {
	api.Api
}

// GetDeptList
// @Summary 列表部门
// @Description 列表部门信息
// @Tags 部门
// @Router /apis/v1/admin-agent/dept [get]
// @Security Bearer
func (e Dept) GetDeptList(c *gin.Context) {
	s := service.SysDept{}
	req := dto.DeptReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]models.SysDept, 0)
	err = s.GetDeptList(&list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// CreateDept
// @Summary 管理员创建团队
// @Description  管理员创建团队
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body dto.DeptReq true "部门详情"
// @Router /api/v1/admin-agent/dept [post]
// @Security Bearer
func (e Dept) CreateDept(c *gin.Context) {

	s := service.SysDept{}
	reqIn := dto.DeptReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&reqIn, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Insert(&reqIn)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(reqIn.DeptName, "创建成功")
}

// UpdateDept
// @Summary 管理员更新团队
// @Description  管理员更新团队
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body dto.DeptReq true "部门详情"
// @Router /api/v1/admin-agent/dept/{id} [put]
// @Security Bearer
func (e Dept) UpdateDept(c *gin.Context) {
	s := service.SysDept{}
	reqIn := dto.DeptReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&reqIn, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reqIn.DeptId = id

	err = s.Update(&reqIn)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(reqIn.DeptName, "创建成功")
}

// RemoveDept
// @Summary 管理员删除团队
// @Description  管理员删除团队
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/dept/{id} [delete]
// @Security Bearer
func (e Dept) RemoveDept(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.Remove(id)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "删除成功")
}
