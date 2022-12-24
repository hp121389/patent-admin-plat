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

//----------------------------------------patent----------------------------------------

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
	var object2 models.Patent
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	req.PatentId, err = strconv.Atoi(c.Param("patent_id"))
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "not found params from router")
		return
	}
	err = s.Get(&req, &object2)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object2, "查询成功")
}

// GetPatentLists
// @Summary 列表专利信息数据
// @Description 获取本地专利
// @Tags 专利表
// @Router /api/v1/user-agent/patent [get]
// @Security Bearer
func (e Patent) GetPatentLists(c *gin.Context) { //gin框架里的上下文
	s := service.Patent{}  //service中查询或者返回的结果赋值给s变量
	req := dto.PatentReq{} //被绑定的数据
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

	e.OK(list, "查询成功")
}

//// InsertPatent
//// @Summary 添加专利
//// @Description 添加专利到本地
//// @Tags 专利表
//// @Accept  application/json
//// @Product application/json
//// @Param data body dtos.PatentReq true "专利表数据"
//// @Router /api/v1/user-agent/patent [post]
//// @Security Bearer
//func (e Patent) InsertPatent(c *gin.Context) {
//	s := service.Patent{}
//	req := dtos.PatentReq{}
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
//	err = s.Insert(&req)
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	e.OK(req, "创建成功")
//}

// UpdatePatent
// @Summary 修改专利
// @Description 必须要有主键PatentId值
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "body"
// @Router /api/v1/user-agent/patent [put]
// @Security Bearer
func (e Patent) UpdatePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentReq{}
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
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.UpdateBy = user.GetUserId(c)

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "删除成功")
}

//----------------------------------------user-patent-----------------------------------------------------------------

// GetUserPatentsPages
// @Summary 获取用户的专利列表
// @Description 获取用户的专利列表
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/user [get]
// @Security Bearer
func (e Patent) GetUserPatentsPages(c *gin.Context) {

	s := service.UserPatent{}
	s1 := service.Patent{}
	req := dto.UserPatentObject{}
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

	err = s1.GetPageByIds(&req1, &list1, &count2)

	fmt.Println(list1)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")
}

// ClaimPatent
// @Summary 认领专利
// @Description 认领专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/claim [post]
// @Security Bearer
func (e Patent) ClaimPatent(c *gin.Context) {

	pid, PNM, err := e.internalInsertIfAbsent(c)

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
// @Param data body dto.PatentReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/focus [post]
// @Security Bearer
func (e Patent) FocusPatent(c *gin.Context) {

	pid, PNM, err := e.internalInsertIfAbsent(c)

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

// InsertIfAbsent
// @Summary 添加专利
// @Description 添加专利到本地
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "专利表数据"
// @Router /api/v1/user-agent/patent [post]
// @Security Bearer
func (e Patent) InsertIfAbsent(c *gin.Context) {
	pid, pnm, err := e.internalInsertIfAbsent(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = e.MakeContext(c).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(&dto.PatentBriefInfo{PatentId: pid, PNM: pnm}, "success")
}

func (e Patent) internalInsertIfAbsent(c *gin.Context) (int, string, error) {
	ps := service.Patent{}
	req := dto.PatentReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		return 0, "", err
	}
	req.CreateBy = user.GetUserId(c)
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
	s := service.UserPatent{}
	s1 := service.Patent{}
	req := dto.UserPatentObject{}
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
		MakeService(&s1.Service).
		Errors
	req1.PatentIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}
	err = s1.GetPageByIds(&req1, &list1, &count2)
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
	s := service.UserPatent{}
	s1 := service.Patent{}
	req := dto.UserPatentObject{} //被绑定的数据
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
		MakeService(&s1.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req1.PatentIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s1.GetPageByIds(&req1, &list1, &count2)

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
	s := service.UserPatent{}
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

	s := service.UserPatent{}

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

//----------------------------------------user-patent 修改用户专利关系----------------------------------------

//// UpdateUserPatentRelationship
//// @Summary 修改用户专利关系
//// @Description 需要输入专利id
//// @Tags 专利表
//// @Accept  application/json
//// @Product application/json
//// @Param data body dtos.UpDateUserPatentObject true "body"
//// @Router /api/v1/user-agent/patent [put]
//// @Security Bearer
//func (e Patent) UpdateUserPatentRelationship(c *gin.Context) {
//	s := service.UserPatent{}
//	req := dtos.UpDateUserPatentObject{}
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

//----------------------------------------tag-patent----------------------------------------

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
	if err != nil {
		e.Logger.Error(err)
		return
	}
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
	s := service.PatentTag{}
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

	s := service.PatentTag{}
	s1 := service.Patent{}
	req := dto.PatentTagGetPageReq{}
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

	err = s1.GetPageByIds(&req1, &list1, &count2)
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

	s := service.PatentTag{}
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
	if err != nil {
		e.Logger.Error(err)
		return
	}
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
