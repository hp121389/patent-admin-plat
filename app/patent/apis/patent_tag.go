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

type PatentTag struct {
	api.Api
}

// GetTags
// @Summary 获得该PatentId的Tag列表
// @Description 获取标签
// @Tags 专利标签关系表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/patent-tag/tags/{patent_id} [get]
// @Security Bearer
func (e PatentTag) GetTags(c *gin.Context) {

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

// GetPatent
// @Summary 获得该TagId的Patents列表
// @Description 获取标签
// @Tags 专利标签关系表
// @Param TagId query string false "标签ID"
// @Router /api/v1/patent-tag/patents/{tag_id} [get]
// @Security Bearer
func (e PatentTag) GetPatent(c *gin.Context) {

	s := service.PatentTag{}
	req := dto.TagPageGetReq{}
	req1 := dto.PatentsByIdsForRelationshipTags{}

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
		req1.PatentIds[i] = list[i].TagId
	}

	err = s.GetPatentPages(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")

}

// InsertPatentTagRelationship
// @Summary 创建专利标签关系
// @Description  创建专利标签关系
// @Tags 专利标签关系表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentTagInsertReq true "TagId和PatentId为必要输入"
// @Router /api/v1/patent-tag/ [post]
// @Security Bearer
func (e PatentTag) InsertPatentTagRelationship(c *gin.Context) {
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
