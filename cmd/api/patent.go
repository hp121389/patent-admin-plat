package api

import "go-admin/app/user-agent/router"

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}
