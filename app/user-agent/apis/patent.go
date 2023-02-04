package apis

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	amodels "go-admin/app/admin/models"
	aservice "go-admin/app/admin/service"
	adto "go-admin/app/admin/service/dto"
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
		e.Error(500, err, err.Error())
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
// todo: remove redundant
func (e Patent) GetUserPatentsPages(c *gin.Context) {

	s := service.UserPatent{}
	s1 := service.Patent{}
	req := dto.UserPatentObject{}

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

	var count int64

	err = s.GetUserPatentIds(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var count2 int64
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	res, err := s1.GetPatentsByIds(ids, &count2)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(res, "查询成功")
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

	pid, PNM, desc, err := e.internalInsertIfAbsent(c)
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

	req := dto.NewUserPatentClaim(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), PNM, desc)

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

	pid, PNM, desc, err := e.internalInsertIfAbsent(c)

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

	req := dto.NewUserPatentFocus(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), PNM, desc)

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
	pid, pnm, _, err := e.internalInsertIfAbsent(c)
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

func (e Patent) internalInsertIfAbsent(c *gin.Context) (int, string, string, error) {
	ps := service.Patent{}
	req := dto.PatentReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		return 0, "", "", err
	}
	req.CreateBy = user.GetUserId(c)
	p, err := ps.InsertIfAbsent(&req)
	if err != nil {
		return 0, "", "", err
	}
	return p.PatentId, p.PNM, req.Desc, nil
}

// GetFocusPages
// @Summary 获取关注列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/focus [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Security Bearer
func (e Patent) GetFocusPages(c *gin.Context) {
	ups := service.UserPatent{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ups.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	list := make([]models.UserPatent, 0)
	userID := user.GetUserId(c)
	err = ups.GetFocusLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ps := service.Patent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.PatentPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}
	var count int64
	res, err := ps.GetPatentPagesByIds(ids, req, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	for i := range res {
		res[i].Desc = list[i].Desc
	}

	e.PageOK(res, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// FindFocusPages
// @Summary 搜索关注列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/focus/search [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e Patent) FindFocusPages(c *gin.Context) {
	ups := service.UserPatent{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ups.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	list := make([]models.UserPatent, 0)
	userID := user.GetUserId(c)
	err = ups.GetFocusLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ps := service.Patent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	query := c.Query("query")
	req := dto.FindPatentPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = query

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}
	var count int64
	res, err := ps.FindPatentPages(ids, req, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	for i := range res {
		res[i].Desc = list[i].Desc
	}

	e.PageOK(res, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// GetClaimPages
// @Summary 获取认领列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/claim [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Security Bearer
func (e Patent) GetClaimPages(c *gin.Context) {
	s := service.UserPatent{}

	userID := user.GetUserId(c)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.UserPatent, 0)
	err = s.GetClaimLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ps := service.Patent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.PatentPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize

	ids := make([]int, len(list))

	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	var count int64
	res, err := ps.GetPatentPagesByIds(ids, req, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	for i := range res {
		res[i].Desc = list[i].Desc
	}

	e.PageOK(res, int(count), req.PageIndex, req.PageSize, "查询成功")
}

// FindClaimPages
// @Summary 搜索认领专利
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/claim/search [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e Patent) FindClaimPages(c *gin.Context) {
	s := service.UserPatent{}

	userID := user.GetUserId(c)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.UserPatent, 0)
	err = s.GetClaimLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ps := service.Patent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	query := c.Query("query")
	req := dto.FindPatentPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = query

	ids := make([]int, len(list))

	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	var count int64
	res, err := ps.FindPatentPages(ids, req, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	for i := range res {
		res[i].Desc = list[i].Desc
	}

	e.PageOK(res, int(count), req.PageIndex, req.PageSize, "查询成功")
}

// DeleteFocus
// @Summary 取消关注
// @Description  取消关注
// @Tags 专利表
// @Param PNM query string false "专利PNM"
// @Router /api/v1/user-agent/patent/focus/{PNM}  [delete]
// @Security Bearer
func (e Patent) DeleteFocus(c *gin.Context) {
	var err error
	s := service.UserPatent{}
	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), -1, user.GetUserId(c), user.GetUserId(c), PNM, "")

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
// @Param PNM query string false "专利PNM"
// @Router /api/v1/user-agent/patent/claim/{PNM} [delete]
// @Security Bearer
func (e Patent) DeleteClaim(c *gin.Context) {
	var err error
	s := service.UserPatent{}

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentClaim(user.GetUserId(c), -1, user.GetUserId(c), user.GetUserId(c), PNM, "")

	err = e.MakeContext(c).
		MakeOrm().
		Bind(req). //修改&
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.RemoveClaim(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "取消认领成功")
}

// UpdateClaimDesc
// @Summary 更新认领专利备注
// @Description  更新认领专利备注
// @Tags 专利表
// @Param data body dto.PatentDescReq true "专利描述"
// @Router /api/v1/user-agent/patent/claim/{PNM}/desc [put]
// @Security Bearer
func (e Patent) UpdateClaimDesc(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.NewEmptyClaim()
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.UpdateUserPatentDesc(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "更新成功")
}

// UpdateFocusDesc
// @Summary 更新认领专利备注
// @Description  更新认领专利备注
// @Tags 专利表
// @Param data body dto.PatentDescReq true "专利描述"
// @Router /api/v1/user-agent/patent/focus/{PNM}/desc [put]
// @Security Bearer
func (e Patent) UpdateFocusDesc(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.NewEmptyFocus()
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.UpdateUserPatentDesc(req)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "更新成功")
}

//-----------------------------------------------foucs-graph--------------------------------------------------

// GetRelationGraphByFocus
// @Summary 获取关注专利的关系图谱
// @Description  获取关注专利的关系图谱
// @Tags 专利表
// @Router /api/v1/user-agent/patent/focus/graph/relation [get]
// @Security Bearer
func (e Patent) GetRelationGraphByFocus(c *gin.Context) {
	sp := service.Patent{}
	sup := service.UserPatent{}
	InventorGraph := models.Graph{}
	upList := make([]models.UserPatent, 0)
	var err error
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	userID := user.GetUserId(c)
	err = sup.GetFocusLists(userID, &upList)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, len(upList))
	for i := 0; i < len(upList); i++ {
		ids[i] = upList[i].PatentId
	}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sp.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var count int64
	listp, err := sp.GetPatentsByIds(ids, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	Inventors, Relations, err := sp.FindInventorsAndRelationsFromPatents(listp) //relations is an Upper Triangle
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	InventorGraph, err = sp.GetGraphByPatents(Inventors, Relations)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(InventorGraph, "查询成功")
}

// GetTechGraphByFocus
// @Summary 获取关注专利的技术图谱
// @Description  获取关注专利的技术图谱
// @Tags 专利表
// @Router /api/v1/user-agent/patent/focus/graph/tech [get]
// @Security Bearer
func (e Patent) GetTechGraphByFocus(c *gin.Context) {
	sp := service.Patent{}
	sup := service.UserPatent{}
	InventorGraph := models.Graph{}
	upList := make([]models.UserPatent, 0)
	var err error
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	userID := user.GetUserId(c)
	err = sup.GetFocusLists(userID, &upList)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, len(upList))
	for i := 0; i < len(upList); i++ {
		ids[i] = upList[i].PatentId
	}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sp.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var count int64
	listp, err := sp.GetPatentsByIds(ids, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	keyWords, Relations, err := sp.FindKeywordsAndRelationsFromPatents(listp) //relations is an Upper Triangle
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	InventorGraph, err = sp.GetGraphByPatents(keyWords, Relations)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(InventorGraph, "查询成功")
}

// ---------------------------------------------------patent--graph-------------------------------------------------------

// GetTheGraphByUserId
// @Summary 获取专利关系图
// @Description  获取专利关系图
// @Tags 专利表
// @Router /api/v1/user-agent/patent/relationship [get]
// @Security Bearer
func (e Patent) GetTheGraphByUserId(c *gin.Context) {
	//spp := service.PatentPackage{}
	sup := service.UserPatent{}
	su := aservice.SysUser{}
	sp := service.Patent{}
	gservice := service.Node{}
	//reqpp := dto.PackagePageGetReq{} //patent-package
	reqp := dto.PatentsIds{} //patents
	requ := adto.SysUserById{}
	//fmt.Println("get the line 471")
	//fmt.Println(c)
	var err error
	//reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
	//fmt.Println("get the line 474")
	//fmt.Println(reqpp.PackageId)
	//s := service.UserPatent{}

	requp := dto.UserPatentObject{}
	//reqp := dto.PatentsIds{}

	requp.UserId = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	listup1 := make([]models.UserPatent, 0)
	listp := make([]models.Patent, 0)

	var count int64 //not used
	//err = sup.GetUserFocusPatentIds(&requp, &listup1) //
	err = sup.GetFocusLists(&requp, &listup1, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64 //not used
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&reqp).
		MakeService(&sp.Service).
		Errors

	reqp.PatentIds = make([]int, len(listup1))
	for i := 0; i < len(listup1); i++ {
		reqp.PatentIds[i] = listup1[i].PatentId
	}

	err = sp.GetPageByIds(&reqp, &listp, &count2)
	fmt.Println("找到了所有的关注的patent")
	fmt.Println(listp)

	//err = e.MakeContext(c).
	//	MakeOrm().
	//	MakeService(&sup.Service).
	//	Errors
	////fmt.Println("get the line 480")
	//if err != nil {
	//	e.Logger.Error(err)
	//	e.Error(500, err, err.Error())
	//	return
	//}
	//reqpp.SetUpdateBy(user.GetUserId(c))
	//fmt.Println("get the line 486")
	//listpp := make([]models.PatentPackage, 0)
	//var count int64 //  not used
	//err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
	//fmt.Println(listpp)
	//fmt.Println(reqp)
	//for i := 0; i < len(listpp); i++ {
	//	fmt.Println(listpp[i].PatentId)
	//}
	//reqp.PatentIds = make([]int, len(listpp))
	//for i := 0; i < len(listpp); i++ {
	//	reqp.PatentIds[i] = listpp[i].PatentId
	//}
	//fmt.Println("get line 496")
	//1 := make([]models.Node, 0) //resultnode
	links := make([]models.Link, 0) //resultlink
	//listup2, members, usertimes := e.AddGraphNodeByReq(c, &nodes, reqp, 5)
	//------------------------------------already get the patents id  now get the users id
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	//fmt.Println("get line 505")
	listup2 := make([]models.UserPatent, 0) //√
	//fmt.Println(reqp.PatentIds)
	sup.GetFocusUsersByPatentId(&listup2, &reqp) //√
	//for i := 0; i < len(listup2); i++ {
	//	fmt.Println(listup2[i])
	//}
	//fmt.Println("the listup2 u success")
	//-----------------------------------already get the users id  now sort the users  and pick 8 users
	usertimes := make(map[int]int) // k-v is uid-times

	for i := 0; i < len(listup2); i++ { //建立map usertimes
		if usertimes[listup2[i].UserId] == 0 {
			usertimes[listup2[i].UserId] = 1
		} else {
			usertimes[listup2[i].UserId]++
		}

	}
	//fmt.Println("show the usertimes:")
	//fmt.Println(usertimes)
	usertimes1 := rankByWordCount(usertimes) //usertimes1有序的 key-value  uid-uidtimes 以times排序
	fmt.Println("show the usertimes1:")
	fmt.Println(usertimes1)
	//get 5th user
	var members int
	if len(usertimes1) < 75 {
		members = len(usertimes1)
	} else {
		members = 75
	}
	listu := make([]amodels.SysUser, members)
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&requ, nil).
		MakeService(&su.Service).
		Errors
	for i := 0; i < members; i++ {
		requ.Id = usertimes1[i].Key
		su.Get(&requ, &listu[i])
		fmt.Println("this is %d", i)
		fmt.Println(listu[i])
	}

	NodeList := make([]models.Node, 0)
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&gservice.Service).
		Errors
	//err = gservice.GetNodes(&NodeList)
	for i := 0; i < members; i++ {
		fmt.Println(i)
		err = gservice.GetNodes(&NodeList, i) //√
		NodeList[i].NodeValue = usertimes1[i].Value
		NodeList[i].NodeName = listu[i].Username
		//NodeList[i].NodeX = 0
		//NodeList[i].NodeY = 0
		//fmt.Println(NodeList[i])
		//nownode.GraphName =    应该加上
	}

	//有75个人的uid，package里的patent的pid  计算边的值，listpatent1为memberi的所包含的patent，relationExist为关系值
	for i := 0; i < members; i++ {
		for j := i + 1; j < members; j++ {
			//pantentssum := len(listpp)
			//var ispatent [2][pantentssum]bool
			//RelationExist, _ := sup.GetTwoUserRelationshipInThisPackage(&reqp, usertimes1[i].Key, usertimes1[j].Key)
			RelationExist := 0
			listpatent1 := make([]int, 0)
			listpatent2 := make([]int, 0)
			for z := 0; z < len(listup2); z++ {
				if listup2[z].UserId == usertimes1[i].Key {
					listpatent1 = append(listpatent1, listup2[z].PatentId)
				}
			}
			for z := 0; z < len(listup2); z++ {
				if listup2[z].UserId == usertimes1[j].Key {
					listpatent2 = append(listpatent2, listup2[z].PatentId)
				}
			}
			for z := 0; z < len(listpatent1); z++ {
				for z1 := 0; z1 < len(listpatent2); z1++ {
					if listpatent1[z] == listpatent2[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				var nowlink models.Link
				nowlink.Source = strconv.FormatInt((int64(usertimes1[i].Key)), 10)
				nowlink.Target = strconv.FormatInt((int64(usertimes1[j].Key)), 10)
				nowlink.Value = RelationExist
				links = append(links, nowlink)
				//fmt.Println(usertimes1[i].Key, usertimes1[j].Key)
				//fmt.Println(i, j, RelationExist, nowlink)
			}

		}
	} //√
	//fmt.Println(links)
	result := dto.GraphResult{}
	result.GetNodesAndLinks(&NodeList, &links)
	//fmt.Println(NodeList)
	//fmt.Println(links)
	//fmt.Println(result)
	e.GraphOK(NodeList, links, "查询成功")

	////get each 8 users of members users
	//for i := 0; i < members; i++ { //each node be added in nodes[]
	//	reqp2 := dto.PatentsIds{} //find patent listup2 of node[i]
	//	z := 0
	//	for j := 0; j < len(listup2); j++ {
	//		nodeid := usertimes[i].Key //get node id
	//		if listup2[j].UserId == nodeid {
	//			reqp2.PatentIds[z] = listup2[j].PatentId
	//			z++
	//		}
	//	}
	//	_, _, _ = e.AddGraphNodeByReq(c, &nodes, reqp2, i*14+5)
	//}

}

// GetTheGraphByUserId2
// @Summary 获取专利关系图2
// @Description  获取专利关系图2
// @Tags 专利表
// @Router /api/v1/user-agent/patent/relationship2 [get]
// @Security Bearer
func (e Patent) GetTheGraphByUserId2(c *gin.Context) {
	//spp := service.PatentPackage{}
	sup := service.UserPatent{}
	su := aservice.SysUser{}
	sp := service.Patent{}
	gservice := service.Node{}
	//reqpp := dto.PackagePageGetReq{} //patent-package
	reqp := dto.PatentsIds{} //patents
	requ := adto.SysUserById{}
	//fmt.Println("get the line 471")
	//fmt.Println(c)
	var err error
	//reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
	//fmt.Println("get the line 474")
	//fmt.Println(reqpp.PackageId)
	//s := service.UserPatent{}

	requp := dto.UserPatentObject{}
	//reqp := dto.PatentsIds{}

	requp.UserId = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	listup1 := make([]models.UserPatent, 0) //本用户所关注的专利
	listp := make([]models.Patent, 0)

	var count int64 //not used
	//err = sup.GetUserFocusPatentIds(&requp, &listup1) //
	err = sup.GetFocusLists(&requp, &listup1, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	var count2 int64 //not used
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&reqp).
		MakeService(&sp.Service).
		Errors

	reqp.PatentIds = make([]int, len(listup1))
	for i := 0; i < len(listup1); i++ {
		reqp.PatentIds[i] = listup1[i].PatentId
	}

	err = sp.GetPageByIds(&reqp, &listp, &count2)
	fmt.Println("找到了所有的关注的patent")
	fmt.Println(listp)

	fmt.Println("专利包所有的所有的patent的properties")
	for i := 0; i < len(listp); i++ {
		fmt.Println(listp[i].PatentProperties)
	}
	links := make([]models.Link, 0) //resultlink

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	//fmt.Println("get line 505")
	listup2 := make([]models.UserPatent, 0) //√          关注了专利(本用户所关注的)的所有用户数据
	//fmt.Println(reqp.PatentIds)
	sup.GetFocusUsersByPatentId(&listup2, &reqp) //√
	//for i := 0; i < len(listup2); i++ {
	//	fmt.Println(listup2[i])
	//}
	//fmt.Println("the listup2 u success")

	usertimes := make(map[int]int) // k-v is uid-times

	for i := 0; i < len(listup2); i++ { //建立map usertimes
		if usertimes[listup2[i].UserId] == 0 {
			usertimes[listup2[i].UserId] = 1
		} else {
			usertimes[listup2[i].UserId]++
		}

	}
	//fmt.Println("show the usertimes:")
	//fmt.Println(usertimes)
	usertimes1 := rankByWordCount(usertimes) //usertimes1有序的 自定义map(key-value)结构  uid-uidtimes 以times排序
	fmt.Println("show the usertimes1:")
	fmt.Println(usertimes1)
	//get 5th user
	var members int //找寻两两关系的结点的结点个数
	//var members2 int //最后输出的结点的个数
	if len(usertimes1) < 500 {
		members = len(usertimes1)
	} else {
		members = 500
	}
	fmt.Println("show the member:")
	fmt.Println(members)
	UserIsNode := make([]bool, members) //判断usertimes1中的uid是否在node中
	//--------------------------------------------------------------------------给强关系点赋值(最多10个)

	var StrongRelationNode int
	if members >= 100 {
		StrongRelationNode = 10
	} else {
		StrongRelationNode = members / 10
	}
	NodeList := make([]models.Node, StrongRelationNode) //需要输出的node列表,初始值为10
	for i := 0; i < StrongRelationNode; i++ {           //设置初始颜色类型
		NodeList[i].NodeCategory = i
		NodeList[i].NodeId = strconv.FormatInt(int64(usertimes1[i].Key), 10)
		NodeList[i].NodeValue = usertimes1[i].Value
		UserIsNode[i] = true
		//NodeList[i].NodeName =
	}
	fmt.Println("show the StrongNodelist:")
	fmt.Println(NodeList)
	//-------------------------------------------------------------------------------统计每个用户的专利
	userspatents := make([]models.OneUserPatents, members) //已排序的usertimes中每个用户的专利（patentid数组）
	for i := 0; i < len(listup2); i++ {
		for j := 0; j < members; j++ {
			if listup2[i].UserId == usertimes1[j].Key {
				userspatents[j].Patentsid = append(userspatents[j].Patentsid, listup2[i].PatentId)
				break
			}
		}
	}
	fmt.Println("show the userspatents:")
	fmt.Println(userspatents)

	useruserrelation1 := make(map[int]int) //前十结点的两两关系
	first10 := StrongRelationNode          //第一次处理的点数
	firstlinks := first10 * 2              //第一次处理的边数
	//-------------------------------------------------------------------------------处理10两两结点关系
	for i := 0; i < first10; i++ {
		for j := i + 1; j < first10; j++ {
			RelationExist := 0
			for z := 0; z < len(userspatents[i].Patentsid); z++ {
				for z1 := 0; z1 < len(userspatents[j].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelation1[i*10+j] = RelationExist
			}
		}
	}
	useruserrelation2 := rankByWordCount(useruserrelation1) //给边排序
	fmt.Println("show the useruserrelation2:")
	fmt.Println(useruserrelation2)
	for i := 0; i < minresult(firstlinks, len(useruserrelation2)); i++ {
		var nowlink models.Link
		nowlink.Source = strconv.FormatInt(int64(usertimes1[useruserrelation2[i].Key/10].Key), 10)
		nowlink.Target = strconv.FormatInt(int64(usertimes1[useruserrelation2[i].Key%10].Key), 10)
		nowlink.Value = useruserrelation2[i].Value
		links = append(links, nowlink)
	}

	fmt.Println("show the Nodelist1:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList1:")
	fmt.Println(links)
	fmt.Println("------------------------------------------------------------------------------------------")
	//--------------------------------------------------------------------------------------------------
	useruserrelation3 := make(map[int]int) //前10，490结点的两两关系
	//MaxRelationNode := 5
	secondLinks := 200
	ExtendNodeTime := make([]int, members) //strongNode可扩展的点 和 regularNode可扩展的边
	//----------------------------------------------------------------处理10,490两两结点的关系

	//有最多500个人的uid，package里的patent的pid   ,listup是查出来的所有的patent-user关系
	for i := 1; i < first10; i++ { //这里的i和j不会重复  第一个点单独处理
		for j := first10; j < members; j++ {
			RelationExist := 0
			for z := 0; z < len(userspatents[i].Patentsid); z++ { //平均关注的patent不多的话复杂度不高，可以用map优化(后续做)
				for z1 := 0; z1 < len(userspatents[j].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelation3[i*500+j] = RelationExist
			}
		}
	} //√
	useruserrelation4 := rankByWordCount(useruserrelation3) // source*500+target---重复次数    key value     source,target为再usertimes中的排序序号
	NodelistIdToTimeLst := make(map[string]int)             //key-value   点id-----在usertimes1,usersPatents中的排序
	fmt.Println("show the useruserrelation4:")
	fmt.Println(useruserrelation4)

	for i := 0; i < minresult(secondLinks, len(useruserrelation4)); i++ {
		source := useruserrelation4[i].Key / 500
		target := useruserrelation4[i].Key % 500
		if ExtendNodeTime[source] >= 5 { //strongnode扩展超过5个点
			continue
		} else {
			if UserIsNode[target] == false {
				//fmt.Println("show the nextnode:")
				//fmt.Println(i)
				//fmt.Println(target)
				UserIsNode[target] = true
				ExtendNodeTime[source]++
				var nowlink models.Link
				var nowNode models.Node
				nowNode.NodeCategory = NodeList[source].NodeCategory //这里NodeList的source是和usertime1的序号相同的
				nowNode.NodeId = strconv.FormatInt(int64(usertimes1[target].Key), 10)
				NodelistIdToTimeLst[nowNode.NodeId] = target
				nowlink.Source = strconv.FormatInt(int64(usertimes1[source].Key), 10)
				nowlink.Target = strconv.FormatInt(int64(usertimes1[target].Key), 10)
				nowlink.Value = useruserrelation4[i].Value
				fmt.Println("show the nextnode and link:")
				fmt.Println(nowNode)
				fmt.Println(nowlink)
				links = append(links, nowlink)       //边增加
				NodeList = append(NodeList, nowNode) //点增加
			}
		}
	}
	fmt.Println("show the Nodelist2:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList2:")
	fmt.Println(links)
	// --------------------------------------------------建立后续点的关系表   建议用相同的边排序算法进行计算
	//regularNodetime := make([]int, len(NodeList)-StrongRelationNode)
	//for i := StrongRelationNode; i < len(NodeList); i++ {
	//	for j := i+1; j <len(NodeList) ; j++ {
	//		if
	//		for z := 0; z < len(userspatents[NodelistIdToTimeLst[i]].Patentsid); z++ {
	//			for z1 := 0; z1 < ; z1++ {
	//
	//			}
	//		}
	//	}
	//}
	useruserrelation5 := make(map[int]int)
	for i := first10; i < len(NodeList); i++ { //后续regular点加边
		for j := i + 1; j < len(NodeList); j++ {
			RelationExist := 0
			iToUserspatentsPosition := NodelistIdToTimeLst[NodeList[i].NodeId]
			jToUserspatentsPosition := NodelistIdToTimeLst[NodeList[j].NodeId]
			for z := 0; z < len(userspatents[iToUserspatentsPosition].Patentsid); z++ { //平均关注的patent不多的话复杂度不高，可以用map优化(后续做)
				for z1 := 0; z1 < len(userspatents[jToUserspatentsPosition].Patentsid); z1++ {
					if userspatents[iToUserspatentsPosition].Patentsid[z] == userspatents[jToUserspatentsPosition].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelation5[i*500+j] = RelationExist //根据NodeList中的位置为key进行的排序(按relationExist排序)
			}

		}
	}

	useruserrelation6 := rankByWordCount(useruserrelation5) //第三次边的关系
	thirdlinks := 50                                        //第三次边的数量
	for i := 0; i < minresult(thirdlinks, len(useruserrelation6)); i++ {
		source := useruserrelation6[i].Key / 500 //source在NodeList中的位置
		target := useruserrelation6[i].Key % 500

		if ExtendNodeTime[source] >= 3 { //strongnode扩展超过5个点
			continue
		} else {
			ExtendNodeTime[source]++ //NodeList中的source位置
			var nowlink models.Link
			nowlink.Source = NodeList[source].NodeId
			nowlink.Target = NodeList[target].NodeId
			nowlink.Value = useruserrelation6[i].Value
			links = append(links, nowlink) //边增加
		}
	}
	fmt.Println("show the Nodelist3:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList3:")
	fmt.Println(links)
	//-----------------------------------建立第一个点的关系  先不建立了偷个懒

	//------------------------补全要显示的点的信息

	//members = len(usertimes1)                 //先不限制显示的user对象的数量
	listu := make([]amodels.SysUser, len(NodeList)) //需要显示的user对象
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&requ, nil).
		MakeService(&su.Service).
		Errors
	for i := 0; i < len(NodeList); i++ { //找名字
		//int, err := strconv.Atoi(string)
		requ.Id, err = strconv.Atoi(NodeList[i].NodeId)
		su.Get(&requ, &listu[i]) //查找需要显示的user对象
		fmt.Println("this is %d", i)
		fmt.Println(listu[i])
	}

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&gservice.Service).
		Errors
	//err = gservice.GetNodes(&NodeList)
	max := 0
	min := 100000
	NodeList[0].NodeValue = usertimes1[0].Value //先把最大的点  操作一下
	NodeList[0].NodeSymbolizeSize = 60
	NodeList[0].NodeName = listu[0].Username //点的名字
	for i := 1; i < first10; i++ {           //strongNode 取value（有多少重复的patent（time））
		NodeList[i].NodeValue = usertimes1[i].Value
		fmt.Println("show the value:")
		fmt.Println(NodeList[i].NodeValue)
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	for i := first10; i < len(NodeList); i++ { //regularNode 取value（有多少重复的patent（time））
		NodeList[i].NodeValue = usertimes1[NodelistIdToTimeLst[NodeList[i].NodeId]].Value
		fmt.Println("show the value:")
		fmt.Println(NodeList[i].NodeValue)
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	fmt.Println("show the max:")
	fmt.Println(max)
	fmt.Println("show the min:")
	fmt.Println(min)

	for i := 1; i < len(NodeList); i++ {
		//err = gservice.GetNodes(&NodeList, i) //√
		fmt.Println(NodeList[i].NodeValue)
		fmt.Println(float32(float32(NodeList[i].NodeValue*60) / float32(maxresult((max), 1))))
		NodeList[i].NodeSymbolizeSize = float32(float32(NodeList[i].NodeValue*60) / float32(maxresult((max), 1)))
		NodeList[i].NodeName = listu[i].Username //点的名字
		//NodeList[i].NodeId = strconv.FormatInt(int64(usertimes1[i].Key), 10) //点的id(db中的user-id)
		//fmt.Println(NodeList[i])
		//nownode.GraphName =    应该加上
	}
	fmt.Println("show the Nodelist4:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList4:")
	fmt.Println(links)
	//-----------------规范输出格式
	result := dto.GraphResult{}
	result.GetNodesAndLinks(&NodeList, &links)
	//fmt.Println(NodeList)
	//fmt.Println(links)
	//fmt.Println(result)
	e.GraphOK(NodeList, links, "查询成功")

}

// GetTheGraphByUserId3
// @Summary 获取专利发明人关系图
// @Description  获取专利发明人关系图
// @Tags 专利表
// @Router /api/v1/user-agent/patent/relationship3 [get]
// @Security Bearer
func (e Patent) GetTheGraphByUserId3(c *gin.Context) {
	//spp := service.PatentPackage{}
	sup := service.UserPatent{}
	su := aservice.SysUser{}
	sp := service.Patent{}
	gservice := service.Node{}
	//reqpp := dto.PackagePageGetReq{} //patent-package
	reqp := dto.PatentsIds{} //patents
	requ := adto.SysUserById{}
	//fmt.Println("get the line 471")
	//fmt.Println(c)
	var err error
	//reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
	//fmt.Println("get the line 474")
	//fmt.Println(reqpp.PackageId)
	//s := service.UserPatent{}

	requp := dto.UserPatentObject{}
	//reqp := dto.PatentsIds{}

	requp.UserId = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	listup1 := make([]models.UserPatent, 0) //本用户所关注的专利
	listp := make([]models.Patent, 0)

	var count int64 //not used
	//err = sup.GetUserFocusPatentIds(&requp, &listup1) //
	err = sup.GetFocusLists(&requp, &listup1, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	var count2 int64 //not used
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&reqp).
		MakeService(&sp.Service).
		Errors

	reqp.PatentIds = make([]int, len(listup1))
	for i := 0; i < len(listup1); i++ {
		reqp.PatentIds[i] = listup1[i].PatentId
	}

	err = sp.GetPageByIds(&reqp, &listp, &count2)
	fmt.Println("找到了所有的关注的patent")
	fmt.Println(listp)

	links := make([]models.Link, 0) //resultlink

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	//fmt.Println("get line 505")
	listup2 := make([]models.UserPatent, 0) //√          关注了专利(本用户所关注的)的所有用户数据
	//fmt.Println(reqp.PatentIds)
	sup.GetFocusUsersByPatentId(&listup2, &reqp) //√
	//for i := 0; i < len(listup2); i++ {
	//	fmt.Println(listup2[i])
	//}
	//fmt.Println("the listup2 u success")

	usertimes := make(map[int]int) // k-v is uid-times

	for i := 0; i < len(listup2); i++ { //建立map usertimes
		if usertimes[listup2[i].UserId] == 0 {
			usertimes[listup2[i].UserId] = 1
		} else {
			usertimes[listup2[i].UserId]++
		}

	}
	//fmt.Println("show the usertimes:")
	//fmt.Println(usertimes)
	usertimes1 := rankByWordCount(usertimes) //usertimes1有序的 自定义map(key-value)结构  uid-uidtimes 以times排序
	fmt.Println("show the usertimes1:")
	fmt.Println(usertimes1)
	//get 5th user
	var members int //找寻两两关系的结点的结点个数
	//var members2 int //最后输出的结点的个数
	if len(usertimes1) < 500 {
		members = len(usertimes1)
	} else {
		members = 500
	}
	fmt.Println("show the member:")
	fmt.Println(members)
	UserIsNode := make([]bool, members) //判断usertimes1中的uid是否在node中
	//--------------------------------------------------------------------------给强关系点赋值(最多10个)

	var StrongRelationNode int
	if members >= 100 {
		StrongRelationNode = 10
	} else {
		StrongRelationNode = members / 10
	}
	NodeList := make([]models.Node, StrongRelationNode) //需要输出的node列表,初始值为10
	for i := 0; i < StrongRelationNode; i++ {           //设置初始颜色类型
		NodeList[i].NodeCategory = i
		NodeList[i].NodeId = strconv.FormatInt(int64(usertimes1[i].Key), 10)
		NodeList[i].NodeValue = usertimes1[i].Value
		UserIsNode[i] = true
		//NodeList[i].NodeName =
	}
	fmt.Println("show the StrongNodelist:")
	fmt.Println(NodeList)
	//-------------------------------------------------------------------------------统计每个用户的专利
	userspatents := make([]models.OneUserPatents, members) //已排序的usertimes中每个用户的专利（patentid数组）
	for i := 0; i < len(listup2); i++ {
		for j := 0; j < members; j++ {
			if listup2[i].UserId == usertimes1[j].Key {
				userspatents[j].Patentsid = append(userspatents[j].Patentsid, listup2[i].PatentId)
				break
			}
		}
	}
	fmt.Println("show the userspatents:")
	fmt.Println(userspatents)

	useruserrelation1 := make(map[int]int) //前十结点的两两关系
	first10 := StrongRelationNode          //第一次处理的点数
	firstlinks := first10 * 2              //第一次处理的边数
	//-------------------------------------------------------------------------------处理10两两结点关系
	for i := 0; i < first10; i++ {
		for j := i + 1; j < first10; j++ {
			RelationExist := 0
			for z := 0; z < len(userspatents[i].Patentsid); z++ {
				for z1 := 0; z1 < len(userspatents[j].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelation1[i*10+j] = RelationExist
			}
		}
	}
	useruserrelation2 := rankByWordCount(useruserrelation1) //给边排序
	fmt.Println("show the useruserrelation2:")
	fmt.Println(useruserrelation2)
	for i := 0; i < minresult(firstlinks, len(useruserrelation2)); i++ {
		var nowlink models.Link
		nowlink.Source = strconv.FormatInt(int64(usertimes1[useruserrelation2[i].Key/10].Key), 10)
		nowlink.Target = strconv.FormatInt(int64(usertimes1[useruserrelation2[i].Key%10].Key), 10)
		nowlink.Value = useruserrelation2[i].Value
		links = append(links, nowlink)
	}

	fmt.Println("show the Nodelist1:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList1:")
	fmt.Println(links)
	fmt.Println("------------------------------------------------------------------------------------------")
	//--------------------------------------------------------------------------------------------------
	useruserrelation3 := make(map[int]int) //前10，490结点的两两关系
	//MaxRelationNode := 5
	secondLinks := 200
	ExtendNodeTime := make([]int, members) //strongNode可扩展的点 和 regularNode可扩展的边
	//----------------------------------------------------------------处理10,490两两结点的关系

	//有最多500个人的uid，package里的patent的pid   ,listup是查出来的所有的patent-user关系
	for i := 1; i < first10; i++ { //这里的i和j不会重复  第一个点单独处理
		for j := first10; j < members; j++ {
			RelationExist := 0
			for z := 0; z < len(userspatents[i].Patentsid); z++ { //平均关注的patent不多的话复杂度不高，可以用map优化(后续做)
				for z1 := 0; z1 < len(userspatents[j].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelation3[i*500+j] = RelationExist
			}
		}
	} //√
	useruserrelation4 := rankByWordCount(useruserrelation3) // source*500+target---重复次数    key value     source,target为再usertimes中的排序序号
	NodelistIdToTimeLst := make(map[string]int)             //key-value   点id-----在usertimes1,usersPatents中的排序
	fmt.Println("show the useruserrelation4:")
	fmt.Println(useruserrelation4)

	for i := 0; i < minresult(secondLinks, len(useruserrelation4)); i++ {
		source := useruserrelation4[i].Key / 500
		target := useruserrelation4[i].Key % 500
		if ExtendNodeTime[source] >= 5 { //strongnode扩展超过5个点
			continue
		} else {
			if UserIsNode[target] == false {
				//fmt.Println("show the nextnode:")
				//fmt.Println(i)
				//fmt.Println(target)
				UserIsNode[target] = true
				ExtendNodeTime[source]++
				var nowlink models.Link
				var nowNode models.Node
				nowNode.NodeCategory = NodeList[source].NodeCategory //这里NodeList的source是和usertime1的序号相同的
				nowNode.NodeId = strconv.FormatInt(int64(usertimes1[target].Key), 10)
				NodelistIdToTimeLst[nowNode.NodeId] = target
				nowlink.Source = strconv.FormatInt(int64(usertimes1[source].Key), 10)
				nowlink.Target = strconv.FormatInt(int64(usertimes1[target].Key), 10)
				nowlink.Value = useruserrelation4[i].Value
				fmt.Println("show the nextnode and link:")
				fmt.Println(nowNode)
				fmt.Println(nowlink)
				links = append(links, nowlink)       //边增加
				NodeList = append(NodeList, nowNode) //点增加
			}
		}
	}
	fmt.Println("show the Nodelist2:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList2:")
	fmt.Println(links)
	// --------------------------------------------------建立后续点的关系表   建议用相同的边排序算法进行计算
	//regularNodetime := make([]int, len(NodeList)-StrongRelationNode)
	//for i := StrongRelationNode; i < len(NodeList); i++ {
	//	for j := i+1; j <len(NodeList) ; j++ {
	//		if
	//		for z := 0; z < len(userspatents[NodelistIdToTimeLst[i]].Patentsid); z++ {
	//			for z1 := 0; z1 < ; z1++ {
	//
	//			}
	//		}
	//	}
	//}
	useruserrelation5 := make(map[int]int)
	for i := first10; i < len(NodeList); i++ { //后续regular点加边
		for j := i + 1; j < len(NodeList); j++ {
			RelationExist := 0
			iToUserspatentsPosition := NodelistIdToTimeLst[NodeList[i].NodeId]
			jToUserspatentsPosition := NodelistIdToTimeLst[NodeList[i].NodeId]
			for z := 0; z < len(userspatents[iToUserspatentsPosition].Patentsid); z++ { //平均关注的patent不多的话复杂度不高，可以用map优化(后续做)
				for z1 := 0; z1 < len(userspatents[jToUserspatentsPosition].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelation5[i*500+j] = RelationExist //根据NodeList中的位置为key进行的排序(按relationExist排序)
			}

		}
	}

	useruserrelation6 := rankByWordCount(useruserrelation5) //第三次边的关系
	thirdlinks := 50                                        //第三次边的数量
	for i := 0; i < minresult(thirdlinks, len(useruserrelation6)); i++ {
		source := useruserrelation6[i].Key / 500 //source在NodeList中的位置
		target := useruserrelation6[i].Key % 500

		if ExtendNodeTime[source] >= 3 { //strongnode扩展超过5个点
			continue
		} else {
			ExtendNodeTime[source]++ //NodeList中的source位置
			var nowlink models.Link
			nowlink.Source = NodeList[source].NodeId
			nowlink.Target = NodeList[target].NodeId
			nowlink.Value = useruserrelation6[i].Value
			links = append(links, nowlink) //边增加
		}
	}
	fmt.Println("show the Nodelist3:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList3:")
	fmt.Println(links)
	//-----------------------------------建立第一个点的关系  先不建立了偷个懒

	//------------------------补全要显示的点的信息

	//members = len(usertimes1)                 //先不限制显示的user对象的数量
	listu := make([]amodels.SysUser, len(NodeList)) //需要显示的user对象
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&requ, nil).
		MakeService(&su.Service).
		Errors
	for i := 0; i < len(NodeList); i++ { //找名字
		//int, err := strconv.Atoi(string)
		requ.Id, err = strconv.Atoi(NodeList[i].NodeId)
		su.Get(&requ, &listu[i]) //查找需要显示的user对象
		fmt.Println("this is %d", i)
		fmt.Println(listu[i])
	}

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&gservice.Service).
		Errors
	//err = gservice.GetNodes(&NodeList)
	max := 0
	min := 100000
	NodeList[0].NodeValue = usertimes1[0].Value //先把最大的点  操作一下
	NodeList[0].NodeSymbolizeSize = 60
	NodeList[0].NodeName = listu[0].Username //点的名字
	for i := 1; i < first10; i++ {           //strongNode 取value（有多少重复的patent（time））
		NodeList[i].NodeValue = usertimes1[i].Value
		fmt.Println("show the value:")
		fmt.Println(NodeList[i].NodeValue)
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	for i := first10; i < len(NodeList); i++ { //regularNode 取value（有多少重复的patent（time））
		NodeList[i].NodeValue = usertimes1[NodelistIdToTimeLst[NodeList[i].NodeId]].Value
		fmt.Println("show the value:")
		fmt.Println(NodeList[i].NodeValue)
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	fmt.Println("show the max:")
	fmt.Println(max)
	fmt.Println("show the min:")
	fmt.Println(min)

	for i := 1; i < len(NodeList); i++ {
		//err = gservice.GetNodes(&NodeList, i) //√
		fmt.Println(NodeList[i].NodeValue)
		fmt.Println(float32(float32(NodeList[i].NodeValue*60) / float32(maxresult((max), 1))))
		NodeList[i].NodeSymbolizeSize = float32(float32(NodeList[i].NodeValue*60) / float32(maxresult((max), 1)))
		NodeList[i].NodeName = listu[i].Username //点的名字
		//NodeList[i].NodeId = strconv.FormatInt(int64(usertimes1[i].Key), 10) //点的id(db中的user-id)
		//fmt.Println(NodeList[i])
		//nownode.GraphName =    应该加上
	}
	fmt.Println("show the Nodelist4:")
	fmt.Println(NodeList)
	fmt.Println("show the LinkList4:")
	fmt.Println(links)
	//-----------------规范输出格式
	result := dto.GraphResult{}
	result.GetNodesAndLinks(&NodeList, &links)
	//fmt.Println(NodeList)
	//fmt.Println(links)
	//fmt.Println(result)
	e.GraphOK(NodeList, links, "查询成功")

}

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
