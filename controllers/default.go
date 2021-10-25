package controllers

import (
	"leonet-gateway/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (ctx *MainController) Get() {
	user := new(models.User)
	var users []*models.User
	orm.NewOrm().QueryTable(&user).All(&users)
	ctx.Data["json"] = &users
	ctx.ServeJSON()
}
