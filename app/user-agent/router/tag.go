package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTagRouter)
}

// 需认证的路由代码
func registerTagRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Tag{}
	r := v1.Group("/tag").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.PUT("", api.Update)    //更新标签信息
		r.GET("/:id", api.Get)   //根据标签id获取标签
		r.POST("", api.Insert)   //增加标签并增加用户--标签信息
		r.DELETE("", api.Delete) //删除标签
	}

}
