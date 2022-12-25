package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"strconv"
)

type Dept struct {
	api.Api
}

// GetDeptList
// @Summary 列表部门
// @Description 列表部门信息
// @Tags 管理员-部门
// @Router /apis/v1/admin-agent/dept/list [get]
// @Security Bearer
func (e Dept) GetDeptList(c *gin.Context) {
	s := service.Dept{}
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

// GetDeptRelaList
// @Summary 列表部门-用户申请信息
// @Description 列表部门-用户申请信息
// @Tags 管理员-部门
// @Router /apis/v1/admin-agent/dept/relaList [get]
// @Security Bearer
func (e Dept) GetDeptRelaList(c *gin.Context) {
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
	err = s.GetDeptRelaList(&list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// GetRelaListById
// @Summary 通过部门ID列表部门-用户申请信息
// @Description 通过部门ID列表部门-用户申请信息
// @Tags 管理员-部门
// @Router /apis/v1/admin-agent/dept/relaListById/{dept_id} [get]
// @Security Bearer
func (e Dept) GetRelaListById(c *gin.Context) {
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
	var deptId int
	deptId, err = strconv.Atoi(c.Param("dept_id"))
	err = s.GetDeptRelaListById(deptId, &list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// GetDeptById
// @Summary 通过ID获取部门信息（待修改）
// @Description 通过ID获取部门信息
// @Tags 管理员-部门
// @Router /apis/v1/admin-agent/dept/{deptId} [get]
// @Security Bearer
func (e Dept) GetDeptById(c *gin.Context) {
	//s := service.Report{}
	//req := dtos.ReportGetPageReq{}
	//err := e.MakeContext(c).
	//	MakeOrm().
	//	Bind(&req).
	//	MakeService(&s.Service).
	//	Errors
	//if err != nil {
	//	e.Logger.Error(err)
	//	e.Error(500, err, err.Error())
	//	return
	//}
	//req.Type = c.Param("type")
	//list := make([]model.Report, 0)
	//
	//err = s.GetPagesByType(req.Type, &list)
	//if err != nil {
	//	e.Error(500, err, "查询失败")
	//	return
	//}
	//
	//e.OK(list, "查询成功")
}

// InsertDept
// @Summary 管理员创建团队
// @Description  用户申请报告
// @Tags 管理员-部门
// @Accept  application/json
// @Product application/json
// @Param data body dtos.DeptReq true "时间、部门名称、id"
// @Router /api/v1/admin-agent/dept [post]
// @Security Bearer
func (e Dept) InsertDept(c *gin.Context) {

	s := service.Dept{}
	reqIn := dtos.DeptReq{}
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
	reqIn.DeptStatus = dtos.Online
	err = s.InsertDept(&reqIn)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(reqIn.DeptName, "创建成功")
}

// OfflineDept
// @Summary 下线部门
// @Description 修改状态为暂时下线
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/offline/{dept_id} [put]
// @Security Bearer
func (e Dept) OfflineDept(c *gin.Context) {
	s := service.Dept{}
	req := dtos.DeptReq{}
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
	fmt.Println(req.DeptId)
	req.SetUpdateBy(user.GetUserId(c))
	req.DeptStatus = dtos.Offline

	err = s.UpdateDept(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "下线成功")
}

// ReOnlineDept
// @Summary 重新上线部门
// @Description 修改状态为存在
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/reOnline/{dept_id} [put]
// @Security Bearer
func (e Dept) ReOnlineDept(c *gin.Context) {
	s := service.Dept{}
	req := dtos.DeptReq{}
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
	req.SetUpdateBy(user.GetUserId(c))
	req.DeptStatus = dtos.Online

	err = s.UpdateDept(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.DeptStatus, "重新上线成功")
}

//--------------------------------------------------用户申请PUT-------------------------------------

// IfJoinDept
// @Summary 管理员批准用户加入团队
// @Description 管理员批准用户加入团队
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/join/{dept_id}/{user_id} [put]
// @Security Bearer
func (e Dept) IfJoinDept(c *gin.Context) {
	s := service.Dept{}
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
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	req.MemStatus = dtos.AlreadyMember
	req.ExamineStatus = dtos.OKTag
	req.MemType = dtos.Member

	err = s.UserDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "用户加入成功")
}

// RecoverJoin
// @Summary 管理员将用户移出团队
// @Description 管理员将用户移出团队
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/recoverJoin/{dept_id}/{user_id} [put]
// @Security Bearer
func (e Dept) RecoverJoin(c *gin.Context) {
	s := service.Dept{}
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
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	req.MemStatus = dtos.ApplyMember
	req.ExamineStatus = dtos.OKTag
	req.MemType = dtos.Member

	err = s.UserDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "用户移出成功")
}

// UnJoinDept
// @Summary 管理员将用户移出团队
// @Description 管理员将用户移出团队
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/unJoin/{dept_id}/{user_id} [put]
// @Security Bearer
func (e Dept) UnJoinDept(c *gin.Context) {
	s := service.Dept{}
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
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	req.MemStatus = dtos.NotMember
	req.ExamineStatus = dtos.OKTag
	req.MemType = dtos.NotMember

	err = s.UserDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "用户移出成功")
}

// JoinReject
// @Summary 管理员拒绝用户加入团队
// @Description 管理员拒绝用户加入团队
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/joinReject/{dept_id}/{user_id} [put]
// @Security Bearer
func (e Dept) JoinReject(c *gin.Context) {
	s := service.Dept{}
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
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	req.MemStatus = dtos.NotMember
	req.ExamineStatus = dtos.RejectTag
	req.MemType = dtos.NotMember

	err = s.UserDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "驳回成功")
}

// UnReject
// @Summary 管理员撤销驳回
// @Description 管理员撤销驳回
// @Tags 管理员-部门
// @Param DeptId query string false "部门ID"
// @Router /api/v1/admin-agent/dept/unReject/{dept_id}/{user_id} [put]
// @Security Bearer
func (e Dept) UnReject(c *gin.Context) {
	s := service.Dept{}
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
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	req.MemStatus = dtos.ApplyMember
	req.ExamineStatus = dtos.ProcessTag
	req.MemType = dtos.NotMember

	err = s.UserDeptRela(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "撤销驳回成功")
}
