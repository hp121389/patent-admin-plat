package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/router"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	router.RouterCheckRole = append(router.RouterCheckRole, registerUserPatentRouter)
}

// 需认证的路由代码
func registerUserPatentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Patent{}

	r1 := v1.Group("/user-patent").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.GET("", api.GetPatentLists)
		r1.GET("/get-patent-lists-by-userid/:user_id", api.GetPatentById)
	}

}

//一般来说,Controller是Handler,但Handler不一定是Controller。
