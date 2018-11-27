package controllers

import (
	"encoding/json"
	"github.com/bbhj/baobeihuijia/models"
	_ "strings"

	"fmt"
	"github.com/astaxie/beego"
	_ "strconv"
	_ "time"
)

// Operations about Articles
type SendWechatTemplateMessage struct {
	beego.Controller
}

// @Title CreateArticle
// @Description create users
// @Param	body		body 	models.Article	true		"body for user content"
// @Success 200 {int} models.Article.Id
// @Failure 403 body is empty
// @router / [post]
func (u *SendWechatTemplateMessage) Post() {
	u.Data["json"] = map[string]string{"status": ""}
	u.ServeJSON()
	return
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router / [get]
func (u *SendWechatTemplateMessage) Get() {
	var sm models.WechatServiceMessage
	sm.ToUser = "oPPbr0M2h0geV-jgzUPve9g3x3jg"
	sm.TemplateID = beego.AppConfig.String("wechat_templateMessageId")
	sm.Data.Keyword1.Value = "帮忙寻找xxx"
	sm.Data.Keyword2.Value = "阿正_Dean"
	sm.Data.Keyword3.Value = "常州市新北区河海东路200号8栋2单元301"
	sm.Data.Keyword4.Value = "这是备注"
	sm.Data.Keyword5.Value = "已有500人参与提供帮助"

	aa, err := json.Marshal(sm)
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	fmt.Println ( "WechatTemplateMessage: ",   string(aa) )


	fmt.Println(beego.AppConfig.String("wechatApiSendTemplateMessage") )
	// u.Data["json"] = map[string]string{"WeCosUrl": WeCosUrl}
	u.ServeJSON()
}
