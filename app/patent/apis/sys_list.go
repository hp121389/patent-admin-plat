package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
	"net/http"

	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
)

type SysList struct {
	api.Api
}

// GetPatentById
// @Summary 通过专利id获取单个对象
// @Description 获取JSON
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Param TI  query string false "专利名"
// @Param PNM  query string false "申请号"
// @Param AD query string false "申请日"
// @Param PD query string false "公开日"
// @Param CL query string false "简介"
// @Param PA query string false "申请单位"
// @Param AR  query string false "地址"
// @Param INN  query string false "申请人"
// @Router /api/v1/sys-list/patentid/{patent_id} [get]
// @Security Bearer
func (e SysList) GetPatentById(c *gin.Context) {
	s := service.SysList{}
	req := dto.SysListById{}
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
	var object models.SysList
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

//// GetPatentByName
//// @Summary 通过专利名称获取单个对象
//// @Description 获取JSON
//// @Tags 专利表
//// @Param PatentId query string false "专利ID"
//// @Param TI  query string false "专利名"
//// @Param PNM  query string false "申请号"
//// @Param AD query string false "申请日"
//// @Param PD query string false "公开日"
//// @Param CL query string false "简介"
//// @Param PA query string false "申请单位"
//// @Param AR  query string false "地址"
//// @Param INN  query string false "申请人"
//// @Router /api/v1/sys-list/patentname [get]
//// @Security Bearer
//func (e SysList) GetPatentByName(c *gin.Context) {
//	s := service.SysList{}
//	req := dto.SysListByName{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, nil).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	var object models.SysList
//	//数据权限检查
//	//p := actions.GetPermissionFromContext(c)
//	err = s.GetByName(&req, &object)
//	if err != nil {
//		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
//		return
//	}
//	e.OK(object, "查询成功")
//}

// GetLists
// @Summary 列表专利信息数据
// @Description 获取JSON
// @Tags 专利表
// @Router /api/v1/sys-list [get]
// @Security Bearer
func (e SysList) GetLists(c *gin.Context) {
	s := service.SysList{}
	req := dto.SysListGetPageReq{}
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

	list := make([]models.SysList, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// InsertListsByPatentId
// @Summary 创建专利
// @Description 获取JSON
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysListInsertReq true "专利表数据"
// @Router /api/v1/sys-list [post]
// @Security Bearer
func (e SysList) InsertListsByPatentId(c *gin.Context) {
	s := service.SysList{}
	req := dto.SysListInsertReq{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.InsertListsByPatentId(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetPatentId(), "创建成功")
}

// UpdateLists
// @Summary 修改专利表数据
// @Description 修改JSON
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysListUpdateReq true "body"
// @Router /api/v1/sys-list/ [put]
// @Security Bearer
func (e SysList) UpdateLists(c *gin.Context) {
	s := service.SysList{}
	req := dto.SysListUpdateReq{}
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
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.UpdateLists(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetPatentId(), "更新成功")
}

// DeleteLists
// @Summary 输入id删除专利表
// @Description 专利删除
// @Tags 专利表
// @Param data body dto.SysListDeleteReq true "body"
// @Router /api/v1/sys-list [delete]
// @Security Bearer
func (e SysList) DeleteLists(c *gin.Context) {
	s := service.SysList{}
	req := dto.SysListDeleteReq{}
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

	// 设置编辑人
	//req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetPatentId(), "删除成功")
}
