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

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ipipdotnet/datx-go"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		fmt.Println("check token=====", ctx.Input.Header("Authorization"))
		ipip(ctx.Input.IP())
		// cookie, err := ctx.Request.Cookie("Authorization")
		// fmt.Println("aaa", cookie, err)

		// fmt.Println("check token=====", ctx.Request.Header("Authorization"))
		// if err != nil || !hjwt.CheckToken(cookie.Value) {
		//  http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		// }
	})

	ns := beego.NewNamespace("/lastest",
		// beego.NSNamespace("/wechat/small/login", beego.NSInclude(&controllers.LoginController{})),
		// beego.NSNamespace("/wechat/small/login", beego.NSInclude(&controllers.LoginController{})),
		beego.NSNamespace("/db/init", beego.NSInclude(&controllers.DBController{})),
		// beego.NSNamespace("/wechatapi/small/topics", beego.NSInclude(&controllers.ArticleController{})),
		beego.NSNamespace("/wechatapi/small/admin", beego.NSInclude(&controllers.AdminController{})),
		beego.NSNamespace("/wechatapi/small/article", beego.NSInclude(&controllers.ArticleController{})),
		// beego.NSNamespace("/wechatapi/small/summary", beego.NSInclude(&controllers.SummaryController{})),
		beego.NSNamespace("/wechatapi/small/comment", beego.NSInclude(&controllers.CommentController{})),
		beego.NSNamespace("/wechatapi/small/contact", beego.NSInclude(&controllers.ContactController{})),
		beego.NSNamespace("/wechatapi/small/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/wechatapi/small/location", beego.NSInclude(&controllers.LocationController{})),
		beego.NSNamespace("/wechatapi/small/config", beego.NSInclude(&controllers.ConfigController{})),
		beego.NSNamespace("/wechatapi/qcloud/wecos", beego.NSInclude(&controllers.WecosController{})),
		beego.NSNamespace("/wechatapi/customer/service", beego.NSInclude(&controllers.CustomerController{})),
	)
	beego.AddNamespace(ns)
}

func ipip(ip string) {
	path := "keys/17monipdb.datx"
	city, err := datx.NewCity(path) // For City Level IP Database
	if err == nil {
		var ipip models.Ipip
		// loc, err := city.FindLocation("218.17.197.194")
		loc, err := city.FindLocation(ip)
		if err == nil {
			// fmt.Println(string(loc.ToJSON()))
			fmt.Println(loc)
			err := json.Unmarshal(loc.ToJSON(), &ipip)
			if err == nil {
				ipip.IP = ip
				fmt.Println(ipip)
				models.InsertIpip(ipip)
			} else {
				fmt.Println("json unmarshal err: ", err)
			}
		}
	}
}
