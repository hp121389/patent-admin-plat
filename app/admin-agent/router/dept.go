package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin-agent/apis"
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

		r.GET("/list", api.GetDeptList)                          //  管理员列表部门信息    √
		r.GET("/relaList", api.GetDeptRelaList)                  //  管理员列表用户团队申请信息    √
		r.GET("/:dept_id", api.GetDeptById)                      //  管理员通过NAME查询部门信息
		r.GET("/relaListById/:dept_id", api.GetRelaListById)     //  管理员通过团队ID查询关系列表    √
		r.POST("", api.InsertDept)                               //  管理员创建部门团队    √
		r.PUT("/offline/:dept_id", api.OfflineDept)              //  管理员下线部门团队    √
		r.PUT("/reOnline/:dept_id", api.ReOnlineDept)            //  管理员重新上线部门团队    √
		r.PUT("/unJoin/:dept_id/:user_id", api.UnJoinDept)       //  管理员将用户踢出团队    √
		r.PUT("/recoverJoin/:dept_id/:user_id", api.RecoverJoin) //  管理员将用户恢复加入团队    √
		r.PUT("/join/:dept_id/:user_id", api.IfJoinDept)         //  管理员批准用户加入团队    √
		r.PUT("/joinReject/:dept_id/:user_id", api.JoinReject)   //  管理员拒绝用户加入团队    √
		r.PUT("/unReject/:dept_id/:user_id", api.UnReject)       //  撤销驳回     √

		//r.PUT("/leader/:dept_id", api.GetDeptById)       //  管理员批准用户成为组长
		//r.PUT("/leaderReject/:dept_id", api.GetDeptById) //  管理员拒绝用户成为组长
		//r.PUT("/exit/:dept_id", api.GetDeptById)       //  管理员批准用户退出团队
		//r.PUT("/exitReject/:dept_id", api.GetDeptById) //  管理员拒绝用户退出团队

	}
}
