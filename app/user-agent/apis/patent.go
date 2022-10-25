package apis

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"net/http"
	"strconv"
)

type Patent struct {
	api.Api
}

// GetPatentById
// @Summary 检索专利
// @Description  通过PatentId检索专利
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/{patent_id} [get]
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
// @Description 获取本地专利
// @Tags 专利表
// @Router /api/v1/user-agent/patent [get]
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
// @Summary 添加专利
// @Description 添加专利到本地
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "专利表数据"
// @Router /api/v1/user-agent/patent [post]
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
// @Summary 修改专利
// @Description 必须要有主键PatentId值
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentUpdateReq true "body"
// @Router /api/v1/user-agent/patent [put]
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
	e.OK(req, "更新成功")
}

// DeletePatent
// @Summary 删除专利
// @Description  输入专利id删除专利表
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeletePatent(c *gin.Context) {
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

	req.PatentId, err = strconv.Atoi(c.Param("patent_id"))
	req.UpdateBy = user.GetUserId(c)

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
// @Description 认领专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/claim [post]
// @Security Bearer
func (e Patent) ClaimPatent(c *gin.Context) {

	pid, PNM, err := e.internalInsertIfAbsent(c)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.Patent{}
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

	req := dto.NewUserPatentClaim(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), PNM)

	if err = s.InsertUserPatent(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "认领成功")
}

// FocusPatent
// @Summary 关注专利
// @Description 关注专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentInsertReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/focus [post]
// @Security Bearer
func (e Patent) FocusPatent(c *gin.Context) {

	pid, PNM, err := e.internalInsertIfAbsent(c)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.Patent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), PNM)

	if err = s.InsertUserPatent(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "关注成功")
}

func (e Patent) internalInsertIfAbsent(c *gin.Context) (int, string, error) {
	ps := service.Patent{}
	req := dto.PatentInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		return 0, "", err
	}
	p, err := ps.InsertIfAbsent(&req)
	if err != nil {
		return 0, "", err
	}
	return p.PatentId, p.PNM, nil
}

// GetFocusPages
// @Summary 获取关注列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/focus [get]
// @Security Bearer
func (e Patent) GetFocusPages(c *gin.Context) {
	s := service.Patent{}             //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req.UserId = user.GetUserId(c)
	req1 := dto.PatentsIds{}

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
	list := make([]models.UserPatent, 0)
	list1 := make([]models.Patent, 0)
	var count int64
	err = s.GetFocusLists(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	var count2 int64
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors
	req1.PatentIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}
	err = s.GetPageByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")
}

// GetClaimPages
// @Summary 获取认领列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/claim [get]
// @Security Bearer
func (e Patent) GetClaimPages(c *gin.Context) {
	s := service.Patent{}             //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req1 := dto.PatentsIds{}

	req.UserId = user.GetUserId(c)

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
	list := make([]models.UserPatent, 0)
	list1 := make([]models.Patent, 0)

	var count int64

	err = s.GetClaimLists(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors

	req1.PatentIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s.GetPageByIds(&req1, &list1, &count2)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list1, "查询成功")
}

// DeleteFocus
// @Summary 取消关注
// @Description  取消关注
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/focus/{patent_id}  [delete]
// @Security Bearer
func (e Patent) DeleteFocus(c *gin.Context) {
	s := service.Patent{}
	pid, err := strconv.Atoi(c.Param("patent_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), "")

	err = e.MakeContext(c).
		MakeOrm().
		Bind(req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveFocus(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "取消关注成功")
}

// DeleteClaim
// @Summary 取消认领
// @Description  取消认领
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/claim/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeleteClaim(c *gin.Context) {

	s := service.Patent{}

	pid, err := strconv.Atoi(c.Param("patent_id"))
	if err != nil {
		e.Logger.Error(err)
		return
	}

	req := dto.NewUserPatentClaim(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), "")

	err = e.MakeContext(c).
		MakeOrm().
		Bind(req). //修改&
		MakeService(&s.Service).
		Errors

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveClaim(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "取消认领成功")
}

// DeleteTag
// @Summary 取消给该专利添加的该标签
// @Description  取消给该专利添加的该标签
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Param TagId query string false "标签ID"
// @Router /api/v1/user-agent/patent/tags/{tag_id}/patent/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeleteTag(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentTagInsertReq{}
	req.SetUpdateBy(user.GetUserId(c))
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

	req.TagId, err = strconv.Atoi(c.Param("tag_id"))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveRelationship(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "删除成功")
}

// InsertTag
// @Summary 为该专利添加该标签
// @Description  为该专利添加该标签
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentTagInsertReq true "TagId和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/tag [post]
// @Security Bearer
func (e Patent) InsertTag(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentTagInsertReq{}

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

	if req.PatentId == 0 || req.TagId == 0 {
		e.Logger.Error(err)
		e.Error(404, err, "您输入的专利id不存在！")
		return
	}

	err = s.InsertPatentTagRelationship(&req)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "创建成功")
}

// GetPatent
// @Summary 显示该标签下的专利
// @Description 显示该标签下的专利
// @Tags 专利表
// @Param TagId query string false "标签ID"
// @Router /api/v1/user-agent/patent/tag-patents/{tag_id} [get]
// @Security Bearer
func (e Patent) GetPatent(c *gin.Context) {

	s := service.Patent{}
	req := dto.TagPageGetReq{}
	req1 := dto.PatentsIds{}

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

	req.TagId, err = strconv.Atoi(c.Param("tag_id"))

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]models.PatentTag, 0)
	list1 := make([]models.Patent, 0)
	var count int64

	err = s.GetPatentIdByTagId(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors

	req1.PatentIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s.GetPageByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")

}

// GetTags
// @Summary 显示专利的标签
// @Description 显示专利的标签
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/tags/{patent_id} [get]
// @Security Bearer
func (e Patent) GetTags(c *gin.Context) {

	s := service.Patent{}
	req := dto.PatentTagGetPageReq{}
	req1 := dto.TagsByIdsForRelationshipPatents{}

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

	req.PatentId, err = strconv.Atoi(c.Param("patent_id"))

	list := make([]models.PatentTag, 0)
	list1 := make([]models.Tag, 0)
	var count int64

	err = s.GetTagIdByPatentId(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors

	req1.TagIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.TagIds[i] = list[i].TagId
	}

	err = s.GetTagPages(&req1, &list1, &count2)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list1, "查询成功")

}

// GetPackagePatents
// @Summary 显示专利包内专利
// @Description 显示专利包内专利
// @Tags 专利表
// @Param TagId query string false "标签ID"
// @Router /api/v1/user-agent/patent/package/{package_id} [get]
// @Security Bearer
func (e Patent) GetPackagePatents(c *gin.Context) {

	s := service.Patent{}
	req := dto.PackagePageGetReq{}
	req1 := dto.PatentsIds{}

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

	req.PackageId, err = strconv.Atoi(c.Param("package_id"))

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]models.PatentPackage, 0)
	list1 := make([]models.Patent, 0)
	var count int64

	err = s.GetPatentIdByPackageId(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors

	req1.PatentIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s.GetPageByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")

}

// InsertPackagePatent
// @Summary 将专利加入专利包
// @Description  将专利加入专利包
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackagePageGetReq true "PackageId和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/package [post]
// @Security Bearer
func (e Patent) InsertPackagePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PackagePageGetReq{}

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

	if req.PatentId == 0 || req.PackageId == 0 {
		e.Logger.Error(err)
		e.Error(404, err, "您输入的专利id不存在！")
		return
	}
	err = s.InsertPatentPackage(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.PackageBack, "创建成功")
}

// DeletePackagePatent
// @Summary 取消加入该专利包
// @Description  取消加入该专利包
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Param PackageId query string false "专利包ID"
// @Router /api/v1/user-agent/patent/package/{package_id}/patent/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeletePackagePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PackagePageGetReq{}
	req.SetUpdateBy(user.GetUserId(c))
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

	req.PatentId, err = strconv.Atoi(c.Param("patent_id"))
	req.PackageId, err = strconv.Atoi(c.Param("package_id"))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemovePackagePatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.PackageBack, "删除成功")
}

// GetUserPatentsPages
// @Summary 获取用户的专利列表
// @Description 获取用户的专利列表
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/user [get]
// @Security Bearer
func (e Patent) GetUserPatentsPages(c *gin.Context) {

	s := service.Patent{}             //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req1 := dto.PatentsIds{}

	req.UserId = user.GetUserId(c)

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
	list := make([]models.UserPatent, 0)
	list1 := make([]models.Patent, 0)

	var count int64
	err = s.GetUserPatentIds(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s.Service).
		Errors

	req1.PatentIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s.GetPageByIds(&req1, &list1, &count2)

	fmt.Println(list1)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")
}

//// UpdateUserPatentRelationship
//// @Summary 修改用户专利关系
//// @Description 需要输入专利id
//// @Tags 专利表
//// @Accept  application/json
//// @Product application/json
//// @Param data body dto.UpDateUserPatentObject true "body"
//// @Router /api/v1/user-agent/patent [put]
//// @Security Bearer
//func (e Patent) UpdateUserPatentRelationship(c *gin.Context) {
//	s := service.UserPatent{}
//	req := dto.UpDateUserPatentObject{}
//	req.UserId = user.GetUserId(c)
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, binding.JSON).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	req.SetUpdateBy(user.GetUserId(c))
//	//数据权限检查
//	//p := actions.GetPermissionFromContext(c)
//
//	if req.PatentId == 0 {
//		e.Logger.Error(err)
//		e.Error(404, err, "请输入专利id")
//		return
//	}
//
//	err = s.UpdateUserPatent(&req)
//
//	if err != nil {
//		e.Logger.Error(err)
//		return
//	}
//	e.OK(req, "更新成功")
//}
//// InsertUserPatentRelationship
//// @Summary 创建用户专利关系
//// @Description Type和PatentId为必要输入，Type只能是 认领 或者 关注 关系
//// @Tags 用户专利关系表
//// @Accept  application/json
//// @Product application/json
//// @Param data body dto.UserPatentInsertReq true "Type和PatentId为必要输入"
//// @Router /api/v1/user-patent/ [post]
//// @Security Bearer
//func (e UserPatent) InsertUserPatentRelationship(c *gin.Context) {
//	s := service.UserPatent{}
//	req := dto.UserPatentInsertReq{}
//	req.UserId = user.GetUserId(c)
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req, binding.JSON).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	// 设置创建人
//	req.SetCreateBy(user.GetUserId(c))
//
//	if req.Type == "认领" {
//		err = s.Insert(&req)
//	} else if req.Type == "关注" {
//		err = s.InsertCollectionRelationship(&req)
//	} else {
//		e.Logger.Error(err)
//		e.Error(404, err, fmt.Sprintf("invalid req.Type: %s, should be 认领/关注", req.Type))
//		return
//	}
//
//	if req.PatentId == 0 {
//		e.Logger.Error(err)
//		e.Error(404, err, "您输入的专利id不存在！")
//		return
//	}
//
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	e.OK(req, "创建成功")
//}
