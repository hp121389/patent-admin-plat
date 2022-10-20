package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
)

type UserPatent struct {
	api.Api
}

// GetClaims
// @Summary 获得该UserId的认领专利列表
// @Description 获取认领关系
// @Tags 用户专利关系表
// @Router /api/v1/user-patent/claim [get]
// @Security Bearer
func (e UserPatent) GetClaims(c *gin.Context) { //gin框架里的上下文

	s := service.UserPatent{}         //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req.UserId = user.GetUserId(c)
	req1 := dto.PatentsByIdsForRelationshipUsers{}

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
	err = s.GetPatentPagesByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")
}

// GetCollections
// @Summary 获得该UserId的关注专利列表
// @Description 获取关注关系
// @Tags 用户专利关系表
// @Router /api/v1/user-patent/collection [get]
// @Security Bearer
func (e UserPatent) GetCollections(c *gin.Context) { //gin框架里的上下文

	s := service.UserPatent{}         //service中查询或者返回的结果赋值给s变量
	req := dto.UserPatentGetPageReq{} //被绑定的数据
	req1 := dto.PatentsByIdsForRelationshipUsers{}

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

	err = s.GetCollectionLists(&req, &list, &count)

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

	err = s.GetPatentPagesByIds(&req1, &list1, &count2)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list1, "查询成功")
}

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

// DeleteUserPatentRelationship
// @Summary 根据专利id、TYPE删除用户专利关系
// @Description  根据专利id、TYPE删除用户专利关系
// @Tags 用户专利关系表
// @Param PatentId query string false "专利ID"
// @Param Type query string false "关系类型"
// @Router /api/v1/user-patent/{patent_id}/{type} [delete]
// @Security Bearer
func (e UserPatent) DeleteUserPatentRelationship(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.UserPatentObject{}
	req.UserId = user.GetUserId(c)

	req.SetUpdateBy(user.GetUserId(c))

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req). //在这一步传入request数据
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
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

// UpdateUserPatentRelationship
// @Summary 修改用户专利关系
// @Description 需要输入专利id更新用户专利关系
// @Tags 用户专利关系表
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpDateUserPatentObject true "body"
// @Router /api/v1/user-patent [put]
// @Security Bearer
func (e UserPatent) UpdateUserPatentRelationship(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.UpDateUserPatentObject{}
	req.UserId = user.GetUserId(c)
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

	req.SetUpdateBy(user.GetUserId(c))
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	if req.PatentId == 0 {
		e.Logger.Error(err)
		e.Error(404, err, "请输入专利id")
		return
	}

	err = s.UpdateUserPatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "更新成功")
}
