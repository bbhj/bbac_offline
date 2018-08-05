package controllers

import (
	"encoding/json"
	"github.com/bbhj/baobeihuijia/models"
	"github.com/google/uuid"
	"strings"

	"fmt"
	"github.com/astaxie/beego"
	_ "strconv"
	"time"
)

type PostValue struct {
	Article string `json:"article"`
	Code    string `json:"code"`
}

type BabyTime struct {
	BirthedAt string `json:"birthedAt"`
	MissedAt  string `json:"missedAt"`
}

// Operations about Articles
type ArticleController struct {
	beego.Controller
}

// @Title Get Articles Count
// @Description get all Articles count
// @Success 200 {object} models.Article
// @router /count [get]
func (u *ArticleController) GetCount() {
	count := models.GetArticlesCount()
	u.Data["json"] = count
	// u.Data["json"] = map[string]string{"openid": openid}
	u.ServeJSON()
}

// @Title Get Article Infomation
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /:uuid [get]
func (u *ArticleController) Get() {
	uuid := strings.TrimLeft(u.GetString(":uuid"), ":")
	if uuid != "" {
		article, err := models.GetArticle(uuid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = article
		}
	}
	u.ServeJSON()
}

// PUT
// @Title Create Article (API)
// @Description Post Article API (For other system)
// @Param	body		body 	models.Article	true		"body for article content"
// @Success 200 status: {int} , msg: "",  errmsg:"",  data: models.Article.UUID
// @Failure 403 body is empty
// @router /new [post]
func (u *ArticleController) PostArticle() {
	var article models.Article
	var retmsg models.RetMsg
	// fmt.Println(u.Ctx.Input.RequestBody)
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &article)

	if err != nil {
		fmt.Println(article)
		retmsg.Errmsg = fmt.Sprintf("Unmarshal post json data error, msg: %s.", err)
		fmt.Println(retmsg.Errmsg)
	}

	if models.IsExistsBabyid(article.Babyid) {
		retmsg.Errmsg = fmt.Sprintf("Baby id alreay exists, id: %d.", article.Babyid)
	} else {

		if "" != article.UUID {
			_, err = uuid.Parse(article.UUID)
			if err != nil {
				retmsg.Errmsg = fmt.Sprintf("post artcile uuid invalid, error: %s", err)
				fmt.Println(retmsg.Errmsg)
			}
		}
		article.UUID = uuid.New().String()
		fmt.Println("recive :", article)
		uuid := models.AddArticle(article)
		if "" == uuid {
			retmsg.Errmsg = fmt.Sprintf("Insert artcile to db, error: %s", uuid)
		}
	}

	if "" != retmsg.Errmsg {
		retmsg.Status = -1
	}
	if 0 == retmsg.Status {
		retmsg.Data = article.UUID
	}
	u.Data["json"] = retmsg
	u.ServeJSON()
	return
}

// @Title CreateArticle
// @Description create users
// @Param	body		body 	models.Article	true		"body for user content"
// @Success 200 {int} models.Article.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ArticleController) Post() {
	var postvalue PostValue
	json.Unmarshal(u.Ctx.Input.RequestBody, &postvalue)
	fmt.Println("post value=====", postvalue.Article)
	fmt.Println("post value=====", postvalue.Article)

	var article models.Article
	err := json.Unmarshal([]byte(postvalue.Article), &article)
	fmt.Println("已知错误：", article)

	var babytime BabyTime
	err = json.Unmarshal([]byte(postvalue.Article), &babytime)
	// loc, _ := time.LoadLocation("Local")
	article.BirthedAt, err = time.ParseInLocation("2006-01-02", babytime.BirthedAt, time.Local)
	if nil != err {
		fmt.Println("BirthedAt Parse to time", err)
	}
	// article.MissedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", babytime.MissedAt, loc)
	article.MissedAt, err = time.ParseInLocation("2006-01-02 15:04:05", "2008-03-03 00:00:00", time.Local)
	if nil != err {
		fmt.Println("MissedAt Parse to time", err)
	}
	fmt.Println(article)
	uuid := models.AddArticle(article)

	var summary models.Summary
	summary.UUID = uuid

	models.AddSummary(summary)
	u.Data["json"] = map[string]string{"status": "success", "uuid": uuid}
	u.ServeJSON()
	return
}

// @Title GetAll
// @Description get all Articles
// @Success 200 {object} models.Article
// @router / [get]
func (u *ArticleController) GetAll() {
	aa := models.GetArticles(1)
	u.Data["json"] = aa
	// u.Data["json"] = map[string]string{"openid": openid}
	u.ServeJSON()
}

// @Title Get Article by baby id
// @Description get all Articles
// @Success 200 {object} models.Article
// @router /keywords [get]
func (u *ArticleController) GetArticleByKeywords() {
	nickname := u.GetString("nickname")
	babyid := u.GetString("babyid")
	fmt.Println("=====", nickname, babyid)

	aa, _ := models.GetArticleBySearchBar(nickname, babyid)
	u.Data["json"] = aa
	// u.Data["json"] = map[string]string{"openid": openid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Articles
// @Success 200 {object} models.Article
// @router /topics [get]
func (u *ArticleController) GetTopics() {
	page, err := u.GetInt("page")
	if err != nil {
		page = 1
	}
	fmt.Println("get topics, page:", u.GetString("page"))
	aa := models.GetArticles(page)
	u.Data["json"] = aa
	// u.Data["json"] = map[string]string{"openid": openid}
	u.ServeJSON()
}

// @Title Check baby id exist or not
// @Description get all Articles
// @Success 200 {object} models.Article
// @router /checkid [get]
func (u *ArticleController) CheckId() {
	var retmsg models.RetMsg
	babyid, err := u.GetInt64("id")
	if err != nil {
		retmsg.Errmsg = fmt.Sprintf("baby id not int, id: %d", babyid)
		retmsg.Status = -1
	}

	if models.IsExistsBabyid(babyid) {
		retmsg.Msg = fmt.Sprintf("baby id is exists, id: %d", babyid)
	} else {
		retmsg.Errmsg = fmt.Sprintf("baby id is not exists, id: %d", babyid)
		retmsg.Status = -1
	}
	u.Data["json"] = retmsg
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Article	true		"body for user content"
// @Success 200 {object} models.Article
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *ArticleController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.Article
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		// uu, err := models.UpdateArticle(uid, &user)
		// if err != nil {
		// 	u.Data["json"] = err.Error()
		// } else {
		// 	u.Data["json"] = uu
		// }
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the artile
// @Param	uuid		path 	string	true		"The uuid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uuid [delete]
func (u *ArticleController) Delete() {
	uuid := u.GetString(":uuid")
	models.DeleteArticle(uuid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
