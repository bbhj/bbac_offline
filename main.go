package main

import (
	"github.com/bbhj/bbac/models"
	"github.com/bbhj/bbac/crontab"
	_ "github.com/bbhj/bbac/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/esap/wechat"
	"github.com/jinzhu/gorm"
)

// var Pool *models.Pool

var conn *gorm.DB

func main() {

	conn, _ := models.Connect()
	defer conn.Close()

	beego.SetLogger("file", `{"filename":"logs/bbhj.log", "level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.SetLogFuncCall(true)


	logs.Info("applicatoin start...")
	wechat.Debug = true
	// wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))
	wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))

	// openid := "oPPbr0M2h0geV-jgzUPve9g3x3jg"
	// msg := "xxx"
	// wechat.SendText(openid, 0, msg)

	go crontab.CronTask()
	if beego.BConfig.RunMode != "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
