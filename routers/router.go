package routers

import (
	"beego-basic/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func MyJWTMiddleware(ctx *context.Context) {
	// ctx.Output.Status = http.StatusUnauthorized
	// ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	// ctx.WriteString("Invalid token")
	ctx.Input.SetData("userId", "1233")

}

func init() {

	beego.Router("/auth", &controllers.MainController{})
	beego.InsertFilter("/index/*", beego.BeforeRouter, MyJWTMiddleware)
	beego.Router("/index", &controllers.MainController{}, "get:CreateAt")
	beego.Router("/index/1", &controllers.MainController{}, "get:CreateAt")

}
