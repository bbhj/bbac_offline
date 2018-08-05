package controllers

import (
	"github.com/bbhj/baobeihuijia/models"
	"github.com/astaxie/beego"
)

// Operations about Articles
type DBController struct {
	beego.Controller
}

// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router / [get]
func (u *DBController) Get() {
	models.Init()
	u.Data["json"] = map[string]string{"msg": "init db successfully!"}
	u.ServeJSON()

}
