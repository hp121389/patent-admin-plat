package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/user-agent/models"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	amodels "go-admin/app/admin/models"
	aservice "go-admin/app/admin/service"
	adto "go-admin/app/admin/service/dto"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

//import

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
//	req := dto.PackageGetPageReq{}
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
// @Param data body dto.PackageInsertReq true "body"
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
