package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/user-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDeptRouter)
}

// 需认证的路由代码
func registerDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Dept{}

	r := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/list", api.GetDeptList)                           //  用户查看所有团队    √
		r.GET("/relaList", api.GetRelaListByUserId)               //  用户列表自己加入/管理的部门信息    √
		r.GET("/relaListOfUserId/:dept_id", api.RelaListOfUserId) //  用户在该部门的状态    √
		r.POST("/joinApply/:dept_id", api.UserJoinDept)           //  用户加入团队申请    √
		r.PUT("/joinCancel/:dept_id", api.UserCancelJoinDept)     //  用户撤销加入团队申请    √
		r.PUT("/reJoin/:dept_id", api.ReJoinDept)                 //  用户重新申请    √
		r.PUT("/exit/:dept_id", api.ExitDept)                     //  用户申请退出团队    √

		//r.PUT("/leaderApply/:dept_id", api.GetDeptById)  //  用户成为组长申请
		//r.PUT("/leaderCancel/:dept_id", api.GetDeptById) //  用户撤销成为组长申请
		//r.PUT("/leaderUndo/:dept_id", api.GetDeptById)   //  用户撤销组长申请

		//r.GET("/:dept_id", api.GetDeptById)            //  通过NAME查询部门信息
	}
}
