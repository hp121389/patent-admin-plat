package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/user-agent/apis"
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
		r.GET("", api.GetPatentLists)           //显示专利本地数据
		r.GET("/:patent_id", api.GetPatentById) //查询专利
		r.GET("/claim", api.GetClaimPages)      //显示认领专利
		r.GET("/focus", api.GetFocusPages)      //显示关注专利
		//显示专利包内专利
		//显示标签为xxx的专利
		//显示专利的标签
		//显示专利所属的专利包

		r.POST("", api.InsertPatent)      //添加专利
		r.POST("/claim", api.ClaimPatent) //认领专利
		r.POST("/focus", api.FocusPatent) //关注专利
		//为专利添加标签
		//将专利加入专利包

		r.PUT("", api.UpdatePatent) //修改专利

		r.DELETE("/:patent_id", api.DeletePatentByPatentId)
		r.DELETE("/claim/:patent_id", api.DeleteClaim) //取消关注
		r.DELETE("/focus/:patent_id", api.DeleteFocus) //取消认领
		//取消添加该标签
		//取消加入该专利包

	}

}
