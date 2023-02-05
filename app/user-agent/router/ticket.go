package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	aApis "go-admin/app/admin-agent/apis"
	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTicketRouter)
}

// 需认证的路由代码
func registerTicketRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Ticket{}
	adminApi := aApis.Ticket{}
	r := v1.Group("/tickets").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetTicketPages)
		r.POST("", adminApi.CreateTicket)
		r.PUT("/:id", adminApi.UpdateTicket)
		r.PUT("/:id/close", adminApi.CloseTicket)
	}
}
