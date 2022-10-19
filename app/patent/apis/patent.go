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

type Patent struct {
	api.Api
}

// GetPatentById
// @Summary 通过专利id获取单个对象
// @Description 获取JSON,希望可以通过以下参数高级搜索，暂时只支持patentId
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/patent-list/get_by_patent_id/{patent_id} [get]
// @Security Bearer
func (e Patent) GetPatentById(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentById{}
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
	var object models.Patent
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// GetPatentLists
// @Summary 列表专利信息数据
// @Description 获取JSON
// @Tags 专利表
// @Router /api/v1/patent-list/get_patent_lists [get]
// @Security Bearer
func (e Patent) GetPatentLists(c *gin.Context) { //gin框架里的上下文
	s := service.Patent{}         //service中查询或者返回的结果赋值给s变量
	req := dto.PatentGetPageReq{} //被绑定的数据
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

	list := make([]models.Patent, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// InsertPatent
// @Summary 创建专利
// @Description 不是必须要有主键PatentId值（自增），其他需要修改什么输入什么
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "专利表数据"
// @Router /api/v1/patent-list/post_a_patent/ [post]
// @Security Bearer
func (e Patent) InsertPatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentInsertReq{}
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

// UpdatePatent
// @Summary 修改专利表数据
// @Description 在post的json数组必须要有主键PatentId值（默认0不可重复），其他需要修改什么输入什么
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentUpdateReq true "body"
// @Router /api/v1/patent-list/change_a_patent/ [put]
// @Security Bearer
func (e Patent) UpdatePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentUpdateReq{}
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

// DeletePatentByPatentId
// @Summary 输入专利id删除专利表
// @Description  输入专利id删除专利表
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/patent-list/delete_a_patent_by_id/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeletePatentByPatentId(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentById{}
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
