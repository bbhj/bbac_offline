package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bbhj/baobeihuijia/models"
	"github.com/imroc/req"
	"github.com/esap/wechat"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	upload_path string = "./upload/"
	APPID       string = "wxc4e11081e3d5bdf7"
	SECRET      string = "e3284a3123d4ed4e06ee09bf0171bef7"
	AESKEY      string = "RHtxqJ3CQyspdfT9GZSTReNfoOC58ocEzj47HYAU4Us"
	TOKEN       string = "deanhome"
)

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
	beego.Info("111111======")

	var wxlogin models.WechatLogin
	json.Unmarshal(u.Ctx.Input.RequestBody, &wxlogin)

	// beego.AppConfig.String("wehcat_getAnalysisDailySummary") + wechat.GetAccessToken()
	apiurl := fmt.Sprintf("%sappid=%s&secret=%s&&js_code=%s&grant_type=authorization_code", beego.AppConfig.String("WechatAuthUrl"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), wxlogin.Code)

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	a.ToJSON(&wxlogin.User)

	models.AddUserInfo(wxlogin.User)

	var login models.Login
	login.IP = u.Ctx.Input.IP()
	login.Openid =  wxlogin.User.Openid
	login.Unionid = wxlogin.User.Unionid
	login.Platform = wxlogin.MobileInfo.Platform
	login.NetworkType = wxlogin.MobileInfo.NetworkType
	login.Brand = wxlogin.MobileInfo.PhoneBrand
	login.Pmodel = wxlogin.MobileInfo.Model
	login.Longitude = wxlogin.MobileInfo.Longitude
	login.Latitude = wxlogin.MobileInfo.Latitude
	fmt.Println("login record: ", login)

	models.AddLogin(login)

	var st models.ServiceTime
	st.Openid = wxlogin.User.Openid
	st.Type = "login"
	st.Duration = 3
	models.AddServiceTime(st)

	// u.Ctx.Output.Header("Authorization", tokenString)
	// u.Ctx.SetCookie("Authorization", tokenString)
	u.Ctx.SetCookie("openid", wxlogin.User.Openid)

	var auth models.Auth
	auth.Openid = wxlogin.User.Openid
	auth.Unionid = wxlogin.User.Unionid
	auth.YourCity = u.Ctx.Input.Header("city")
	auth.IsAdmin = false
	auth.IsVolunteer = true
	auth.ServiceTime = "100.5"
	auth.Rights = 3
	fmt.Println("====auth: ", wechat.GetAccessToken())
	auth.Token = "dean"
	auth.AccessToken =  wechat.GetAccessToken()
	auth.WeCosUrl = fmt.Sprintf("https://%s.file.myqcloud.com/files/v2/%s/%s/%s", beego.AppConfig.String("QcloudCosRegion"), beego.AppConfig.String("QcloudCosAPPID"), beego.AppConfig.String("QcloudCosBucket"), beego.AppConfig.String("QcloudCosUploadDir"))
	fmt.Println("====auth: ", auth)

	retData, _ := json.Marshal(auth)
	u.Data["json"] = string(retData)

	// time.Sleep(time.Second * 1)

	u.ServeJSON()
}

// @Title FormID
// @Description Colloect Wechet User Form Id, prepare for sending template message.
// @Param	openid 	string	true		"wechat user openid"
// @Param	formid 	string	true		"wechat formid"
// @Success 200 {string} login success
// @Failure 403 forbidden
// @router /formid [post]
func (u *UserController) FormID() {
	fmt.Println("===form====:")
	var tform models.TemplateFormID
	json.Unmarshal(u.Ctx.Input.RequestBody, &tform)
	fmt.Println("===form====:", tform)
	logs.Info("wechat formid: ", tform)
	models.AddTemplateFormID(tform)

	u.Data["json"] = "status: 0" 
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /get_info [get]
func (u *UserController) GetInfo() {
	var userinfo  models.UserInfo
	userinfo.Avator = "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJuPPo6bJsib4LVLTfDCaI5Q3yULbiciaDDnxoq5uKejypd8xMM3qFbSdkogLSSpgXw7ZFSPBhpzUcPA/132"
	userinfo.Name = "xxxxx"
	userinfo.UserId = "1"
	userinfo.Access = `["super_admin", "admin"]`
	u.Data["json"] = userinfo
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /get_drag_list [get]
func (u *UserController) GetDragList() {
        msg := "xxx"
        u.Ctx.WriteString(msg)
	return
}


// @Title for webapi user login
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /login [get]
func (u *UserController) WxLogin() {
	var wxlogin models.OpenWeixinAccessToken

	// uri?code=021cljRR1DUhk81klvSR1xCxRR1cljRM&state=bbhj
	var wxauth models.WechatAuth
	wxauth.Code = u.GetString("code")
	wxauth.State = u.GetString("state")
	wxauth.Scene = u.GetString("scene")
	wxauth.Path = u.GetString("path")
	wxauth.ShareTicket = u.GetString("shareTicket")

	// json.Unmarshal(u.Ctx.Input.RequestBody, &wxauth)
	logs.Info("-------------", wxauth)
	if "" == wxauth.Code {
		return
	}

	if "" != wxauth.Scene {
		// apiurl := fmt.Sprintf("%sappid=%s&secret=%s&js_code=%s&grant_type=authorization_code", beego.AppConfig.String("WechatAuthUrl"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), wxauth.Code)
		apiurl := beego.AppConfig.String("wechat_mina") + wxauth.Code

		req.SetTimeout(50 * time.Second)
		a, _ := req.Get(apiurl)

		logs.Debug(a)
		a.ToJSON(&wxlogin)
		logs.Debug("xxxxx", wxlogin)

		return
	} else {
		apiurl := beego.AppConfig.String("weixin_oauth") + wxauth.Code

		r, _ := req.Get(apiurl)
		// logs.Info(r)
		r.ToJSON(&wxlogin)
		if wxlogin.Openid == "" {
			return
		}
		models.AddOpenWeixinAccessToken(wxlogin)
		logs.Info("weixin login: ", wxlogin)

		// get userinfo
		apiurl = "https://api.weixin.qq.com/sns/userinfo?access_token=" + wxlogin.AccessToken + "&openid=" + wxlogin.Openid
		r, _ = req.Get(apiurl)
		var webuserinfo models.User
		r.ToJSON(&webuserinfo)
		models.AddUserInfo(webuserinfo)
		logs.Info("get weixin userinfo: ", webuserinfo)

		// u.Ctx.SetCookie("access_token", wxlogin.AccessToken, 7200)
		u.Ctx.SetCookie("token", wxlogin.AccessToken, 7200)
	}
	// u.Ctx.Redirect(302, "https://wechat.baobeihuijia.com/index.html")
	u.Data["json"] = wxlogin
	u.ServeJSON()
}

