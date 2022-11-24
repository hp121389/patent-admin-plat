package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"strconv"
)

type Search struct {
	api.Api
}

// AuthSearch
// @Summary 专利搜索
// @Description 根据查询字符串进行搜索（已登陆）
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.SwagSearchListResp
// @Router /api/v1/user-agent/auth-search [post]
// @Security Bearer
func (e Search) AuthSearch(c *gin.Context) {
	ic := service.GetCurrentInnojoy()
	s := service.UserPatent{}
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

	upReq := &dto.UserPatentObject{UserId: req.UserId}
	relatedPatents := make([]models.UserPatent, 0)
	err = s.GetAllRelatedPatentsByUserId(upReq, &relatedPatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ps, err := ic.Search(&req, relatedPatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(ps, req.PageSize, req.PageIndex, len(ps), "查询成功")
}

// Search
// @Summary 专利搜索
// @Description 根据查询字符串进行搜索（未登录）
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.SwagSearchListResp
// @Router /api/v1/user-agent/search [post]
// @Security Bearer
func (e Search) Search(c *gin.Context) {
	ic := service.GetCurrentInnojoy()
	req := dto.SimpleSearchReq{}

	err := e.MakeContext(c).
		Bind(&req).
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

	ps, err := ic.Search(&req, nil)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(ps, req.PageSize, req.PageIndex, len(ps), "查询成功")
}

// GetChart
// @Summary 专利搜索统计图
// @Description 根据检索结果返回echarts option
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.ChartProfile
// @Router /api/v1/user-agent/charts/{aid} [POST]
// @Security Bearer
func (e Search) GetChart(c *gin.Context) {
	ic := service.GetCurrentInnojoy()
	req := dto.SimpleSearchReq{}

	err := e.MakeContext(c).
		Bind(&req).
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

	pathParam := c.Param("aid")
	aid, err := strconv.Atoi(pathParam)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	cp, err := ic.GetChart(aid, &req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(cp, "获取成功")
}
