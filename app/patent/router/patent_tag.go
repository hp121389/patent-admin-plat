package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPatentTagRouter)
}

// 需认证的路由代码
func registerPatentTagRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.PatentTag{}

	r1 := v1.Group("/patent-tag").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{

		r1.GET("/tags/:patent_id", api.GetTags)      //测试√
		r1.GET("/patents/:tag_id", api.GetPatent)    //测试√
		r1.POST("", api.InsertPatentTagRelationship) //测试√

		r1.PUT("", api.GetTags)

		r1.DELETE("/:id", api.GetTags)
	}

}
