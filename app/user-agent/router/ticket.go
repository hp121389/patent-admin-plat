package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTicketRouter)
}

// 需认证的路由代码
func registerTicketRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Ticket{}
	r := v1.Group("/tickets").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetTicketPages)
		r.POST("", api.CreateTicket)
		r.PUT("/:id", api.UpdateTicket)
		r.PUT("/:id/close", api.CloseTicket)
	}
}
