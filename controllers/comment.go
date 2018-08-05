package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bbhj/baobeihuijia/models"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type CommentController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *CommentController) Post() {
	var postValue models.Comment
	// var postValue postRecord
	json.Unmarshal(u.Ctx.Input.RequestBody, &postValue)
	// fmt.Println(u.Ctx.Input.RequestBody)
	fmt.Println(postValue)
	models.AddComment(postValue)

	return
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *CommentController) GetAll() {
	aa := models.GetComment()
	u.Data["json"] = aa
	u.ServeJSON()
}

// @Title Get
// @Description get user by uuid
// @Param	uuid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uuid is empty
// @router /:uuid [get]
func (u *CommentController) Get() {
	uuid := strings.TrimLeft(u.GetString(":uuid"), ":")
	if uuid != "" {
		// user, err := models.GetUser(uid)
		// if err != nil {
		// 	u.Data["json"] = err.Error()
		// 	u.Data["json"] = user
		// }
	}
	// aa := models.GetCommentsByUUID(uuid)
	aa := models.GetCommentsByUUID(uuid)
	u.Data["json"] = aa
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *CommentController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		// var user models.User
		// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		// uu, err := models.UpdateUser(uid, &user)
		// if err != nil {
		// 	u.Data["json"] = err.Error()
		// } else {
		// 	u.Data["json"] = uu
		// }
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *CommentController) Delete() {
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
func (u *CommentController) Login() {
	// username := u.GetString("username")
	// password := u.GetString("password")
	// if models.Login(username, password) {
	// 	u.Data["json"] = "login success"
	// } else {
	// 	u.Data["json"] = "user not exist"
	// }
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *CommentController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
