package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSearchRouter)
	routerCheckRole = append(routerCheckRole, registerAuthedSearchRouter)
}

// 需认证的路由代码
func registerSearchRouter(v1 *gin.RouterGroup) {
	api := apis.Search{}
	r := v1.Group("/search")
	{
		r.POST("", api.Search)
		//r.GET("advance", api.Insert)
	}
}

// 需认证的路由代码
func registerAuthedSearchRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Search{}
	r := v1.Group("/auth-search").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("", api.AuthSearch)
		//r.GET("advance", api.Insert)
	}
}
