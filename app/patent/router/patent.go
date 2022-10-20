package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatentRouter)
}

// 需认证的路由代码
func registerPatentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Patent{}

	r := v1.Group("/patent").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPatentLists)
		r.GET("/:patent_id", api.GetPatentById)
		//r.GET("/patent-name/:ti", api.GetPatentByName)
		r.POST("", api.InsertPatent)
		r.PUT("", api.UpdatePatent)
		r.DELETE("/:patent_id", api.DeletePatentByPatentId)

		r.POST("/claim", api.ClaimPatent)
		r.POST("/focus", api.FocusPatent)
	}

}
