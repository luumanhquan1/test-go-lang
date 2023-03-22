package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// @Title TestEndpoint
// @Description Tests the API
// @Success 200 {object} model.User
// @Failure 403 body is empty
// @router /hello [get]
func (c *MainController) CreateAt() {
	fmt.Println(c.Data["userId"])
	// var users []models.User
	// conf.Db.Preload("SoThich").Preload("Img").Where(&models.User{
	// 	Id: "4c558f57-5744-4381-9016-d1cf6698141d",
	// }).Find(&users)

	c.Data["json"] = "wqewq"
	c.ServeJSON()

}
