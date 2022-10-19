package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/router"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	router.RouterCheckRole = append(router.RouterCheckRole, registerPatentTagRouter)
}

// 需认证的路由代码
func registerPatentTagRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.PatentTag{}

	r1 := v1.Group("/patent_tag").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.GET("", api.GetPatentListsByUserId)
		r1.GET("/get-patent-lists-by-userid/:user_id", api.GetPatentListsByUserId)
	}

}

//一般来说,Controller是Handler,但Handler不一定是Controller。
