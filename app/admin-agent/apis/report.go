package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	service "go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/models"
	service2 "go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"net/http"
	"strconv"
)

type Report struct {
	api.Api
}

//----------------------------------------GET report----------------------------------------

// GetReportListByPatentId
// @Summary 通过PatentId 获取对应报告列表
// @Description  通过PatentId 获取对应报告列表
// @Tags 管理员-报告
// @Param ReportId query string false "专利ID"
// @Router /apis/v1/admin-agent/report/reportList/{patent_id} [get]
// @Security Bearer
func (e Report) GetReportListByPatentId(c *gin.Context) {

	s := service.Report{}
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

// GetPatentByReId
// @Summary 通过报告Id获取对应Patent
// @Description  通过报告Id获取对应专利( 报告Id 多对一 专利Id )
// @Tags 管理员-报告
// @Param ReportId query string false "报告ID"
// @Router /apis/v1/admin-agent/report/patent/{report_id} [get]
// @Security Bearer
func (e Report) GetPatentByReId(c *gin.Context) {

	s := service.Report{}
	sUser := service2.Patent{}
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

// GetReportById
// @Summary 查看报告
// @Description  通过ReportId检索报告
// @Tags 管理员-报告
// @Param ReportId query string false "报告ID"
// @Router /apis/v1/admin-agent/report/{report_id} [get]
// @Security Bearer
func (e Report) GetReportById(c *gin.Context) {

	s := service.Report{}
	req := dtos.ReportById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object model.Report
	req.ReportId, err = strconv.Atoi(c.Param("report_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.GetReportById(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// GetListsByType
// @Summary 列表Type类型报告
// @Description 列表Type类型报告
// @Tags 管理员-报告
// @Router /apis/v1/admin-agent/report/type/{type} [get]
// @Security Bearer
func (e Report) GetListsByType(c *gin.Context) {
	s := service.Report{}
	req := dtos.ReportGetPageReq{}
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
	req.Type = c.Param("type")
	list := make([]model.Report, 0)

	err = s.GetPagesByType(req.Type, &list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

//----------------------------------------upload report----------------------------------------

// UploadReport
// @Summary 上传报告
// @Description 请求体必须包含reportId，仅传文件名和路径（json）（本质是更新）
// @Tags 管理员-报告
// @Accept  application/json
// @Product application/json
// @Param data body dtos.ReportGetPageReq true "body"
// @Router /api/v1/admin-agent/report/upload [put]
// @Security Bearer
func (e Report) UploadReport(c *gin.Context) {
	s := service.Report{}
	req := dtos.ReportGetPageReq{}
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

	req.SetUpdateBy(user.GetUserId(c))
	req.FilesOpt = dto.FilesAdd
	req.RejectTag = dtos.UploadTag

	err = s.UploadReport(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}

	e.OK(req, "更新成功")
}

// DeleteUploadReport
// @Summary 删除已经上传报告
// @Description
// @Tags 管理员-报告
// @Accept  application/json
// @Product application/json
// @Param ReportId query string false "报告ID"
// @Router /api/v1/admin-agent/report/files/{report_id} [put]
// @Security Bearer
func (e Report) DeleteUploadReport(c *gin.Context) {
	s := service.Report{}
	req := dtos.ReportGetPageReq{}
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

	req.SetUpdateBy(user.GetUserId(c))
	req.FilesOpt = dto.FilesDelete
	req.RejectTag = dtos.ProcessTag
	req.ReportId, err = strconv.Atoi(c.Param("report_id"))

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.UploadReport(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}

	e.OK(req, "更新成功")
}

//----------------------------------------驳回请求----------------------------------------

// Reject
// @Summary 驳回请求
// @Description 修改tag+清空文件
// @Tags 管理员-报告
// @Router /apis/v1/admin-agent/report/reject/{report_id} [put]
// @Security Bearer
func (e Report) Reject(c *gin.Context) {

	s := service.Report{}
	req := dtos.ReportGetPageReq{}
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

	req.SetUpdateBy(user.GetUserId(c))

	req.ReportId, err = strconv.Atoi(c.Param("report_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.RejectTag = dtos.RejectTag

	err = s.UpdateReports(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}

	e.OK(req, "更新成功")

}

// UnReject
// @Summary 撤销驳回请求
// @Description 必须要有主键ReportId值
// @Tags 管理员-报告
// @Router /apis/v1/admin-agent/report/unReject/{report_id} [put]
// @Security Bearer
func (e Report) UnReject(c *gin.Context) {

	s := service.Report{}
	req := dtos.ReportGetPageReq{}
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

	req.SetUpdateBy(user.GetUserId(c))
	req.ReportId, err = strconv.Atoi(c.Param("report_id"))
	req.RejectTag = dtos.ProcessTag

	err = s.UpdateReports(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}

	e.OK(req, "更新成功")

}
