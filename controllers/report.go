package controllers

import (
	_ "encoding/json"
	_ "fmt"
	_ "github.com/bbhj/bbac/models"
	_ "github.com/imroc/req"
	_ "github.com/esap/wechat"
	_ "time"

	"github.com/astaxie/beego"
)

// Operations about Users
type ReportController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /save_error_logger [post]
func (u *ReportController) SaveErrorLogger() {
	rbody := u.Ctx.Input.RequestBody
	beego.Info("error log:", string(rbody))
	// {"type":"script","code":0,"mes":"Cannot read property 'then' of undefined","url":"https://wechat.baobeihuijia.com/#/home"}
	// u.Data["json"] = "status: 0" 
	msg := "xxx"
	u.Ctx.WriteString(msg)
	u.ServeJSON()
}
