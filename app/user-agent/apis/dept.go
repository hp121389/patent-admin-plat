package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
)

type Dept struct {
	api.Api
}

// GetDeptByCurrentUser
// @Summary 获取当前用户部门
// @Description 获取当前用户部门
// @Tags 部门
// @Router /apis/v1/user-agent/dept [get]
// @Security Bearer
func (e Dept) GetDeptByCurrentUser(c *gin.Context) {
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

	userId := user.GetUserId(c)
	dept, err := s.GetDeptByUserID(userId)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(dept, "查询成功")
}
