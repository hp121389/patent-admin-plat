package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTagRouter)
}

// 需认证的路由代码
func registerTagRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Tag{}
	r := v1.Group("/tag").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.PUT("", api.Update)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.DELETE("", api.Delete)
	}
	//r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	//{
	//	//r1.PUT("/role-status", api.Update2Status)
	//	//r1.PUT("/roledatascope", api.Update2DataScope)
	//}
}
