package controllers

import (
	"encoding/json"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/bbhj/baobeihuijia/models"
	"github.com/astaxie/beego"
	_ "io"
	"math/rand"
	"time"
)

// Operations about Users
type WecosController struct {
	beego.Controller
}

// @Title auth
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /auth [get]
func (u *WecosController) Auth() {
	curtime := time.Now().Unix()
	randnum := rand.Intn(1000000000)

	multi_effect_signature := fmt.Sprintf("a=%s&b=%s&k=%s&e=%d&t=%d&r=%d&f=",
		beego.AppConfig.String("qcloud_appid"),
		beego.AppConfig.String("qcloud_bucket"),
		beego.AppConfig.String("qcloud_secretid"),
		curtime+60, curtime, randnum)
	beego.Info(multi_effect_signature)

	mac := hmac.New(sha1.New, []byte(beego.AppConfig.String("qcloud_secretkey")))
	mac.Write([]byte(multi_effect_signature))
	macSum := mac.Sum(nil)

	// data64 := base64.StdEncoding.EncodeToString([]byte(string(macSum) + multi_effect_signature))
	u.Data["json"] = base64.StdEncoding.EncodeToString([]byte(string(macSum) + multi_effect_signature))
	u.ServeJSON()

}

// @Title upload
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /upload [post]
func (u *WecosController) Upload() {
	var image models.Image
	json.Unmarshal(u.Ctx.Input.RequestBody, &image)
	models.UploadImage(image)
	u.Data["json"] = "upload success"
	u.ServeJSON()
}
