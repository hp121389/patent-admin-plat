package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	serviceAdmin "go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/models"
	serviceUser "go-admin/app/user-agent/service"
	"strconv"
)

type Report struct {
	api.Api
}

//----------------------------------------GET report----------------------------------------

// UserGetReportByType
// @Summary 用户按类型查看报告
// @Description 用户查看自己认领的专利申请的type类报告
// @Tags 用户-报告
// @Param type query string false "报告类型"
// @Router /api/v1/user-agent/report/{type} [get]
// @Security Bearer
func (e Report) UserGetReportByType(c *gin.Context) {

	sRela := serviceAdmin.Report{}
	reqRe := dtos.ReportRelaReq{}
	req := dtos.ReportIds{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&sRela.Service). //注意绑定的service到底是谁
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	reqRe.UserId = user.GetUserId(c)
	reqRe.Type = c.Param("type")
	list := make([]model.ReportRelation, 0)

	err = sRela.GetReportIdsByType(&reqRe, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	listRepo := make([]model.Report, 0)
	req.ReportIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req.ReportIds[i] = list[i].ReportId
	}

	err = sRela.GetReportListByIds(&req, &listRepo)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(listRepo, "查询成功")
}

// UserGetReports
// @Summary 用户查看所有报告
// @Description 用户查看自己认领的专利申请的所有报告
// @Tags 用户-报告
// @Router /api/v1/user-agent/report [get]
// @Security Bearer
func (e Report) UserGetReports(c *gin.Context) {

	sRela := serviceAdmin.Report{}
	reqRe := dtos.ReportRelaReq{}
	req := dtos.ReportIds{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&sRela.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	reqRe.UserId = user.GetUserId(c)
	list := make([]model.ReportRelation, 0)
	err = sRela.GetReportIdsByUserId(reqRe.UserId, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	listRepo := make([]model.Report, 0)
	req.ReportIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req.ReportIds[i] = list[i].ReportId
	}

	err = sRela.GetReportListByIds(&req, &listRepo)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(listRepo, "查询成功")
}

//----------------------------------------Insert report----------------------------------------

// InsertReport
// @Summary 用户申请报告
// @Description  用户申请报告
// @Tags 用户-报告
// @Accept  application/json
// @Product application/json
// @Param data body dtos.ReportInsertGetReq true "时间、报告类型、专利id"
// @Router /api/v1/user-agent/report [post]
// @Security Bearer
func (e Report) InsertReport(c *gin.Context) {

	s := serviceUser.Report{}
	req := dtos.ReportGetPageReq{}
	reqIn := dtos.ReportInsertGetReq{}
	reqRela := dtos.ReportRelaReq{}

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

	var f *model.Report

	req.CreateBy = user.GetUserId(c)
	req.RejectTag = dtos.ApplyTag
	req.Type = reqIn.Type
	req.CreatedAt = reqIn.CreatedAt

	err, f = s.InsertReport(&req, req.Type)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	reqRela.UserId = user.GetUserId(c)
	reqRela.ReportId = f.ReportId
	reqRela.Type = f.Type
	reqRela.PatentId = reqIn.PatentId
	reqRela.CreateBy = user.GetUserId(c)

	err = s.InsertRelation(&reqRela)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "创建成功")
}

//----------------------------------------put report----------------------------------------

// CancelReport
// @Summary 撤销申请报告
// @Description 撤销申请报告
// @Tags 用户-报告
// @Param ReportId query string false "报告ID"
// @Router /api/v1/user-agent/report/cancel/{report_id} [put]
// @Security Bearer
func (e Report) CancelReport(c *gin.Context) {
	s := serviceUser.Report{}
	req := dtos.ReportGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.ReportId, err = strconv.Atoi(c.Param("report_id"))
	req.SetUpdateBy(user.GetUserId(c))
	req.RejectTag = dtos.CancelTag

	err = s.UpdateReport(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "撤销成功")
}

// ReapplyReport
// @Summary 重新申请报告
// @Description 重新申请报告
// @Tags 用户-报告
// @Param ReportId query string false "报告ID"
// @Router /api/v1/user-agent/report/reApp/{report_id} [put]
// @Security Bearer
func (e Report) ReapplyReport(c *gin.Context) {
	s := serviceUser.Report{}
	req := dtos.ReportGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.ReportId, err = strconv.Atoi(c.Param("report_id"))
	req.SetUpdateBy(user.GetUserId(c))
	req.RejectTag = dtos.ApplyTag

	err = s.UpdateReport(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "撤销成功")
}

//----------------------------------------------------------------------------------

// GetPatentByReId
// @Summary 通过报告Id获取对应Patent
// @Description  通过报告Id获取对应专利(报告Id 多对一 专利Id)
// @Tags 用户-报告
// @Param ReportId query string false "报告ID"
// @Router /apis/v1/user-agent/report/patent/{report_id} [get]
// @Security Bearer
func (e Report) GetPatentByReId(c *gin.Context) {

	s := serviceAdmin.Report{}
	sUser := serviceUser.Patent{}
	req := dtos.ReportRelaReq{}
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

	req.ReportId, err = strconv.Atoi(c.Param("report_id")) //接受成功
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var object model.ReportRelation

	err = s.GetPatentByReId(req.ReportId, &object)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	model := models.Patent{}
	req.PatentId = object.PatentId

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&sUser.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = sUser.GeById(&req, &model)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(model, "查询成功")
}

// GetReportListByPatentId
// @Summary 通过PatentId 获取对应报告列表
// @Description  通过PatentId 获取对应报告列表
// @Tags 用户-报告
// @Param ReportId query string false "专利ID"
// @Router /apis/v1/user-agent/report/reportList/{patent_id} [get]
// @Security Bearer
func (e Report) GetReportListByPatentId(c *gin.Context) {

	s := serviceAdmin.Report{}
	req := dtos.ReportRelaReq{}
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

	req.PatentId, err = strconv.Atoi(c.Param("patent_id")) //接受成功
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]model.ReportRelation, 0)
	err = s.GetReportIdsByPatentId(req.PatentId, user.GetUserId(c), &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req1 := dtos.ReportIds{}
	req1.ReportIds = make([]int, len(list)) //必须在这里初始化！

	for i := 0; i < len(list); i++ {
		req1.ReportIds[i] = list[i].ReportId
	}

	list1 := make([]model.Report, 0)

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.GetReportListByIds(&req1, &list1)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(list1, "查询成功")
}

// UserGetReportById
// @Summary 查询报告
// @Description  通过ReportId检索报告
// @Tags 用户-报告
// @Param ReportId query string false "报告ID"
// @Router /apis/v1/user-agent/report/query/{report_id} [get]
// @Security Bearer
func (e Report) UserGetReportById(c *gin.Context) {

	s := serviceAdmin.Report{}
	sUser := serviceUser.Report{}
	req := dtos.ReportRelaReq{}
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
	req.ReportId, err = strconv.Atoi(c.Param("report_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object model.ReportRelation
	err = s.GetPatentByReId(req.ReportId, &object)
	//判断是否是认领的专利的报告 ：如果此处查询咩问题，说明用户认领过patent
	if err != nil {
		e.Logger.Error(err)
		e.Error(401, err, "您没有认领该专利，无法查看报告")
		return
	}
	model := model.Report{}
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&sUser.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = sUser.UserGetRepoById(req.ReportId, &model)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(model, "查询成功")
}
