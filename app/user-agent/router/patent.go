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
		r.GET("", api.GetPatentLists)                               //显示专利本地数据√
		r.GET("/:patent_id", api.GetPatentById)                     //查询专利√
		r.GET("/claim", api.GetClaimPages)                          //显示认领专利√
		r.GET("/claim/search", api.FindClaimPages)                  //搜索认领专利√
		r.GET("/focus", api.GetFocusPages)                          //显示关注专利√
		r.GET("/focus/search", api.FindFocusPages)                  //搜索关注专利√
		r.GET("/user", api.GetUserPatentsPages)                     //获取该用户所有专利列表
		r.GET("/focus/graph/relation", api.GetRelationGraphByFocus) //获取该用户关注的专利的发明人关系图谱(点位置随机、大小数量和线的数量、值根据数据生成)
		r.GET("/focus/graph/tech", api.GetTechGraphByFocus)         //获取该用户关注的专利的技术关系图谱(点位置随机、大小数量和线的数量、值根据数据生成)

		r.POST("", api.InsertIfAbsent)    //添加专利√
		r.POST("/claim", api.ClaimPatent) //认领专利√
		r.POST("/focus", api.FocusPatent) //关注专利√

		r.PUT("", api.UpdatePatent)                    //修改专利
		r.PUT("/claim/:PNM/desc", api.UpdateClaimDesc) //修改认领专利简介
		r.PUT("/focus/:PNM/desc", api.UpdateFocusDesc) //修改关注专利简介

		r.DELETE("/:patent_id", api.DeletePatent) //删除该专利√
		r.DELETE("/claim/:PNM", api.DeleteClaim)  //取消关注√
		r.DELETE("/focus/:PNM", api.DeleteFocus)  //取消认领√
	}

}
