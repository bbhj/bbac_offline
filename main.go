package main

import (
	"github.com/bbhj/baobeihuijia/models"
	_ "github.com/bbhj/baobeihuijia/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/esap/wechat"
	"github.com/jinzhu/gorm"
	"github.com/imroc/req"
)

// var Pool *models.Pool

var conn *gorm.DB

func main() {

	conn, _ := models.Connect()
	defer conn.Close()

	wechat.Debug = true
	// wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))
	wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))

        var tm models.TemplateMessage
        tm.Touser  = "oPPbr0M2h0geV-jgzUPve9g3x3jg"
        tm.TemplateID = "nYL5mnIUq5aybu8uZWaF1qwWPp_6Up4EhhUmrbHXWmc"
        tm.Page = "/pages/home/main"
        tm.FormID = "2fc6e039e32e2c77035d7c7311920ab1"
        tm.EmphasisKeyword = ""
        tm.Data.Keyword1.Value = "帮忙寻找xxx"
        tm.Data.Keyword2.Value = "阿正_Dean"
        tm.Data.Keyword3.Value = "常州市新北区河海东路200号8栋2单元301"
        tm.Data.Keyword4.Value = "这是备注"
        tm.Data.Keyword5.Value = "已有500人参与提供帮助"

        a2, _ := req.Post(beego.AppConfig.String("wechatApiSendTemplateMessage") + wechat.GetAccessToken(), req.BodyJSON(tm))
	logs.Debug( "send template message: ", a2 )

	// openid := "oPPbr0M2h0geV-jgzUPve9g3x3jg"
	// msg := "xxx"
	// wechat.SendText(openid, 0, msg)
	if beego.BConfig.RunMode != "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
