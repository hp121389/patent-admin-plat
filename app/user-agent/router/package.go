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
	r := v1.Group("/package").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.ListByCurrentUser)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("/:id", api.Delete)

		r.GET("/:id/patent", api.GetPackagePatents)              //显示专利包内专利√
		r.POST("/:id/patent", api.InsertPackagePatent)           //将专利加入专利包√
		r.DELETE("/:id/patent/:pid", api.DeletePackagePatent)    //取消加入该专利包
		r.GET("/:id/patent/:PNM/isExist", api.IsPatentInPackage) //取消加入该专利包
	}
}
