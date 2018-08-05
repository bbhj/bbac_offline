package controllers

import (
	_ "encoding/json"
	"github.com/bbhj/baobeihuijia/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type SummaryController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *SummaryController) Post() {
	return
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *SummaryController) GetAll() {
	aa := models.GetSummaries()
	u.Data["json"] = aa
	// u.Data["json"] = map[string]string{"openid": openid}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *SummaryController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *SummaryController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *SummaryController) Delete() {
	// uid := u.GetString(":uid")
	// models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title List
// @Description Logs user into the system
// @Param       openid    query   string  true            "wechat openid"
// @Param       status    query   int     true            "summary stauts"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /list [get]
func (u *SummaryController) List() {
	openid := u.GetString("openid")
	status, _ := u.GetInt("status")
	u.Data["json"] = models.GetSummariesByOpenID(openid, status)
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *SummaryController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
