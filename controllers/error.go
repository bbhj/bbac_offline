package controllers

import (
	_ "encoding/json"
	_ "github.com/bbhj/bbac/models"
	_ "strings"

	"github.com/astaxie/beego"
	_ "strconv"
	_ "time"
)

type ErrorController struct {
	beego.Controller
}

func (u *ErrorController) Error404() {
	u.Data["json"] = map[string]string{"msg": "404 page not found"}
	u.ServeJSON()
}

func (u *ErrorController) Error501() {
	u.Data["json"] = map[string]string{"msg": "404 page not found"}
	u.ServeJSON()
}

func (u *ErrorController) ErrorDB() {
	u.Data["json"] = map[string]string{"msg": "404 page not found"}
	u.ServeJSON()
}
