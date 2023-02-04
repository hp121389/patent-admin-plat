package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"

	"go-admin/common/middleware"
)

func init() {
	RouterCheckRole = append(RouterCheckRole, registerDeptRouter)
}

// 需认证的路由代码
func registerDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Dept{}
	r := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{

		r.GET("", api.GetDeptList) //  管理员列表部门信息    √
		r.PUT("/:id", api.UpdateDept)
		r.POST("", api.CreateDept)
		r.DELETE("/:id", api.RemoveDept)
	}
}
