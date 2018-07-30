package controllers

import (
	"encoding/json"
	"github.com/airdb/baobeihuijia/models"

	"fmt"
	"github.com/astaxie/beego"
)

// Operations about Users
type ContactController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ContactController) Post() {
	var contact models.Contact
	json.Unmarshal(u.Ctx.Input.RequestBody, &contact)
	fmt.Println(contact)
	models.AddContact(contact)
	u.Data["json"] = map[string]string{"status": "success"}
	u.ServeJSON()
	return
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *ContactController) GetAll() {
	aa := models.GetContact()
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
func (u *ContactController) Get() {
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
func (u *ContactController) Put() {
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
func (u *ContactController) Delete() {
	// uid := u.GetString(":uid")
	// models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *ContactController) Login() {
	// username := u.GetString("username")
	// password := u.GetString("password")
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *ContactController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
