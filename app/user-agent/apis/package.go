package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	amodels "go-admin/app/admin/models"
	aservice "go-admin/app/admin/service"
	adto "go-admin/app/admin/service/dto"
	"go-admin/app/user-agent/models"
	"net/http"
	"sort"
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

//// GetPage
//// @Summary 列表专利包信息数据
//// @Description 获取JSON
//// @Tags 专利包
//// @Param packageName query string false "packageName"
//// @Router /api/v1/package [get]
//// @Security Bearer
//func (e Package) GetPage(c *gin.Context) {
//	s := service.Package{}
//	req := dtos.PackageGetPageReq{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	//数据权限检查
//	//p := actions.GetPermissionFromContext(c)
//
//	list := make([]models.Package, 0)
//	var count int64
//
//	err = s.GetPage(&req, &list, &count)
//	if err != nil {
//		e.Error(500, err, "查询失败")
//		return
//	}
//
//	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
//}

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
		e.Error(500, err, "查询失败")
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

	req.PackageId, err = strconv.Atoi(c.Param("id"))

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
	e.PageOK(list1, int(count2), req.GetPageIndex(), req.GetPageSize(), "查询成功")

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
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id} [post]
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
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id} [delete]
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

	patentId, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PatentId = patentId

	err = s.RemovePackagePatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.PackageBack, "删除成功")
}

//---------------------------------------------------patent--graph-------------------------------------------------------

// GetTheGraphByPackageId
// @Summary 获取专利关系图
// @Description  获取专利关系图
// @Tags 专利包
// @Router /api/v1/user-agent/package/{package_id}/relationship [get]
// @Security Bearer
func (e Package) GetTheGraphByPackageId(c *gin.Context) {
	spp := service.PatentPackage{}
	sup := service.UserPatent{}
	su := aservice.SysUser{}
	gservice := service.Node{}
	reqpp := dto.PackagePageGetReq{} //patent-package
	reqp := dto.PatentsIds{}         //patents
	requ := adto.SysUserById{}
	//fmt.Println("get the line 471")
	fmt.Println(c)
	var err error
	reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
	//fmt.Println("get the line 474")
	fmt.Println(reqpp.PackageId)
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&spp.Service).
		Errors
	//fmt.Println("get the line 480")
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reqpp.SetUpdateBy(user.GetUserId(c))
	//fmt.Println("get the line 486")
	listpp := make([]models.PatentPackage, 0)
	var count int64 //  not used
	err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
	fmt.Println(listpp)
	fmt.Println(reqp)
	for i := 0; i < len(listpp); i++ {
		fmt.Println(listpp[i].PatentId)
	}
	reqp.PatentIds = make([]int, len(listpp))
	for i := 0; i < len(listpp); i++ {
		reqp.PatentIds[i] = listpp[i].PatentId
	}
	//fmt.Println("get line 496")
	//1 := make([]models.Node, 0) //resultnode
	links := make([]models.Link, 0) //resultlink
	//listup, members, usertimes := e.AddGraphNodeByReq(c, &nodes, reqp, 5)
	//------------------------------------already get the patents id  now get the users id
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	//fmt.Println("get line 505")
	listup := make([]models.UserPatent, 0) //√
	//fmt.Println(reqp.PatentIds)
	sup.GetUsersByPatentId(&listup, &reqp) //√
	//for i := 0; i < len(listup); i++ {
	//	fmt.Println(listup[i])
	//}
	//fmt.Println("the list u success")
	//-----------------------------------already get the users id  now sort the users  and pick 8 users
	usertimes := make(map[int]int) // k-v is uid-times

	for i := 0; i < len(listup); i++ { //建立map usertimes
		if usertimes[listup[i].UserId] == 0 {
			usertimes[listup[i].UserId] = 1
		} else {
			usertimes[listup[i].UserId]++
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
		//fmt.Println(NodeList[i])
		//nownode.GraphName =    应该加上
	}

	//有75个人的uid，package里的patent的pid
	for i := 0; i < members; i++ {
		for j := i + 1; j < members; j++ {
			//pantentssum := len(listpp)
			//var ispatent [2][pantentssum]bool
			//RelationExist, _ := sup.GetTwoUserRelationshipInThisPackage(&reqp, usertimes1[i].Key, usertimes1[j].Key)
			RelationExist := 0
			listpatent1 := make([]int, 0)
			listpatent2 := make([]int, 0)
			for z := 0; z < len(listup); z++ {
				if listup[z].UserId == usertimes1[i].Key {
					listpatent1 = append(listpatent1, listup[z].PatentId)
				}
			}
			for z := 0; z < len(listup); z++ {
				if listup[z].UserId == usertimes1[j].Key {
					listpatent2 = append(listpatent2, listup[z].PatentId)
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
	//	reqp2 := dto.PatentsIds{} //find patent list of node[i]
	//	z := 0
	//	for j := 0; j < len(listup); j++ {
	//		nodeid := usertimes[i].Key //get node id
	//		if listup[j].UserId == nodeid {
	//			reqp2.PatentIds[z] = listup[j].PatentId
	//			z++
	//		}
	//	}
	//	_, _, _ = e.AddGraphNodeByReq(c, &nodes, reqp2, i*14+5)
	//}

}

// GetTheGraphByPackageId3
// @Summary 获取专利包中专利的发明人的关系
// @Description  获取专利包中专利的发明人的关系
// @Tags 专利表
// @Router /api/v1/user-agent/package/{packageId}/relationship3 [get]
// @Security Bearer
func (e Package) GetTheGraphByPackageId3(c *gin.Context) {
	//spp := service.PatentPackage{}
	sup := service.UserPatent{}
	su := aservice.SysUser{}
	sp := service.Patent{}
	gservice := service.Node{}
	//reqpp := dto.PackagePageGetReq{} //patent-package
	reqp := dto.PatentsIds{} //patents
	requ := adto.SysUserById{}
	spp := service.PatentPackage{}
	reqpp := dto.PackagePageGetReq{} //patent-package
	//fmt.Println("get the line 471")
	//fmt.Println(c)
	var err error
	//reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
	//fmt.Println("get the line 474")
	//fmt.Println(reqpp.PackageId)
	//s := service.UserPatent{}
	//requp := dto.UserPatentObject{}
	//reqp := dto.PatentsIds{}

	reqpp.PackageId, err = strconv.Atoi(c.Param("id")) //get packageId
	//fmt.Println("get the line 474")
	fmt.Println(reqpp.PackageId)
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&spp.Service).
		Errors
	//fmt.Println("get the line 480")
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reqpp.SetUpdateBy(user.GetUserId(c))
	//fmt.Println("get the line 486")
	listpp := make([]models.PatentPackage, 0)
	var count int64 //  not used
	err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
	//fmt.Println(listpp)
	//fmt.Println(reqp)
	for i := 0; i < len(listpp); i++ {
		fmt.Println(listpp[i].PatentId)
	}
	reqp.PatentIds = make([]int, len(listpp))
	for i := 0; i < len(listpp); i++ {
		reqp.PatentIds[i] = listpp[i].PatentId
	}
	listp := make([]models.Patent, 0)
	var count2 int64
	fmt.Println("begin to find patent:")
	fmt.Println(reqp)
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&reqp).
		MakeService(&sp.Service).
		Errors
	err = sp.GetPageByIds(&reqp, &listp, &count2)
	fmt.Println("找到了所有的专利包的patent")
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
	//listup2 := make([]models.UserPatent, 0)  //√          关注了专利(本用户所关注的)的所有用户数据
	listInventorId := make(map[string]int)      //key-value  为   发明人名字-发明人id（自己定的）
	listup2 := make([]models.InventorPatent, 0) //发明了了专利(本数据包内)的所有用户数据
	FindTheInventorFromPatents(&listInventorId, &listup2, listp)
	//按patentproperties格式为：{"patentId":232,"TI":"string","PNM":"232","AD":"string","PD":"string","CL":"string","PA":"string","AR":"string","PINN":"刘贺祥;刘禹辰;刘佳绮","CLS":"string","CreateBy":19,"UpdateBy":0}
	//fmt.Println(reqp.PatentIds)
	//已获得所有patent，获取字符串并且分出发明人-专利关系
	//生成
	//sup.GetFocusUsersByPatentId(&listup2, &reqp) //√
	//for i := 0; i < len(listup2); i++ {
	//	fmt.Println(listup2[i])
	//}
	//fmt.Println("the listup2 u success")

	usertimes := make(map[int]int) // k-v is uid-times   in package is inventorname-times

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
	firstlinks := first10                  //第一次处理的边数
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
	if len(NodeList) < 1 {
		e.GraphOK(NodeList, links, "查询成功")
	}
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
	NodelistIdToTimeList := make(map[string]int)            //key-value   点id-----在usertimes1,usersPatents中的排序
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
				NodelistIdToTimeList[nowNode.NodeId] = target
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
	//		for z := 0; z < len(userspatents[NodelistIdToTimeList[i]].Patentsid); z++ {
	//			for z1 := 0; z1 < ; z1++ {
	//
	//			}
	//		}
	//	}
	//}
	useruserrelation5 := make(map[int]int)

	for i := first10; i < len(NodeList); i++ { //后续regular点加边
		iToUserspatentsPosition := NodelistIdToTimeList[NodeList[i].NodeId]
		for j := i + 1; j < len(NodeList); j++ {
			RelationExist := 0
			jToUserspatentsPosition := NodelistIdToTimeList[NodeList[j].NodeId]
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
				fmt.Println(useruserrelation5[i*500+j])
			}

		}
	}
	fmt.Println(useruserrelation5)
	useruserrelation6 := rankByWordCount(useruserrelation5) //第三次边的关系
	thirdlinks := 50                                        //第三次边的数量
	for i := 0; i < minresult(thirdlinks, len(useruserrelation6)); i++ {
		source := useruserrelation6[i].Key / 500 //source在NodeList中的位置
		target := useruserrelation6[i].Key % 500

		if ExtendNodeTime[source] >= 2 { //regularnode扩展超过3个点
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
	listu := make([]amodels.SysUser, 0) //需要显示的user对象
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&requ, nil).
		MakeService(&su.Service).
		Errors
	fmt.Println("listInventorId:", listInventorId)
	for i := 0; i < len(NodeList); i++ { //找名字
		//int, err := strconv.Atoi(string)
		requ.Id, err = strconv.Atoi(NodeList[i].NodeId)
		//su.Get(&requ, &listu[i]) //查找需要显示的user对象
		var user1 amodels.SysUser
		user1.UserId = requ.Id
		for k, v := range listInventorId {
			if v == requ.Id {
				user1.Username = k
			}
		}
		listu = append(listu, user1)
	}

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&gservice.Service).
		Errors
	//err = gservice.GetNodes(&NodeList)
	max := 0
	min := 100000
	NodeList[0].NodeValue = usertimes1[0].Value //先把最大的点  操作一下
	NodeList[0].NodeSymbolizeSize = 50
	NodeList[0].NodeName = listu[0].Username //点的名字
	for i := 1; i < first10; i++ {           //strongNode 取value（有多少重复的patent（time））
		NodeList[i].NodeValue = usertimes1[i].Value
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	for i := first10; i < len(NodeList); i++ { //regularNode 取value（有多少重复的patent（time））
		NodeList[i].NodeValue = usertimes1[NodelistIdToTimeList[NodeList[i].NodeId]].Value
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
		NodeList[i].NodeSymbolizeSize = float32(float32(NodeList[i].NodeValue*30) / float32(maxresult((max), 1)))
		fmt.Println(listu[i].Username)
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

// 糊弄

func (e Package) GraphOK(list []models.Node, links []models.Link, s string) {

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

// FindTheInventorFromPatents --------------------------------------------------------------------------
// 查找patents中的发明人
func FindTheInventorFromPatents(listInventorId *map[string]int, listup2 *[]models.InventorPatent, listp []models.Patent) error {
	var err error
	//var s string
	fmt.Println("进入findtheinventor")
	count := 0 //第几个新人inventor 也是id号
	for z := 0; z < len(listp); z++ {
		words := make([]string, 0)
		fmt.Println(listp[z].PatentProperties)
		for i := 0; i < len(listp[z].PatentProperties); i++ {
			//fmt.Println(s[i])

			if listp[z].PatentProperties[i] == '"' && listp[z].PatentProperties[i+1] == 'P' && listp[z].PatentProperties[i+2] == 'I' {
				//var nowname string
				now := i + 8
				for j := i + 8; j < len(listp[z].PatentProperties); j++ {
					if listp[z].PatentProperties[j] == '"' {
						words = append(words, listp[z].PatentProperties[now:j])
						break
					}
					if listp[z].PatentProperties[j] == ';' {
						words = append(words, listp[z].PatentProperties[now:j])
						now = j + 1
					}
				}
				break
			}
		}
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
			_, ok := (*listInventorId)[words[i]]
			if !ok {
				(*listInventorId)[words[i]] = count
				count++
			}
			var inventorpatent models.InventorPatent
			inventorpatent.PatentId = listp[z].PatentId
			inventorpatent.UserId = (*listInventorId)[words[i]]
			*listup2 = append(*listup2, inventorpatent)
		}
	}

	return err
	//assertEquals("A", getKeyByLoop(map, 1));  通过value找key的值
}

// --------------------------------------------------------------------------------------------------------------------
// map排序
func rankByWordCount(wordFrequencies map[int]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	//从小到大排序
	//sort.Sort(pl)
	//从大到小排序
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func minresult(a1 int, a2 int) int {
	if a1 >= a2 {
		return a2
	} else {
		return a1
	}
}
func maxresult(a1 int, a2 int) int {
	if a1 >= a2 {
		return a1
	} else {
		return a2
	}
}
