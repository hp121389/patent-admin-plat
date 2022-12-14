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
//// @Param data body dto.PatentReq true "专利表数据"
//// @Router /api/v1/user-agent/patent [post]
//// @Security Bearer
//func (e Patent) InsertPatent(c *gin.Context) {
//	s := service.Patent{}
//	req := dto.PatentReq{}
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

//
//// ---------------------------------------------------patent--graph-------------------------------------------------------
////
//// GetTheGraphByUserId
//// @Summary 获取专利关系图
//// @Description  获取专利关系图
//// @Tags 专利表
//// @Router /api/v1/user-agent/patent/relationship [get]
//// @Security Bearer
//func (e Patent) GetTheGraphByUserId(c *gin.Context) {
//	//spp := service.PatentPackage{}
//	sup := service.UserPatent{}
//	su := aservice.SysUser{}
//	sp := service.Patent{}
//	gservice := service.Node{}
//	//reqpp := dto.PackagePageGetReq{} //patent-package
//	reqp := dto.PatentsIds{} //patents
//	requ := adto.SysUserById{}
//	//fmt.Println("get the line 471")
//	//fmt.Println(c)
//	var err error
//	//reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
//	//fmt.Println("get the line 474")
//	//fmt.Println(reqpp.PackageId)
//	//s := service.UserPatent{}
//
//	requp := dto.UserPatentGetPageReq{}
//	//reqp := dto.PatentsIds{}
//
//	requp.UserId = user.GetUserId(c)
//
//	err = e.MakeContext(c).
//		MakeOrm().
//		MakeService(&sup.Service).
//		Errors
//
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	//数据权限检查
//	//p := actions.GetPermissionFromContext(c)
//	listup1 := make([]models.UserPatent, 0)
//	listp := make([]models.Patent, 0)
//
//	//var count int64                                           //not used
//	err = sup.GetUserFocusPatentIds(&requp, &listup1) //
//
//	if err != nil {
//		e.Error(500, err, "查询失败")
//		return
//	}
//
//	var count2 int64 //not used
//	err = e.MakeContext(c).
//		MakeOrm().
//		//Bind(&reqp).
//		MakeService(&sp.Service).
//		Errors
//
//	reqp.PatentIds = make([]int, len(listup1))
//	for i := 0; i < len(listup1); i++ {
//		reqp.PatentIds[i] = listup1[i].PatentId
//	}
//
//	err = sp.GetPageByIds(&reqp, &listp, &count2)
//	fmt.Println("找到了所有的关注的patent")
//	fmt.Println(listp)
//
//	//err = e.MakeContext(c).
//	//	MakeOrm().
//	//	MakeService(&sup.Service).
//	//	Errors
//	////fmt.Println("get the line 480")
//	//if err != nil {
//	//	e.Logger.Error(err)
//	//	e.Error(500, err, err.Error())
//	//	return
//	//}
//	//reqpp.SetUpdateBy(user.GetUserId(c))
//	//fmt.Println("get the line 486")
//	//listpp := make([]models.PatentPackage, 0)
//	//var count int64 //  not used
//	//err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
//	//fmt.Println(listpp)
//	//fmt.Println(reqp)
//	//for i := 0; i < len(listpp); i++ {
//	//	fmt.Println(listpp[i].PatentId)
//	//}
//	//reqp.PatentIds = make([]int, len(listpp))
//	//for i := 0; i < len(listpp); i++ {
//	//	reqp.PatentIds[i] = listpp[i].PatentId
//	//}
//	//fmt.Println("get line 496")
//	//1 := make([]models.Node, 0) //resultnode
//	links := make([]models.Link, 0) //resultlink
//	//listup2, members, usertimes := e.AddGraphNodeByReq(c, &nodes, reqp, 5)
//	//------------------------------------already get the patents id  now get the users id
//	err = e.MakeContext(c).
//		MakeOrm().
//		MakeService(&sup.Service).
//		Errors
//	//fmt.Println("get line 505")
//	listup2 := make([]models.UserPatent, 0) //√
//	//fmt.Println(reqp.PatentIds)
//	sup.GetFocusUsersByPatentId(&listup2, &reqp) //√
//	//for i := 0; i < len(listup2); i++ {
//	//	fmt.Println(listup2[i])
//	//}
//	//fmt.Println("the listup2 u success")
//	//-----------------------------------already get the users id  now sort the users  and pick 8 users
//	usertimes := make(map[int]int) // k-v is uid-times
//
//	for i := 0; i < len(listup2); i++ { //建立map usertimes
//		if usertimes[listup2[i].UserId] == 0 {
//			usertimes[listup2[i].UserId] = 1
//		} else {
//			usertimes[listup2[i].UserId]++
//		}
//
//	}
//	//fmt.Println("show the usertimes:")
//	//fmt.Println(usertimes)
//	usertimes1 := rankByWordCount(usertimes) //usertimes1有序的 key-value  uid-uidtimes 以times排序
//	fmt.Println("show the usertimes1:")
//	fmt.Println(usertimes1)
//	//get 5th user
//	var members int
//	if len(usertimes1) < 75 {
//		members = len(usertimes1)
//	} else {
//		members = 75
//	}
//	listu := make([]amodels.SysUser, members)
//	err = e.MakeContext(c).
//		MakeOrm().
//		Bind(&requ, nil).
//		MakeService(&su.Service).
//		Errors
//	for i := 0; i < members; i++ {
//		requ.Id = usertimes1[i].Key
//		su.Get(&requ, &listu[i])
//		fmt.Println("this is %d", i)
//		fmt.Println(listu[i])
//	}
//
//	NodeList := make([]models.Node, 0)
//	err = e.MakeContext(c).
//		MakeOrm().
//		MakeService(&gservice.Service).
//		Errors
//	//err = gservice.GetNodes(&NodeList)
//	for i := 0; i < members; i++ {
//		fmt.Println(i)
//		err = gservice.GetNodes(&NodeList, i) //√
//		NodeList[i].NodeValue = usertimes1[i].Value
//		NodeList[i].NodeName = listu[i].Username
//		//fmt.Println(NodeList[i])
//		//nownode.GraphName =    应该加上
//	}
//
//	//有75个人的uid，package里的patent的pid
//	for i := 0; i < members; i++ {
//		for j := i + 1; j < members; j++ {
//			//pantentssum := len(listpp)
//			//var ispatent [2][pantentssum]bool
//			//RelationExist, _ := sup.GetTwoUserRelationshipInThisPackage(&reqp, usertimes1[i].Key, usertimes1[j].Key)
//			RelationExist := 0
//			listpatent1 := make([]int, 0)
//			listpatent2 := make([]int, 0)
//			for z := 0; z < len(listup2); z++ {
//				if listup2[z].UserId == usertimes1[i].Key {
//					listpatent1 = append(listpatent1, listup2[z].PatentId)
//				}
//			}
//			for z := 0; z < len(listup2); z++ {
//				if listup2[z].UserId == usertimes1[j].Key {
//					listpatent2 = append(listpatent2, listup2[z].PatentId)
//				}
//			}
//			for z := 0; z < len(listpatent1); z++ {
//				for z1 := 0; z1 < len(listpatent2); z1++ {
//					if listpatent1[z] == listpatent2[z1] {
//						RelationExist++
//						break
//					}
//				}
//			}
//			if RelationExist != 0 {
//				var nowlink models.Link
//				nowlink.Source = strconv.FormatInt((int64(usertimes1[i].Key)), 10)
//				nowlink.Target = strconv.FormatInt((int64(usertimes1[j].Key)), 10)
//				nowlink.Value = RelationExist
//				links = append(links, nowlink)
//				//fmt.Println(usertimes1[i].Key, usertimes1[j].Key)
//				//fmt.Println(i, j, RelationExist, nowlink)
//			}
//
//		}
//	} //√
//	//fmt.Println(links)
//	result := dto.GraphResult{}
//	result.GetNodesAndLinks(&NodeList, &links)
//	//fmt.Println(NodeList)
//	//fmt.Println(links)
//	//fmt.Println(result)
//	e.GraphOK(NodeList, links, "查询成功")
//
//	////get each 8 users of members users
//	//for i := 0; i < members; i++ { //each node be added in nodes[]
//	//	reqp2 := dto.PatentsIds{} //find patent listup2 of node[i]
//	//	z := 0
//	//	for j := 0; j < len(listup2); j++ {
//	//		nodeid := usertimes[i].Key //get node id
//	//		if listup2[j].UserId == nodeid {
//	//			reqp2.PatentIds[z] = listup2[j].PatentId
//	//			z++
//	//		}
//	//	}
//	//	_, _, _ = e.AddGraphNodeByReq(c, &nodes, reqp2, i*14+5)
//	//}
//
//}

//// AddGraphNodeByReq -----------------
//func (e Package) AddGraphNodeByReq(c *gin.Context, nodes *[]models.Graph, req dto.PatentsIds, times int) ([]models.UserPatent, int, []Pair) {
//	sup := service.UserPatent{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(req).
//		MakeService(&sup.Service).
//		Errors
//	listup := make([]models.UserPatent, 0)
//	err = sup.GetUsersByPatentId(&listup, req)
//	if err != nil {
//		e.Logger.Error(err)
//		return nil, 0, nil
//	}
//	fmt.Println(listup[0].UserId)
//	fmt.Println("the list u success")
//	//-----------------------------------already get the users id  now sort the users  and pick 8 users
//	usertimes := make(map[int]int) // k-v is uid-times
//	for i := 0; i < len(listup); i++ {
//		if usertimes[listup[i].UserId] == 0 {
//			usertimes[listup[i].UserId]++
//		} else {
//			usertimes[listup[i].UserId] = 1
//		}
//
//	}
//
//	usertimes1 := rankByWordCount(usertimes) //有序的 key-value  uid-uidtimes 以times排序
//	ids := []int{48, 11, 27, 24, 55,}
//	//get 5th user
//	//nodes := make([]models.Graph, 0) //resultnode
//	//links := make([]models.Link, 0)  //resultlink
//	begin := 0
//	if times > 5 { //后续节点
//		begin = times
//		times = 8
//	} //time默认是5
//	if len(usertimes1) < times {
//		times = len(usertimes1)
//	}
//	for i := begin; i < begin+times; i++ {
//		var nownode models.Graph
//		nownode.GraphValue = string(usertimes1[i-begin].Value) //第1，2，3，4·····个元素
//		nownode.GraphId = strconv.Itoa(ids[i])
//		//nownode.GraphName =    应该加上
//		*nodes = append(*nodes, nownode)
//	}
//	return listup, begin, usertimes1
//}
