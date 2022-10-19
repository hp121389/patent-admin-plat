package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/router"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	router.RouterCheckRole = append(router.RouterCheckRole, registerPatentRouter)
}

// 需认证的路由代码
func registerPatentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Patent{}

	r := v1.Group("/patent-list").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/get_patent_lists", api.GetPatentLists)
		r.GET("/get_by_patent_id/:patent_id", api.GetPatentById)
		//r.GET("/patent-name/:ti", api.GetPatentByName)
		r.POST("/post_a_patent/", api.InsertPatent)
		r.PUT("/change_a_patent/", api.UpdatePatent)
		r.DELETE("/delete_a_patent_by_id/:patent_id", api.DeletePatentByPatentId)
	}

}

//一般来说,Controller是Handler,但Handler不一定是Controller。
