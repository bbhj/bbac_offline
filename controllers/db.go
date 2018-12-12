package controllers

import (
	"github.com/bbhj/bbac/models"
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
	var msg models.RetMsg
	flag, dbmsg :=  models.Init() 
	msg.Data = dbmsg
	if flag {
		msg.Status = 0
		msg.Msg = "init db successfully!"
	} else {
		msg.Status = -1
		msg.Errmsg = "init db failed!"
	}
	u.Data["json"] = msg
	u.ServeJSON()
}
