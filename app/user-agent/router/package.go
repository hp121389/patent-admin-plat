package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPackageRouter)
}

// 需认证的路由代码
func registerPackageRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Package{}
	api2 := apis.UserPackage{}
	r := v1.Group("/package").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)                                   //显示本地专利包列表
		r.GET("/:id", api.Get)                                   //查询专利包
		r.GET("/user-package/:user_id", api2.GetPackageByUserId) //查询用户的专利包
		//r.GET("/user-package", api2.GetPage)
		r.POST("", api.Insert)    //用户添加专利包
		r.PUT("/", api.Update)    //修改专利包
		r.DELETE("/", api.Delete) //删除专利包
		//r.DELETE("/user-package/:package_id", api2.DeleteUserPackage)
	}
}
