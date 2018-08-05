package controllers

import (
	_ "encoding/json"
	_ "github.com/bbhj/baobeihuijia/models"
	_ "strings"

	"fmt"
	"github.com/astaxie/beego"
	_ "strconv"
	_ "time"
)

// Operations about Articles
type ConfigController struct {
	beego.Controller
}

// @Title CreateArticle
// @Description create users
// @Param	body		body 	models.Article	true		"body for user content"
// @Success 200 {int} models.Article.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ConfigController) Post() {
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
func (u *ConfigController) Get() {

	WeCosUrl := fmt.Sprintf("https://%s.file.myqcloud.com/files/v2/%s/%s/%s", beego.AppConfig.String("QcloudCosRegion"), beego.AppConfig.String("QcloudCosAPPID"), beego.AppConfig.String("QcloudCosBucket"), beego.AppConfig.String("QcloudCosUploadDir"))

	fmt.Println(WeCosUrl)
	u.Data["json"] = map[string]string{"WeCosUrl": WeCosUrl}
	u.ServeJSON()
}
