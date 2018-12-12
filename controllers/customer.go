package controllers

import (
	_ "crypto/hmac"
	"crypto/sha1"
	_ "encoding/base64"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bbhj/baobeihuijia/models"
	"github.com/esap/wechat"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type WechatMsg struct {
	Event        string `json:"Event"`
	SessionFrom  string `json:"SessionFrom"`
	Content      string `json:"Content"`
	CreateTime   int    `json:"CreateTime"`
	FromUserName string `json:"FromUserName"`
	MsgID        int    `json:"MsgId"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
}

// Operations about Users
type CustomerController struct {
	beego.Controller
}

// @Title get
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router / [get]
func (u *CustomerController) Get() {
	signature := u.GetString("signature")
	echostr := u.GetString("echostr")
	timestamp := u.GetString("timestamp")
	nonce := u.GetString("nonce")
	// fmt.Println("customer service", signature, token, timestamp, nonce)
	checkSignature(signature, timestamp, nonce)

	u.Ctx.WriteString(echostr)
}

func checkSignature(signature, timestamp, nonce string) (flag bool) {
	token := "deanhome"
	arr := []string{token, timestamp, nonce}
	fmt.Println("arr: ", arr)
	sort.Sort(sort.StringSlice(arr))
	fmt.Println("arr: ", arr)
	bbb := strings.Join(arr, "")
	fmt.Println("customer service", signature, arr, bbb)

	// mac := hmac.New(sha1.New, []byte(bbb))
	mac := sha1.New()
	io.WriteString(mac, bbb)
	// my_signature := ""
	// mac.Write([]byte(my_signature))
	// macSum := mac.Sum(nil)
	macSum := fmt.Sprintf("%x", mac.Sum(nil))
	if signature == macSum {
		flag = true
	} else {
		fmt.Printf("signature check err, macSum: %s", macSum)
	}
	return
}

// @Title wechat msg xml
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router / [post]
func (u *CustomerController) Post() {
	openid := u.GetString("openid")
	msg := u.GetString("msg")
	fmt.Printf("openid: %s, msg: %s\n", openid, msg)

	wechat.Debug = false
	wechat_appid := beego.AppConfig.String("wechat_appid")

	wechat.Set(beego.AppConfig.String("wechat_token"), wechat_appid, beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))
	wechat.SendText(openid, 0, msg)
	ctx := wechat.VerifyURL(u.Ctx.ResponseWriter, u.Ctx.Request)

	if "user_enter_tempsession" == ctx.Msg.Event {
		fmt.Printf("=====openid: %s,  session from: %s\n", ctx.Msg.FromUserName, ctx.Msg.SessionFrom)
		fmt.Printf("====msg: %v\n", ctx.Msg)
		var customerlogin models.CustomerLogin
		// 数据是xml， 所以使用wechat xml来解析获取
		customerlogin.IP = u.Ctx.Input.IP()
		customerlogin.ToUser = ctx.Msg.ToUserName
		customerlogin.FromUser = ctx.Msg.FromUserName
		customerlogin.CreateTime = ctx.Msg.CreateTime
		customerlogin.MsgType = ctx.Msg.MsgType
		customerlogin.Event = ctx.Msg.Event
		customerlogin.SessionFrom = ctx.Msg.SessionFrom
		fmt.Println(customerlogin)
		models.AddCustomerLogin(customerlogin)
		// 判断是否首次登录
		if models.IsFirstLoginCustomeService(customerlogin.FromUser) {
			// msg := `欢迎智能问答系统。 <a href="CustomerBot" data-miniprogram-appid="` + wechat_appid + `" data-miniprogram-path="pages/hr/main">使用帮助</a>   <a href="http://www.w3school.com.cn/index.html">W3School 在线教程</a>`
			msg = `欢迎使用智能客服系统，您是新用户，建议使用帮助。 <a href="CustomerBot" data-miniprogram-appid="` + wechat_appid + `" data-miniprogram-path="pages/help/main">帮助(暂无此页, 下个版本发布)</a>`
			ctx.NewText(msg).Send()
		} else {
			msg = "I am a robot, :)"
		}
		fmt.Println("Cunstomer return msg: ", msg)

	} else {
		var customertextmsg models.CustomerTextMsg
		customertextmsg.IP = u.Ctx.Input.IP()
		customertextmsg.ToUser = ctx.Msg.ToUserName
		customertextmsg.FromUser = ctx.Msg.FromUserName
		customertextmsg.CreateTime = ctx.Msg.CreateTime
		customertextmsg.MsgType = ctx.Msg.MsgType
		customertextmsg.Content = ctx.Msg.Content
		customertextmsg.MsgId = ctx.Msg.MsgId
		models.AddCustomerTextMsg(customertextmsg)
		fmt.Println("customertextmsg: ", customertextmsg)
		msg = `<a href="http://bbhj.airdb.io">语义识别正在开发中, 敬请期待。 。</a>`

		fmt.Println("isChinese debug:", customertextmsg.MsgType)
		if "text" == customertextmsg.MsgType {
			fmt.Println("isChinese debug:", customertextmsg.Content)
			if isChinese(customertextmsg.Content) {
				nickname := customertextmsg.Content
				fmt.Println("isChinese true, nickname: ", nickname)
				temp, _ := models.GetArticleBySearchBar(nickname, "")
				msg = temp[0].Nickname + "地址：" + temp[0].Province + temp[0].City + temp[0].Address

			} else if isInt(customertextmsg.Content) {
				babyid := customertextmsg.Content
				temp, _ := models.GetArticleBySearchBar("", babyid)
				msg = temp[0].Nickname + "地址：" + temp[0].Province + temp[0].City + temp[0].Address
				msg = `<a href="CustomerBot" data-miniprogram-appid="` + wechat_appid + `" data-miniprogram-path="pages/article/main?data=">` + strconv.FormatInt(temp[0].Babyid, 10) + "-" + temp[0].Nickname + `, 点击进入</a>`
			}
		} else {
		}
		ctx.NewText(msg).Send()

		// 主动通知用户  48小时 5条，慎用。
		// openid = "oPPbr0M2h0geV-jgzUPve9g3x3jg"
		// msg = "这是主动通知"
		// ctx.SendText(openid, 0, msg)
	}

	// wechat.SendText()
	// if "user_enter_tempsession" == wechat.Event {
	// }
	// fmt.Printf("=====", wechat.WxMsg)
}




func isInt(str string) bool {
	var reg = regexp.MustCompile("^[-\\+]?[\\d]+$")
	return reg.MatchString(str)
}

func isChinese(str string) bool {
	var reg = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	return reg.MatchString(str)
}


// @Title get
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /usermsg [get]
func (u *CustomerController) UserMsg() {
}

