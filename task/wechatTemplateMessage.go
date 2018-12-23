package task

import (
	"github.com/astaxie/beego"
	"github.com/bbhj/bbac/models"
	"github.com/imroc/req"
	"github.com/esap/wechat"
)

func SendWechatTemplateMessage() {
	var tm models.TemplateMessage
	tm.Touser = "oPPbr0M2h0geV-jgzUPve9g3x3jg"
	// tm.TemplateID = "nYL5mnIUq5aybu8uZWaF1qwWPp_6Up4EhhUmrbHXWmc"
	// tm.Page = "/pages/home/main"
	// tm.EmphasisKeyword = ""
	// tm.Data.Keyword1.Value = "帮忙寻找xxx"
	// tm.Data.Keyword2.Value = "阿正_Dean"
	// tm.Data.Keyword3.Value = "常州市新北区河海东路200号8栋2单元301"
	// tm.Data.Keyword4.Value = "这是备注"
	// tm.Data.Keyword5.Value = "已有500人参与提供帮助"

	// 签到提醒
	tm.TemplateID = "i_6YW7gfA35W_DNHJYmOMpMWTX24U9vwAEE1QbAjRmQ"
	tm.Page = "/pages/profile/main"
	tm.EmphasisKeyword = ""
	tm.Data.Keyword1.Value = "每日签到"
	tm.Data.Keyword2.Value = "叮！该签到啦~锲而不舍,金石可镂。"
	tm.Data.Keyword3.Value = "本周累计签到2次，加油！"
	tm.Data.Keyword4.Value = "每天最晚签到时间为 02:00。请确认代码是否提交，小程序体验版是否提交。"


	tm.FormID = models.GetFormByOpenid(tm.Touser)

	apiurl := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
	apiurl += wechat.GetAccessToken()
	a2,err := req.Post(apiurl, req.BodyJSON(tm))
	if err != nil {
		beego.Error("post api failed, send wechat template message failed!")
	}
	var errret models.WechatErrorReturn
	a2.ToJSON(&errret)
	beego.Info("-----", errret)
}
