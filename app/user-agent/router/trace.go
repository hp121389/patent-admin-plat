package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTraceRouter)
}

// 需认证的路由代码
func registerTraceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	trace := apis.Trace{}
	r := v1.Group("/tracing").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/logs", trace.SelectTraceLog)
	}
}
