package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTagRouter)
}

// 需认证的路由代码
func registerTagRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	//----------------------------------User-------------------------------
	//用户对report的权限有：get（查看报告，查看是否被驳回状态）；post（申请某个专利的报告(+report-patent关系)）；
	//----------------------------------Admin-------------------------------
	api := apis.Report{}

	r := v1.Group("/report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/unReject/:report_id", api.UnReject)     //管理员撤销驳回，处理中   √
		r.POST("/reject/:report_id", api.Reject)         //管理员驳回请求(标记)   √
		r.GET("/:report_id", api.GetReportById)          //管理员通过ID查看报告申请表   √
		r.GET("/patent/:report_id", api.GetPatentByReId) //管理员通过ID查看报告申请的专利   √
	}

	r1 := v1.Group("/valuation-report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.GET("", api.GetValuationLists) //管理员查看所有估值报告申请表   √
		r1.POST("/upload/:report_id", api.UploadValuationReport)
		//  管理员上传侵权报告Upload;同时把关系存入数据库;
		// （新建report，且新建关系，且上传文件）
	}

	r2 := v1.Group("/infringement-report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r2.GET("", api.GetInfringementLists) //管理员查看所有估值报告申请表   √
		r2.POST("/upload/:report_id", api.UploadInfringementReport)
		//  管理员上传侵权报告Upload;同时把关系存入数据库;
		// （新建report，且新建关系，且上传文件）
	}

}
