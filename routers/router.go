// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/bbhj/baobeihuijia/controllers"
	"github.com/bbhj/baobeihuijia/models"

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
		if ("dev" != beego.AppConfig.String("runmode") ){
			return
		}
		// cookie, _ := ctx.Request.Cookie("access_token")
		accessToken := ctx.GetCookie("access_token")
		fmt.Println("======access_token: ", accessToken)

		// loginPage := "/lastest/wechatapi/wechat/login?"
		loginPage := "/wechatapi/wechat/login?"
		if "" == accessToken {
			if !strings.Contains(ctx.Request.RequestURI, loginPage) {
				fmt.Println("======redirect: Contains")
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
		beego.NSNamespace("/db/init", beego.NSInclude(&controllers.DBController{})),
		beego.NSNamespace("/lbs", beego.NSInclude(&controllers.BaiduLBSController{})),
		beego.NSNamespace("/wechatapi/small/admin", beego.NSInclude(&controllers.AdminController{})),
		beego.NSNamespace("/wechatapi/small/article", beego.NSInclude(&controllers.ArticleController{})),
		beego.NSNamespace("/wechatapi/small/comment", beego.NSInclude(&controllers.CommentController{})),
		beego.NSNamespace("/wechatapi/small/contact", beego.NSInclude(&controllers.ContactController{})),
		beego.NSNamespace("/wechatapi/small/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/wechatapi/small/location", beego.NSInclude(&controllers.LocationController{})),
		beego.NSNamespace("/wechatapi/small/config", beego.NSInclude(&controllers.ConfigController{})),
		beego.NSNamespace("/wechatapi/qcloud/wecos", beego.NSInclude(&controllers.WecosController{})),
		beego.NSNamespace("/wechatapi/wechat/template/message", beego.NSInclude(&controllers.SendWechatTemplateMessage{})),
		// beego.NSNamespace("/wechatapi/small/profile", beego.NSInclude(&controllers.ProfileController{})),
		// 客服系统
		beego.NSNamespace("/wechatapi/customer/service", beego.NSInclude(&controllers.CustomerController{})),
		beego.NSNamespace("/wechatapi/wechat/", beego.NSInclude(&controllers.WechatController{})),
	)
	beego.AddNamespace(ns)
}
