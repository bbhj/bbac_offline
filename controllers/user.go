package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bbhj/baobeihuijia/models"
	"github.com/airdb/sailor/req"
	"github.com/esap/wechat"
	"time"

	"github.com/astaxie/beego"
	"github.com/juusechec/jwt-beego"
)

const (
	upload_path string = "./upload/"
	APPID       string = "wxc4e11081e3d5bdf7"
	SECRET      string = "e3284a3123d4ed4e06ee09bf0171bef7"
	AESKEY      string = "RHtxqJ3CQyspdfT9GZSTReNfoOC58ocEzj47HYAU4Us"
	TOKEN       string = "deanhome"
)

type postRecord struct {
	IP string `json:"ip"`
	// OpenID       string  `json:"open_id"`
	// UionID       string  `json:"uion_id"`
	WechatCode string `json:"code"`
	WechatUserID
	WechatUserInfo  string `json:"userInfo"`
	WechatLoginInfo string `json:"loginInfo"`
}

type WechatLoginInfo struct {
	PhoneNetwork string  `json:"phone_network"`
	PhoneBrand   string  `json:"phone_brand"`
	PhoneModel   string  `json:"phone_model"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

type WechatUserID struct {
	OpenID  string `json:"openid"`
	UnionID string `json:"unionid"`
}

type WechatUserSession struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	SesssionKey string `json:"session_key"`
	WechatUserID
}

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
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
func (u *UserController) Delete() {
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
// @router /login [post]
func (u *UserController) Login() {

	var postValue postRecord
	json.Unmarshal(u.Ctx.Input.RequestBody, &postValue)
	// var userInfo UserInfo
	// json.Unmarshal([]byte(postValue.WechatUserInfo), &userInfo)
	wechat.Debug = true
	wechat.Set(TOKEN, APPID, SECRET)
	apiurl := fmt.Sprintf("%sappid=%s&secret=%s&&js_code=%s&grant_type=authorization_code", beego.AppConfig.String("WechatAuthUrl"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), postValue.WechatCode)

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	var usersession WechatUserSession
	a.ToJSON(&usersession)

	var user models.User
	json.Unmarshal([]byte(postValue.WechatUserInfo), &user)
	user.OpenID = usersession.OpenID
	user.UnionID = usersession.UnionID
	fmt.Println("wechat user info: ", user)
	models.AddUserInfo(user)

	var login models.Login
	fmt.Println("=====", postValue.WechatLoginInfo)
	json.Unmarshal([]byte(postValue.WechatLoginInfo), &login)
	login.IP = u.Ctx.Input.IP()
	login.OpenID = usersession.OpenID
	login.UnionID = usersession.UnionID
	fmt.Println("wechat login info: ", login)
	models.AddLogin(login)

	tokenString := ""
	if "" != usersession.OpenID {
		et := jwtbeego.EasyToken{
			Username: usersession.OpenID,
			Expires:  time.Now().Unix() + 3, //Segundos
			// Expires:  time.Now().Unix() + 3600, //Segundos
		}
		tokenString, _ = et.GetToken()
	}

	// ZZu.Header("Authorization", tokenString)
	// u.Ctx.Output.Header("Authorization", tokenString)
	u.Ctx.SetCookie("Authorization", tokenString)
	u.Ctx.SetCookie("openid", usersession.OpenID)
	data := fmt.Sprintf("{\"openid\": \"%s\", \"unionid\": \"%s\", \"token\": \"%s\"}", usersession.OpenID, usersession.UnionID, tokenString)
	u.Data["json"] = data

	// time.Sleep(time.Second * 1)
	et := jwtbeego.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		u.Ctx.Output.SetStatus(401)
		u.Data["json"] = "Permission Deny, jwt token error"
		u.ServeJSON()
	} else {
		// fmt.Println("bbbbbb vaild")
	}

	u.ServeJSON()
}
