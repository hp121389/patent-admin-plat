package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	internalTrace(searchTracing(req.UserId, req.Query, c.Request.RequestURI), c)

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

// SearchFullText
// @Summary 专利搜索全文
// @Description 根据查询字符串进行搜索全文
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.PatentDetail
// @Router /api/v1/user-agent/auth-search/full [post]
// @Security Bearer
func (e Search) SearchFullText(c *gin.Context) {
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

	ps, err := ic.SearchFullText(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(ps, "查询成功")
}

// GetChart
// @Summary 专利搜索统计图
// @Description 根据检索结果返回echarts option
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.ChartProfile
// @Router /api/v1/user-agent/search/charts/{aid} [POST]
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

// InsertStoredQuery
// @Summary 保存检索表达式
// @Description 为当前用户暂存检索表达式
// @Tags 专利检索
// @Param data body dto.StoredQueryReq true "表达式"
// @Success 200
// @Router /api/v1/user-agent/auth-search/queries [post]
// @Security Bearer
func (e Search) InsertStoredQuery(c *gin.Context) {
	s := service.Search{}
	req := dto.StoredQueryReq{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
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

	err = s.InsertQuery(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetID(), "创建成功")
}

// GetStoredQueryPages
// @Summary 获取暂存的检索表达式（分页）
// @Description 获取暂存的检索表达式（分页）
// @Tags 专利检索
// @Param data body dto.StoredQueryReq true "表达式"
// @Success 200
// @Router /api/v1/user-agent/auth-search/queries [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Security Bearer
func (e Search) GetStoredQueryPages(c *gin.Context) {
	s := service.Search{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.StoredQueryReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize

	if len(req.DB) == 0 {
		// todo: 设置默认数据库配置文件
		req.DB = "wgzl,syxx,fmzl"
	}

	if len(c.GetHeader("Authorization")) != 0 {
		req.UserId = user.GetUserId(c)
	}

	list := make([]models.StoredQuery, 0)
	var count int64
	err = s.GetQueryPage(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// FindStoredQueryPages
// @Summary 搜索暂存的检索表达式（分页）
// @Description 搜索暂存的检索表达式（分页）
// @Tags 专利检索
// @Param data body dto.StoredQueryFindReq true "搜索参数"
// @Success 200
// @Router /api/v1/user-agent/auth-search/queries/search [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e Search) FindStoredQueryPages(c *gin.Context) {
	s := service.Search{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	query := c.Query("query")
	req := dto.StoredQueryFindReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = query

	if len(c.GetHeader("Authorization")) != 0 {
		req.UserId = user.GetUserId(c)
	}

	list := make([]models.StoredQuery, 0)
	var count int64
	err = s.FindQueryPages(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// RemoveStoredQuery
// @Summary 删除检索表达式
// @Description 为当前用户删除检索表达式
// @Tags 专利检索
// @Success 200
// @Router /api/v1/user-agent/auth-search/queries/{query_id} [delete]
// @Security Bearer
func (e Search) RemoveStoredQuery(c *gin.Context) {
	s := service.Search{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.RemoveQuery(id)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(id, "删除成功")
}

// UpdateStoredQuery
// @Summary 更新检索表达式
// @Description 为当前用户更新检索表达式
// @Tags 专利检索
// @Param data body dto.StoredQueryReq true "表达式"
// @Success 200
// @Router /api/v1/user-agent/auth-search/queries/{query_id} [put]
// @Security Bearer
func (e Search) UpdateStoredQuery(c *gin.Context) {
	s := service.Search{}
	req := dto.StoredQueryReq{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
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

	id, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.UpdateQuery(id, &req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetID(), "更新成功")
}
