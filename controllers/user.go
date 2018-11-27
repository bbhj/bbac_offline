package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bbhj/baobeihuijia/models"
	"github.com/imroc/req"
	"github.com/esap/wechat"
	"time"

	"github.com/astaxie/beego"
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

	var wxlogin models.WechatLogin
	json.Unmarshal(u.Ctx.Input.RequestBody, &wxlogin)

	// beego.AppConfig.String("wehcat_getAnalysisDailySummary") + wechat.GetAccessToken()
	param := req.Param{
		"begin_date": "20181023",
		"end_date":  "20181023",
	}
	a1, _ := req.Post(beego.AppConfig.String("wehcat_getAnalysisDailySummary") + wechat.GetAccessToken(), req.BodyJSON(param))
	var analysisDaily models.AnalysisDailySummary
	a1.ToJSON(&analysisDaily)
	fmt.Printf("累计用户数: %d, 转发人数: %d, 转发次数: %d\n", analysisDaily.List[0].VisitTotal, analysisDaily.List[0].ShareUv, analysisDaily.List[0].SharePv)

	apiurl := fmt.Sprintf("%sappid=%s&secret=%s&&js_code=%s&grant_type=authorization_code", beego.AppConfig.String("WechatAuthUrl"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), wxlogin.Code)

	req.SetTimeout(50 * time.Second)
	a, _ := req.Get(apiurl)

	a.ToJSON(&wxlogin.User)

	models.AddUserInfo(wxlogin.User)

	var login models.Login
	login.IP = u.Ctx.Input.IP()
	login.OpenID =  wxlogin.User.Openid
	login.UnionID = wxlogin.User.Unionid
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
	models.AddTemplateFormID(tform)

	u.Data["json"] = "status: 0" 
	u.ServeJSON()
}
