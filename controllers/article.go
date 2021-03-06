package controllers

import (
	"encoding/json"
	"github.com/bbhj/bbac/models"
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
	// https://github.com/ipipdotnet/nginx-ipip-module
	// ipip nginx header info: ip |country_code | country| province| city| owner| isp| latitude| longitude
	// fmt.Println(u.Ctx.Request.Header)
	province := u.Ctx.Input.Header("province")
	city := u.Ctx.Input.Header("city")
	// fmt.Printf("====%s, %s\n", province, city)

	aa := models.GetArticlesByCity(province, city, 1)
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

	// https://github.com/ipipdotnet/nginx-ipip-module
	// ipip nginx header info: ip |country_code | country| province| city| owner| isp| latitude| longitude
	// fmt.Println(u.Ctx.Request.Header)
	// province := u.Ctx.Input.Header("province")
	// city := u.Ctx.Input.Header("city")
	// fmt.Printf("====%s, %s\n", province, city)

	// aa := models.GetArticlesByCity(province, city, page)
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
	// babyid, err := u.GetInt64("id")
	babyid := u.GetString("id")

	if models.IsExistsBabyid(babyid) {
		retmsg.Msg = fmt.Sprintf("baby id is exists, id: %s", babyid)
	} else {
		retmsg.Errmsg = fmt.Sprintf("baby id is not exists, id: %s", babyid)
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

// @Title Delete article
// @Description 使用场景：管理员清空测试数据使用(硬删除)
// @Param	uuid		query 	string	false		"The uuid you want to delete"
// @Param	babyid		query	int64	false		"The babyid of article"
// @Param	deleteType		query	int		true		"删除的方式:删一个则deleteType=1，删全部则deleteType=9"
// @Success 200 {object}	models.RetMsg
// @Failure 403 delete fail
// @router /delete [get]
func (u *ArticleController) Delete() {
	var ret models.RetMsg
	babyid, _ := u.GetInt64("babyid")
	uuid := u.GetString("uuid")
	deleteType, _ := u.GetInt64("deleteType")
	fmt.Println("-------recive:", babyid, uuid, deleteType)
	if deleteType == 0 || (deleteType == 1 && babyid == 0 && uuid == "") {
		beego.Error("delete deleteType or babyid or uuid is null")
		ret.Errmsg = "deleteType or babyid or uuid is null"
	}
	res := false
	switch deleteType {
		case 1: {
			res = models.DeleteArticle(babyid, uuid)
			if res {
				ret.Msg = "Delete success, uuid=" + uuid
			} else {
				ret.Errmsg = "Delete fail"
			}
		}
		case 9: {
			res = models.DeleteAllArticle()
			if res {
				ret.Msg = "Delete success"
			} else {
				ret.Errmsg = "Delete fail"
			}
		}
	}
	if ret.Errmsg != "" {
		ret.Status = -1
	} else {
		ret.Status = 0
	}
	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title 返回文章的浏览、评论、转发等数量
// @Description get all Articles
// @Success 200 {object} models.Article
// @router /summary [get]
func (u *ArticleController) Summary() {
	page, _ := u.GetInt("page")

	var overview []models.ArticleOverview
	if (page >= 0) {
		overview = models.GetArticleOverviewByPage(page)	
	} else {
		overview = models.GetArticleOverview()
	}
	u.Data["json"] = overview
	u.ServeJSON()
}

// @Title 访问量自增
// @Description 文章访问量增加 +1
// @Success 200 {object} models.Article
// @router /updateCount [get]
func (u *ArticleController) UpdateCount() {
	babyid, _ := u.GetInt64("babyid")
	column := u.GetString("column")

	var retmsg models.RetMsg
	retmsg.Status = -2
	if ("" != column) {
		switch column {
			case "visit":
				retmsg.Status = models.UpdateArticleVisit(babyid)
			case "comment":
				retmsg.Status = models.UpdateArticleComment(babyid)
			case "forward":
				retmsg.Status = models.UpdateArticleForward(babyid)
			default:
				retmsg.Status = models.UpdateArticleVisit(babyid)
		}
	}
	if (0 == retmsg.Status) {
		retmsg.Msg = column + "count update +1 succcessfully!"
	} else {
		retmsg.Errmsg = column + "count update +1 failed!"
		beego.Error("update failed, column:", column)
	}
	u.Data["json"] = retmsg
	u.ServeJSON()
}

