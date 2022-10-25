package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

type Search struct {
	api.Api
}

// Search
// @Summary 专利搜索
// @Description 根据查询字符串进行搜索（可传入逻辑表达式或简单字符串）
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.SwagSearchListResp
// @Router /api/v1/search [post]
// @Security Bearer
func (e Search) Search(c *gin.Context) {
	ic := service.GetCurrentInnojoy()
	s := service.Search{}
	req := dto.SimpleSearchReq{}

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

	if len(req.DB) == 0 {
		// todo: 设置默认数据库配置文件
		req.DB = "wgzl,syxx,fmzl"
	}

	if len(c.GetHeader("Authorization")) != 0 {
		req.UserId = user.GetUserId(c)
	}

	ps, err := ic.Search(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(ps, req.PageSize, req.PageIndex, len(ps), "查询成功")
}
