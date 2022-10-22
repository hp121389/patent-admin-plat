package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"net/http"
	"strconv"
)

type Package struct {
	api.Api
}

// GetPage
// @Summary 列表专利包信息数据
// @Description 获取JSON
// @Tags 专利包
// @Param packageName query string false "packageName"
// @Router /api/v1/user-agent/package [get]
// @Security Bearer
func (e Package) GetPage(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageGetPageReq{}
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

	list := make([]models.Package, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取专利包
// @Description 获取JSON
// @Tags 专利包
// @Param packageId path int true "专利包编码"
// @Router /api/v1/user-agent/package/{packageId} [get]
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
	pid, err := s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	fmt.Println("pid已查出:", pid)
	err = e.InsertUserPackage(c, pid)
	//e.InsertUserPackage(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

func (e *Package) InsertUserPackage(c *gin.Context, pid int) error {

	s := service.Package{}
	//req := dto.UserPackageInsertReq{}
	//req.UserId = user.GetUserId(c)
	var err error
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return err
	}
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	// 设置创建人
	//req.SetCreateBy(user.GetUserId(c))

	req := dto.NewUserPackageInsert(user.GetUserId(c), pid)
	req.SetCreateBy(user.GetUserId(c))
	//err = s.InsertUserPackage(&req,user.GetUserId(c),pid)
	//fmt.Println("1pid，uid已查出:", pid, uid)

	err = s.InsertUserPackage(req)
	//fmt.Println("2pid，uid已查出:", pid, uid)
	if req.PackageId == 0 {
		e.Logger.Error(err)
		e.Error(404, err, "您输入的专利包id不存在！")
		return err
	}

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return err
	}
	fmt.Println("插入成功")
	e.OK(req, "创建成功")
	return nil
}

// Update
// @Summary 修改专利包数据
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageInsertReq true "body"
// @Router /api/v1/user-agent/package [put]
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

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除专利包数据
// @Description 删除数据
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package [delete]
// @Param data body dto.ObjectById true "body"
// @Security Bearer
func (e Package) Delete(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageById{}
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

type UserPackage struct {
	api.Api
}

// GetPage
// @Summary 列表用户专利包信息数据
// @Description 获取JSON
// @Tags 专利包
// @Router /api/v1/user-agent/user-package [get]
// @Security Bearer
func (e UserPackage) GetPage(c *gin.Context) {
	s := service.UserPackage{}
	req := dto.UserPackageGetPageReq{}
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

	list := make([]models.UserPackage, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// GetPackageByUserId
// @Summary 获得该UserId的专利包
// @Tags 专利包
// @Router /api/v1/user-agent/user-package/{userId} [get]
// @Security Bearer
func (e UserPackage) GetPackageByUserId(c *gin.Context) { //gin框架里的上下文

	s := service.UserPackage{}         //service中查询或者返回的结果赋值给s变量
	req := dto.UserPackageGetPageReq{} //被绑定的数据
	req1 := dto.PackagesByIdsForRelationshipUsers{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	list := make([]models.UserPackage, 0)
	list1 := make([]models.Package, 0)
	var count int64
	err = s.GetPackageIdsByUserId(&req, &list, &count)
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
	req1.PackageIds = make([]int, len(list))
	for i := 0; i < len(list); i++ {
		req1.PackageIds[i] = list[i].PackageId
	}
	err = s.GetPackagePagesByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list1, "查询成功")
}

// DeleteUserPackage
// @Summary 输入专利包id删除专利表
// @Description  输入专利包id删除专利表
// @Tags 专利包
// @Param PackageId query string false "专利包ID"
// @Router /api/v1/user-agent/user-package/{package_id} [delete]
// @Security Bearer
func (e UserPackage) DeleteUserPackage(c *gin.Context) {
	s := service.UserPackage{}
	req := dto.UserPackageObject{}
	req.UserId = user.GetUserId(c)

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req). //在这一步传入request数据
		MakeService(&s.Service).
		Errors

	fmt.Println(req.PackageId)
	fmt.Println(req.UserId)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveRelationship(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "删除成功")
}
