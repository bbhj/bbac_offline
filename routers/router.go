// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/bbhj/bbac/controllers"
	"github.com/bbhj/bbac/models"

	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

func init() {
	// beego.InsertFilter(`((?!/lastest/wechatapi/wechat/login$).*)`, beego.BeforeRouter, func(ctx *context.Context) {
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		beego.Info("request cookie token:", ctx.GetCookie("token"))
		return
		if "dev" != beego.AppConfig.String("runmode") {
			return
		}
		accessToken := ctx.GetCookie("token")

		// loginPage := "/lastest/wechatapi/wechat/login?"
		loginPage := "/user/login"
		if "" == accessToken {
			// fmt.Println("======token: ", accessToken)
			if !strings.Contains(ctx.Request.RequestURI, loginPage) {
				beego.Info("======redirect: Contains")
				loginPage = beego.AppConfig.String("loginPage")
				http.Redirect(ctx.ResponseWriter, ctx.Request, loginPage, http.StatusMovedPermanently)
			} else {
				fmt.Println("======", ctx.Request.RequestURI, loginPage)
			}
		} else {
			if !models.ValidateAccessToken(accessToken) {
				fmt.Println("======redirect: ValidateAccessToken")
				loginPage = beego.AppConfig.String("loginPage")
				http.Redirect(ctx.ResponseWriter, ctx.Request, loginPage, http.StatusMovedPermanently)
			}
		}

	})

	ns := beego.NewNamespace("/lastest",
		beego.NSNamespace("/pass", beego.NSInclude(&controllers.MainController{})),
		beego.NSNamespace("/db", beego.NSInclude(&controllers.DBController{})),
		beego.NSNamespace("/lbs", beego.NSInclude(&controllers.BaiduLBSController{})),
		beego.NSNamespace("/wechatapi/small/admin", beego.NSInclude(&controllers.AdminController{})),
		beego.NSNamespace("/wechatapi/small/article", beego.NSInclude(&controllers.ArticleController{})),
		beego.NSNamespace("/wechatapi/small/image", beego.NSInclude(&controllers.ImageController{})),
		beego.NSNamespace("/wechatapi/small/comment", beego.NSInclude(&controllers.CommentController{})),
		beego.NSNamespace("/wechatapi/small/contact", beego.NSInclude(&controllers.ContactController{})),
		beego.NSNamespace("/wechatapi/small/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/wechatapi/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/wechatapi/small/location", beego.NSInclude(&controllers.LocationController{})),
		beego.NSNamespace("/wechatapi/small/config", beego.NSInclude(&controllers.ConfigController{})),
		beego.NSNamespace("/wechatapi/small/reliefstation", beego.NSInclude(&controllers.ReliefStationController{})),
		beego.NSNamespace("/wechatapi/small/volunteer", beego.NSInclude(&controllers.VolunteerController{})),
		beego.NSNamespace("/wechatapi/qcloud/wecos", beego.NSInclude(&controllers.WecosController{})),
		beego.NSNamespace("/wechatapi/wechat/template/message", beego.NSInclude(&controllers.SendWechatTemplateMessage{})),
		// beego.NSNamespace("/wechatapi/small/profile", beego.NSInclude(&controllers.ProfileController{})),
		// 客服系统
		beego.NSNamespace("/wechatapi/customer/service", beego.NSInclude(&controllers.CustomerController{})),
		beego.NSNamespace("/wechatapi/wechat/", beego.NSInclude(&controllers.WechatController{})),
		// 机器人系统
		beego.NSNamespace("/api/robot/", beego.NSInclude(&controllers.RobotController{})),

		// webapi
		beego.NSNamespace("/profile/userinfo", beego.NSInclude(&controllers.UserController{})),

		beego.NSNamespace("/webapi/robot", beego.NSInclude(&controllers.RobotController{})),
		beego.NSNamespace("/webapi/report", beego.NSInclude(&controllers.ReportController{})),
		beego.NSNamespace("/webapi/user", beego.NSInclude(&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
