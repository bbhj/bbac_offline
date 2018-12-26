package controllers

import (
	_ "encoding/json"
	"github.com/bbhj/bbac/models"
	_ "strings"

	"github.com/astaxie/beego"
	_ "strconv"
	_ "time"
)

type ErrorController struct {
	beego.Controller
}

func (u *ErrorController) Error404() {
	var ret models.RetMsg
	ret.Status = -1
	ret.Errmsg = "404 page not found."
	u.Data["json"] = ret
	u.ServeJSON()
}

func (u *ErrorController) Error501() {
	var ret models.RetMsg
	ret.Status = -1
	ret.Errmsg = "404 page not found."
	u.Data["json"] = ret
	u.ServeJSON()
}

func (u *ErrorController) ErrorDB() {
	var ret models.RetMsg
	ret.Status = -1
	ret.Errmsg = "404 page not found."
	u.Data["json"] = ret
	u.ServeJSON()
}
