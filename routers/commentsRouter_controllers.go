package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:openid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "CheckId",
			Router: `/checkid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetCount",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetArticleByKeywords",
			Router: `/keywords`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "PostArticle",
			Router: `/new`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Summary",
			Router: `/summary`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetTopics",
			Router: `/topics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "UpdateCount",
			Router: `/updateCount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:BaiduLBSController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:BaiduLBSController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ConfigController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ConfigController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ContactController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CustomerController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CustomerController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CustomerController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "UserMsg",
			Router: `/usermsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:DBController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:DBController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ImageController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ImageController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/getList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/markers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ReliefStationController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ReliefStationController"],
		beego.ControllerComments{
			Method: "GetReliefStations",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ReliefStationController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ReliefStationController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ReportController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:ReportController"],
		beego.ControllerComments{
			Method: "SaveErrorLogger",
			Router: `/save_error_logger`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:RobotController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:RobotController"],
		beego.ControllerComments{
			Method: "QQ",
			Router: `/qq`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:RobotController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:RobotController"],
		beego.ControllerComments{
			Method: "Wechat",
			Router: `/wechat`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SendWechatTemplateMessage"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SendWechatTemplateMessage"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SendWechatTemplateMessage"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SendWechatTemplateMessage"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:SummaryController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "FormID",
			Router: `/formid`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetDragList",
			Router: `/get_drag_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetInfo",
			Router: `/get_info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "WxLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateUserInfo",
			Router: `/updateUserInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WechatController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WechatController"],
		beego.ControllerComments{
			Method: "CreateQRcode",
			Router: `/createqrcode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WechatController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WechatController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WechatController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WechatController"],
		beego.ControllerComments{
			Method: "UserInfoList",
			Router: `/userinfolist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WecosController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WecosController"],
		beego.ControllerComments{
			Method: "Auth",
			Router: `/auth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WecosController"] = append(beego.GlobalControllerRouter["github.com/bbhj/bbac/controllers:WecosController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
