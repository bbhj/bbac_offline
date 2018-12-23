package controllers

import (
	"fmt"
	"github.com/bbhj/bbac/models"
	"github.com/astaxie/beego"
	"strings"
)

// Operations about Admins
type AdminController struct {
	beego.Controller
}

// @Title CreateAdmin
// @Description create users
// @Param	body		body 	models.Admin	true		"body for user content"
// @Success 200 {int} models.Admin.Id
// @Failure 403 body is empty
// @router / [post]
func (u *AdminController) Post() {
	u.ServeJSON()
	return
}

// @Title GetAll
// @Description get all Admins
// @Success 200 {object} models.Admin
// @router / [get]
func (u *AdminController) GetAll() {
	// u.Data["json"] = map[string]string{"openid": openid}
	u.ServeJSON()
}

// @Title Get
// @Description get user by openid
// @Param	openid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Admin
// @Failure 403 :openid is empty
// @router /:openid [get]
func (u *AdminController) Get() {
	openid := strings.TrimLeft(u.GetString(":openid"), ":")
	username := u.GetString("username")
	fmt.Print("=====AdminController=", openid, "=====", username)
	status := 0
	aa := models.GetSummariesByOpenID(openid, status)
	u.Data["json"] = aa

	u.ServeJSON()
}

// // @Title Update
// // @Description update the user
// // @Param	openid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.Admin	true		"body for user content"
// // @Success 200 {object} models.Admin
// // @Failure 403 :openid is not int
// // @router /:openid [put]
// func (u *AdminController) Put() {
// 	openid := strings.TrimLeft(u.GetString(":openid"), ":")
// 	if openid != "" {
// 		// uu, err := models.UpdateAdmin(uid, &user)
// 		// if err != nil {
// 		// 	u.Data["json"] = err.Error()
// 		// } else {
// 		// 	u.Data["json"] = uu
// 		// }
// 	}
// 	u.ServeJSON()
// }

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /list [get]
func (u *AdminController) List() {
	openid := u.GetString("openid")
	status, _ := u.GetInt("status")
	u.Data["json"] = models.GetSummariesByOpenID(openid, status)
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *AdminController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
