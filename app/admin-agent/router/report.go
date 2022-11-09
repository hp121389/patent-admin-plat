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
	api := apis.Report{} //等着改，admin的
	r := v1.Group("/infringement-report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/upload", api.GetInfringementReportById)    //管理员上传侵权报告Upload
		r.POST("/", api.GetInfringementReportById)          //管理员驳回请求(CLICK驳回，传一个值过来，接到之后)Reject
		r.GET("/:report_id", api.GetInfringementReportById) //管理员通过ID查看侵权报告申请表
		r.GET("", api.GetInfringementLists)                 //管理员查看所有侵权报告申请表

	}

	r1 := v1.Group("/valuation-report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.POST("/upload", api.GetValuationLists)         //管理员上传侵权报告Upload
		r1.POST("/", api.GetValuationLists)               //管理员驳回请求(CLICK驳回，传一个值过来，接到之后)Reject
		r1.GET("/:report_id", api.GetValuationReportById) //管理员通过ID查看估值报告申请表
		r1.GET("", api.GetValuationLists)                 //管理员查看所有估值报告申请表
	}

}
