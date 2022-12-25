package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	serviceAdmin "go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/service"
	"strconv"
)

type Dept struct {
	api.Api
}

// GetRelaListByUserId
// @Summary 用户ID查看自己的部门
// @Description 用户ID查看自己的部门
// @Tags 用户-部门
// @Router /apis/v1/user-agent/dept/relaList [get]
// @Security Bearer
func (e Dept) GetRelaListByUserId(c *gin.Context) {
	s := service.Dept{}
	req := dtos.DeptRelaReq{}
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
	list := make([]model.DeptRelation, 0)

	err = s.GetDeptRelaListByUser(user.GetUserId(c), &list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "查询成功")
}

// RelaListOfUserId
// @Summary 用户在该部门的状态
// @Description 用户在该部门的状态
// @Tags 用户-部门
// @Router /apis/v1/user-agent/dept/relaListOfUserId/{dept_id} [get]
// @Security Bearer
func (e Dept) RelaListOfUserId(c *gin.Context) {
	s := service.Dept{}
	req := dtos.DeptRelaReq{}
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
	list := make([]model.DeptRelation, 0)
	req.DeptId, err = strconv.Atoi(c.Param("dept_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.GetDeptUser(req.DeptId, user.GetUserId(c), &list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "查询成功")
}

// GetDeptList
// @Summary 列表部门
// @Description 列表部门信息
// @Tags 用户-部门
// @Router /apis/v1/user-agent/dept/list [get]
// @Security Bearer
func (e Dept) GetDeptList(c *gin.Context) {
	s := serviceAdmin.Dept{}
	req := dtos.DeptReq{}
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
	list := make([]model.Dept, 0)
	err = s.GetDeptList(&list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// UserJoinDept
// @Summary 用户申请加入团队
// @Description  用户申请加入团队
// @Tags 用户-部门
// @Accept  application/json
// @Product application/json
// @Param data body dtos.DeptRelaReq true "用户Id、部门Id"
// @Router /api/v1/user-agent/dept/joinApply/{deptId} [post]
// @Security Bearer
func (e Dept) UserJoinDept(c *gin.Context) {

	s := service.Dept{}
	reqIn := dtos.DeptRelaReq{}
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
	reqIn.MemStatus = dtos.ApplyMember
	reqIn.ExamineStatus = dtos.ApplyTag
	reqIn.MemType = dtos.NotMember
	reqIn.DeptId, err = strconv.Atoi(c.Param("dept_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	reqIn.UserId = user.GetUserId(c)
	reqIn.CreateBy = user.GetUserId(c)
	err = s.InsertDeptRela(&reqIn)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(reqIn.DeptId, "申请成功")
}

// UserCancelJoinDept
// @Summary 用户取消申请加入
// @Description 用户取消申请加入
// @Tags 用户-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/user-agent/dept/joinCancel/{dept_id} [put]
// @Security Bearer
func (e Dept) UserCancelJoinDept(c *gin.Context) {
	s := serviceAdmin.Dept{}
	req := dtos.DeptRelaReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.DeptId, err = strconv.Atoi(c.Param("dept_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	req.ExamineStatus = dtos.CancelTag
	req.MemStatus = dtos.NotMember
	req.MemType = dtos.NotMember
	err = s.UpdateDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.ExamineStatus, "取消加入成功")
}

// ReJoinDept
// @Summary 用户重新申请加入
// @Description 用户重新申请加入
// @Tags 用户-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/user-agent/dept/reJoin/{dept_id} [put]
// @Security Bearer
func (e Dept) ReJoinDept(c *gin.Context) {
	s := serviceAdmin.Dept{}
	req := dtos.DeptRelaReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.DeptId, err = strconv.Atoi(c.Param("dept_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	req.ExamineStatus = dtos.ApplyTag
	req.MemStatus = dtos.ApplyMember
	req.MemType = dtos.NotMember
	err = s.UpdateDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.ExamineStatus, "重新加入成功")
}

// ExitDept
// @Summary 用户申请退出团队
// @Description 用户申请退出团队（前提是，已经是组员）
// @Tags 用户-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/user-agent/dept/exit/{dept_id} [put]
// @Security Bearer
func (e Dept) ExitDept(c *gin.Context) {
	s := serviceAdmin.Dept{}
	req := dtos.DeptRelaReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.DeptId, err = strconv.Atoi(c.Param("dept_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	req.ExamineStatus = dtos.ApplyTag
	req.MemStatus = dtos.ApplyEXIT
	err = s.UpdateDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.MemStatus, "申请退出成功")
}
