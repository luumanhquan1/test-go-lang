package main

import (
	_ "beego-basic/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// @title Todo Application
// @description This is a todo list management application
// @version 1.0
// @host localhost:8081
func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// conf.SetUpConnectDb()
	beego.Run(":8081")

}
