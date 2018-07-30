package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:openid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetCount",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "PostArticle",
			Router: `/new`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetArticleByKeywords",
			Router: `/keywords`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetTopics",
			Router: `/topics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "CheckId",
			Router: `/checkid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ConfigController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ConfigController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CustomerController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CustomerController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:DBController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:DBController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/markers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:WecosController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:WecosController"],
		beego.ControllerComments{
			Method: "Auth",
			Router: `/auth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:WecosController"] = append(beego.GlobalControllerRouter["github.com/airdb/baobeihuijia/controllers:WecosController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
