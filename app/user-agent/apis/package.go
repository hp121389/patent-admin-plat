package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/user-agent/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

type Package struct {
	api.Api
}

// ListByCurrentUser
// @Summary 获取当前用户专利包列表
// @Description 获取JSON
// @Tags 专利包
// @Router /api/v1/user-agent/package [get]
// @Security Bearer
func (e Package) ListByCurrentUser(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageListReq{}
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

	req.UserId = user.GetUserId(c)

	list := make([]models.Package, 0)

	err = s.ListByUserId(&req, &list)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(list, "查询成功")
}

// Find
// @Summary 搜索当前用户专利包列表
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/package/search [get]
// @Param data body dto.FindPatentPagesReq true "搜索参数"
// @Security Bearer
func (e Package) Find(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageFindReq{}
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

	req.UserId = user.GetUserId(c)

	list := make([]models.Package, 0)

	err = s.FindForCurrentUser(&req, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(list, "查询成功")
}

// Get
// @Summary 获取专利包
// @Description 获取JSON
// @Tags 专利包
// @Param packageId path int true "package_id"
// @Router /api/v1/user-agent/package/{package_id} [get]
// @Security Bearer
func (e Package) Get(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageById{}
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
	var object models.Package
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建专利包
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageInsertReq true "专利包数据"
// @Router /api/v1/user-agent/package [post]
// @Security Bearer
func (e Package) Insert(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageInsertReq{}
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

	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改专利包数据
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageUpdateReq true "body"
// @Router /api/v1/user-agent/package/{package_id} [put]
// @Security Bearer
func (e Package) Update(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageUpdateReq{}
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

	if pid, err := strconv.Atoi(c.Param("id")); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	} else {
		req.PackageId = pid
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(nil, "更新成功")
}

// Delete
// @Summary 删除专利包
// @Description 删除专利包
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id} [delete]
// @Security Bearer
func (e Package) Delete(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageById{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.Id, err = strconv.Atoi(c.Param("id"))
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
	e.OK(req.GetId(), "删除成功")
}

//----------------------------------------patent-package---------------------------------------

// todo: please modify the swagger comment

// GetPackagePatents
// @Summary 获取指定专利包中的专利列表
// @Description 获取指定专利包中的专利列表
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id}/patent [get]
// @Security Bearer
func (e Package) GetPackagePatents(c *gin.Context) {

	s := service.PatentPackage{}
	s1 := service.Patent{}
	req := dto.PackagePageGetReq{}

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

	req.PackageId, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	var list []models.PatentPackage
	var count int64

	err = s.GetPatentIdByPackageId(&req, &list, &count)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s1.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	res, err := s1.GetPatentsByIds(ids, &count2)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	for i := range res {
		res[i].Desc = list[i].Desc
	}

	e.PageOK(res, int(count2), req.GetPageIndex(), req.GetPageSize(), "查询成功")

}

// IsPatentInPackage
// @Summary 查询专利是否已在专利包中
// @Description 查询专利是否已在专利包中
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id}/isExist [get]
// @Security Bearer
func (e Package) IsPatentInPackage(c *gin.Context) {
	var err error

	pps := service.PatentPackage{}
	req := dto.PatentPackageReq{}

	req.PNM = c.Param("PNM")

	req.PackageId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.CreateBy = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	existed, err := pps.IsPatentInPackage(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(&dto.IsPatentInPackageResp{Existed: existed}, "查询成功")
}

// InsertPackagePatent
// @Summary 将专利加入专利包
// @Description  将专利加入专利包
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "专利表数据"
// @Router /api/v1/user-agent/package/{package_id}/patent [post]
// @Security Bearer
func (e Package) InsertPackagePatent(c *gin.Context) {
	var err error
	pps := service.PatentPackage{}
	req := dto.PatentPackageReq{}

	ps := service.Patent{}
	patentReq := dto.PatentReq{}
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&patentReq).
		MakeService(&ps.Service).
		Errors
	patentReq.CreateBy = user.GetUserId(c)
	p, err := ps.InsertIfAbsent(&patentReq)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.PatentId = p.PatentId
	req.PNM = p.PNM
	req.PackageId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.CreateBy = user.GetUserId(c)
	req.Desc = patentReq.Desc

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = pps.InsertPatentPackage(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "创建成功")
}

// DeletePackagePatent
// @Summary 删除专利包专利
// @Description  删除专利包专利
// @Tags 专利包
// @Param PatentId query string false "专利ID"
// @Param PackageId query string false "专利包ID"
// @Router /api/v1/user-agent/package/{package_id}/patent/{PNM} [delete]
// @Security Bearer
func (e Package) DeletePackagePatent(c *gin.Context) {
	s := service.PatentPackage{}
	req := dto.PackagePageGetReq{}
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	packageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PackageId = packageId

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.RemovePackagePatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.PackageBack, "删除成功")
}

// UpdatePackagePatentDesc
// @Summary 更新专利包专利描述
// @Description  更新专利包专利描述
// @Tags 专利包
// @Param data body dto.PatentDescReq true "专利描述"
// @Router /api/v1/user-agent/package/{package_id}/patent/{PNM}/desc [put]
// @Security Bearer
func (e Package) UpdatePackagePatentDesc(c *gin.Context) {
	s := service.PatentPackage{}
	req := dto.PatentDescReq{}
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

	packageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PackageID = packageId

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.UpdatePackagePatentDesc(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "更新成功")
}

//---------------------------------------------------patent--graph-------------------------------------------------------

// GetRelationGraphByPackage
// @Summary 获取专利包关系图谱
// @Description  获取专利包关系图谱
// @Tags 专利包
// @Router /api/v1/user-agent/package/{packageId}/graph/relation [get]
// @Security Bearer
func (e Package) GetRelationGraphByPackage(c *gin.Context) {
	sp := service.Patent{}
	spp := service.PatentPackage{}
	reqpp := dto.PackagePageGetReq{}
	InventorGraph := models.Graph{}
	var err error
	reqpp.PackageId, err = strconv.Atoi(c.Param("id"))
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&spp.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reqpp.SetUpdateBy(user.GetUserId(c))
	listpp := make([]models.PatentPackage, 0)
	var count int64
	err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	ids := make([]int, len(listpp))
	for i := 0; i < len(listpp); i++ {
		ids[i] = listpp[i].PatentId
	}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sp.Service).
		Errors

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
		return
	}
	e.OK(InventorGraph, "查询成功")
}

// GetTechGraphByPackage
// @Summary 获取专利包技术图谱
// @Description  获取专利包技术图谱
// @Tags 专利包
// @Router /api/v1/user-agent/package/{packageId}/graph/tech [get]
// @Security Bearer
func (e Package) GetTechGraphByPackage(c *gin.Context) {
	sp := service.Patent{}
	spp := service.PatentPackage{}
	reqpp := dto.PackagePageGetReq{}
	InventorGraph := models.Graph{}
	var err error
	reqpp.PackageId, err = strconv.Atoi(c.Param("id"))
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&spp.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reqpp.SetUpdateBy(user.GetUserId(c))
	listpp := make([]models.PatentPackage, 0)
	var count int64
	err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
	ids := make([]int, len(listpp))
	for i := 0; i < len(listpp); i++ {
		ids[i] = listpp[i].PatentId
	}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sp.Service).
		Errors

	listp, err := sp.GetPatentsByIds(ids, &count)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	keyWords, Relations, err := sp.FindKeywordsAndRelationsFromPatents(listp) //relations is an Upper Triangle
	if err != nil {
		e.Logger.Error(err)
		return
	}
	InventorGraph, err = sp.GetGraphByPatents(keyWords, Relations)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(InventorGraph, "查询成功")
}
