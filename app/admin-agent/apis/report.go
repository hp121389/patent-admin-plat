package apis

import (
	"fmt"
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

//
//----------------------------------------GET report----------------------------------------

// GetPatentByReId
// @Summary 通过报告Id获取对应Patent
// @Description  通过报告Id获取对应专利(报告Id 多对一 专利Id)
// @Tags 报告
// @Param ReportId query string false "报告ID"
// @Router /apis/v1/admin-agent/report/patent/{report_id} [get]
// @Security Bearer
func (e Report) GetPatentByReId(c *gin.Context) {

	s := service.Report{}
	sUser := service2.Patent{}
	req := dtos.PatentReport{}

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

	var object model.PatentReport

	fmt.Println("进入service")

	err = s.GetPatentByReId(req.ReportId, &object)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	model := models.Patent{}

	req.PatentId = object.PatentId

	fmt.Println(req.PatentId)

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&sUser.Service).
		Errors

	err = sUser.GeById(&req, &model)

	e.OK(model, "查询成功")
}

// GetInfringementLists
// @Summary 列表侵权报告
// @Description 列表侵权报告
// @Tags 侵权报告
// @Router /apis/v1/admin-agent/infringement-report/ [get]
// @Security Bearer
func (e Report) GetInfringementLists(c *gin.Context) {
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

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]model.Report, 0)

	err = s.GetInfringementReportPages(&list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// GetReportById
// @Summary 查看报告
// @Description  通过ReportId检索报告
// @Tags 报告
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

// GetValuationLists
// @Summary 列表估值报告
// @Description
// @Tags 估值报告
// @Router /apis/v1/admin-agent/valuation-report [get]
// @Security Bearer
func (e Report) GetValuationLists(c *gin.Context) { //gin框架里的上下文
	s := service.Report{}          //service中查询或者返回的结果赋值给s变量
	req := dtos.ReportGetPageReq{} //被绑定的数据
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

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]model.Report, 0)

	err = s.GetValuationReportPages(&list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

//----------------------------------------upload report----------------------------------------

// UploadReport
// @Summary 上传报告（本质是更新）
// @Description 请求体必须包含reportId，仅传文件名和路径（json）
// @Tags 报告
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

//----------------------------------------驳回请求----------------------------------------

// Reject
// @Summary 驳回请求
// @Description 修改tag+清空文件
// @Tags 报告
// @Router /apis/v1/admin-agent/report/reject/{report_id} [post]
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
	fmt.Println(user.GetUserId(c))
	fmt.Println(req.UpdateBy)
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
// @Tags 报告
// @Router /apis/v1/admin-agent/report/unReject/{report_id} [post]
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
