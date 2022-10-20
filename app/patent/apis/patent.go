package apis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
	"net/http"
)

type Patent struct {
	api.Api
}

// GetPatentById
// @Summary 通过专利id获取单个对象
// @Description 获取JSON,希望可以通过以下参数高级搜索，暂时只支持patentId
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/patent/{patent_id} [get]
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
// @Router /api/v1/patent [get]
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
// @Router /api/v1/patent [post]
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
	err = s.Insert(&req)
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
// @Router /api/v1/patent [put]
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
// @Router /api/v1/patent/{patent_id} [delete]
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
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetPatentId(), "删除成功")
}

// ClaimPatent
// @Summary 认领专利
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.UserPatentInsertReq true "Type和PatentId为必要输入"
// @Router /api/v1/patent/{patent_id}/claim [post]
// @Security Bearer
func (e Patent) ClaimPatent(c *gin.Context) {
	pid, err := e.internalInsertIfAbsent(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.UserPatent{}
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentClaim(user.GetUserId(c), pid, user.GetUserId(c))

	if err = s.Insert(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "创建成功")
}

// FocusPatent
// @Summary 关注专利
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/patent/{patent_id}/focus [post]
// @Security Bearer
func (e Patent) FocusPatent(c *gin.Context) {
	pid, err := e.internalInsertIfAbsent(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.UserPatent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), pid, user.GetUserId(c))

	if err = s.Insert(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "创建成功")
}

func (e Patent) internalInsertIfAbsent(c *gin.Context) (int, error) {
	ps := service.Patent{}
	req := dto.PatentInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		return 0, err
	}
	return ps.InsertIfAbsent(&req)
}
