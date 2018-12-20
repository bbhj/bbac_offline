package main

import (
	"github.com/bbhj/bbac/models"
	"github.com/bbhj/bbac/task"
	"github.com/bbhj/bbac/controllers"
	_ "github.com/bbhj/bbac/routers"

	"github.com/astaxie/beego"
	"github.com/esap/wechat"
)

func main() {

	beego.SetLogFuncCall(true)
	logfile := "logs/" + beego.AppConfig.String("appname") + ".log"
	beego.SetLogger("file", `{"filename":"` + logfile + `", "level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)


	models.Connect()
	defer models.Close()
	beego.Info(beego.BConfig.AppName, "start...")
	beego.Info("runmode:", beego.BConfig.RunMode)

	wechat.Debug = false
	wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))

	// openid := "oPPbr0M2h0geV-jgzUPve9g3x3jg"
	// msg := "xxx"
	// wechat.SendText(openid, 0, msg)

	go task.CronTask()
	if beego.BConfig.RunMode != "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
