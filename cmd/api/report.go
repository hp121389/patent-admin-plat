package api

import "go-admin/app/admin-agent/router"

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}
