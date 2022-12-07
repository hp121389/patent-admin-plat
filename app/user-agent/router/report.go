package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	apiUser "go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerReportRouter)
}

// 需认证的路由代码
func registerReportRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	apiUser := apiUser.Report{}
	r := v1.Group("/report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/reportList/:patent_id", apiUser.GetReportListByPatentId) // 用户查看该专利的报告有哪些   √
		r.GET("/query/:report_id", apiUser.UserGetReportById)            // 用户查询报告   √
		r.GET("", apiUser.UserGetReports)                                // 用户查看认领专利的报告申请表   √
		r.GET("/patent/:report_id", apiUser.GetPatentByReId)             // 用户通过 reportID 查看报告申请的专利   √
		r.GET("/:type", apiUser.UserGetReportByType)                     // 用户通过 类型 查看报告   √
		r.PUT("/cancel/:report_id", apiUser.CancelReport)                // 用户撤销申请报告   √
		r.PUT("/reApp/:report_id", apiUser.ReapplyReport)                // 用户重新申请报告   √
		r.POST("", apiUser.InsertReport)                                 // 用户申请报告   √
	}
}
