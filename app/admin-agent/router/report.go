package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerReportRouter)
}

// 需认证的路由代码
func registerReportRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Report{}

	r := v1.Group("/report").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.PUT("/unReject/:report_id", api.UnReject)                  //管理员撤销驳回，处理中   √
		r.PUT("/reject/:report_id", api.Reject)                      //管理员驳回请求(标记)   √
		r.PUT("/upload", api.UploadReport)                           //管理员更新上传文件相关数据   √（真正的上传接口在public）
		r.PUT("/files/:report_id", api.DeleteUploadReport)           //管理员删除报告（删除files）   √
		r.GET("/:report_id", api.GetReportById)                      //管理员通过ID查看报告申请表   ！！！
		r.GET("/patent/:report_id", api.GetPatentByReId)             //管理员通过ID查看报告申请的专利   ！！！
		r.GET("/reportList/:patent_id", api.GetReportListByPatentId) //管理员通过 PatentId 去查看 申请 的报告有哪些   ！！！
		r.GET("/type/:type", api.GetListsByType)                     //管理员查看 Type类型 报告申请表   ！！！
	}

}
