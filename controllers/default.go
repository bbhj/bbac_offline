package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param       body            body    models.User     true            "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (u *MainController) Get() {
	beego.Info("this is default controller")
	msg := "hello"
	u.Ctx.WriteString(msg)
	return
}
