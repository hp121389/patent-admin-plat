package main

import (
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title PatentAdminPlat API
// @version 0.1.0
// @description 专利检测平台v0.1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
