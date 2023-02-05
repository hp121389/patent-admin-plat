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
		e.Error(500, err, err.Error())
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
