package main

import (
	"github.com/bbhj/baobeihuijia/models"
	_ "github.com/bbhj/baobeihuijia/routers"

	"github.com/astaxie/beego"
	"github.com/esap/wechat"
	"github.com/jinzhu/gorm"
)

// var Pool *models.Pool

var conn *gorm.DB

func main() {

	conn, _ := models.Connect()
	defer conn.Close()

	wechat.Debug = true
	// wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))
	wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))
	openid := "oPPbr0M2h0geV-jgzUPve9g3x3j"
	msg := "xxx"
	wechat.SendText(openid, 0, msg)
	if beego.BConfig.RunMode != "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
