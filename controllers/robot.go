package controllers

import (
	"github.com/bbhj/baobeihuijia/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"

	_ "encoding/json"
	_ "crypto/hmac"
	_ "crypto/sha1"
	_ "encoding/base64"
	"fmt"
	_ "io"
	_ "math/rand"
	_ "time"
)

// Operations about Users
type RobotController struct {
	beego.Controller
}

// @Title QQ 聊天机器人
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /qq [get]
func (u *RobotController) QQ() {
	// qq=137780017&group=311401898&command=bbhj&message=
	// var rmessage models.QQRobotReciveMessage
	// rmessage.QQ = u.GetString("qq")
	// rmessage.Groupid= u.GetString("group")
	// rmessage.Command = u.GetString("command")
	// rmessage.Message = u.GetString("message")

	var infoList []models.PreForumPost

	keyword := u.GetString("message")

	logs.Info("the reqeust keyword message: ",  keyword)
	if ("" != keyword) {
		infoList = models.GetAllPreForumPostByKeyword(keyword)
	} else {
		infoList = models.GetAllPreForumPostList()
	}
	msg := ""
	for _, info := range infoList {
		datafrom := fmt.Sprintf("https://bbs.baobeihuijia.com/thread-%s-1-1.html", strconv.FormatInt(int64(info.Tid), 10))
		logs.Info("PreForumPost info: ", datafrom, info.Subject)
		msg += info.Subject +"\n" + datafrom  + "\n"	
	}
	msg += "(出处: 宝贝回家论坛)\n"
	u.Ctx.WriteString(msg)
	return
}
// @Title QQ 聊天机器人
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /wechat [get]
func (u *RobotController) Wechat() {
	msg := "(出处: 宝贝回家论坛)\n"
	a := models.TestGetUserInfo()
	logs.Error("=========wechat robot", a.Openid)
	u.Ctx.WriteString(msg)
	return
}
