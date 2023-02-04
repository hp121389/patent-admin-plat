package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

type Trace struct {
	api.Api
}

var traceAPI Trace

func init() {
	traceAPI = Trace{}
}

// SelectTraceLog
// @Summary 检索用户行为Tracing
// @Description 检索用户行为Tracing
// @Tags 用户行为Tracing
// @Param data body dto.TracePageReq true "用户数据"
// @Success 200 {object} []models.TraceLog
// @Router /api/v1/user-agent/tracing/logs [get]
// @Security Bearer
func (e Trace) SelectTraceLog(c *gin.Context) {
	s := service.Tracer{}
	req := dto.TracePageReq{}

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

	if len(c.GetHeader("Authorization")) != 0 {
		req.UserID = user.GetUserId(c)
	}

	list := make([]models.TraceLog, 0)
	var count int64
	err = s.SelectTraceLog(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.PageIndex, req.PageSize, "查询成功")
}

func internalTrace(req *dto.TraceReq, c *gin.Context) {
	s := service.Tracer{}
	err := traceAPI.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		traceAPI.Logger.Error(err)
		traceAPI.Error(500, err, err.Error())
		return
	}

	err = s.Trace(req)
	if err != nil {
		traceAPI.Logger.Error(err)
		traceAPI.Error(500, err, err.Error())
		return
	}
}

func searchTracing(userID int, query string, route string) *dto.TraceReq {
	return &dto.TraceReq{
		UserID:  userID,
		Action:  "Search",
		Desc:    fmt.Sprintf("查询操作，表达式：%s", query),
		Request: fmt.Sprintf("请求URI：%s", route),
	}
}
