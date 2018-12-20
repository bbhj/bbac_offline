package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"strconv"
)

type (
	Article struct {
		gorm.Model
		// ID        string `gorm:"primary_key"`
		//  Timestamp int64
		// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
		UUID      string
		AvatarUrl string
		Nickname  string
		// 值为1时是男性，值为2时是女性，值为0时是未知
		Gender   uint
		Province string
		City     string
		Country  string
		Address  string
		// Title      string `gorm:"size:60"`
		Title           string
		Arcid           string // 规档编号
		Age             int
		Characters      string
		Details         string
		DataFrom        string
		BirthedProvince string
		BirthedCity     string
		BirthedCountry  string
		BirthedAddress  string
		BirthedAt       time.Time

		MissedCountry  string
		MissedProvince string
		MissedCity     string
		MissedAddress  string
		MissedAt       time.Time
		Handler        string
		Babyid         int64 `gorm:"type:int64;unique_index"`
		Category       string
		Height         string
	}

	ArticleSummary struct {
		gorm.Model
		Babyid int64 // `gorm:"type:int64"`
		// UUID        string `gorm:"type:string;unique_index"`
		UUID string
		// status, 0 未找到， 1 已找回, 其他预留，如紧急
		Status  int `gorm:"type:int;default:0"`
		Visit   int64
		Forward int64
		Comment int64
	}

	ArticleOverview struct {
		Babyid int64
		UUID   string
		// status, 0 未找到， 1 已找回, 其他预留，如紧急
		Status    int `gorm:"type:int;default:0"`
		Visit     int64
		Forward   int64
		Comment   int64
		DataFrom  string
		Title     string
		Nickname  string
		AvatarUrl string
		// 值为1时是男性，值为2时是女性，值为0时是未知
		Gender         uint
		BirthedAt      time.Time
		Characters     string
		Details        string
		MissedProvince string
		MissedCity     string
		MissedAddress  string
		MissedAt       time.Time
	}
)

func AddArticle(article Article) (uuid string) {
	// err := conn.Debug().Where(Article{Babyid: article.Babyid}).FirstOrCreate(&article)
	// err := conn.Debug().FirstOrCreate(&article, Article{Babyid: article.Babyid})
	beego.Error(article.Babyid)
	err := conn.Debug().FirstOrCreate(&article, Article{Babyid: 12248})
	if err.Error != nil {
		fmt.Println("xxxxx, db", err.Error)
	} else {
		uuid = article.UUID
		var articleSummary ArticleSummary
		articleSummary.UUID = article.UUID
		articleSummary.Babyid = article.Babyid
		AddArticleSummary(articleSummary)
	}
	beego.Info("===!!!!=", article)
	
	return
}

func AddArticleSummary(articleSummary ArticleSummary) {
	conn.Save(&articleSummary)
	return
}

func GetArticle(uuid string) (article Article, err error) {

	conn.Where("uuid = ?", uuid).Last(&article)
	fmt.Println("%s", article)
	return
}

func GetArticleBySearchBar(nickname, babyid string) (article []Article, err error) {
	//keyword = fmt.Sprintf("\"%%%v%%\"", keyword)
	// nickname = fmt.Sprintf("%%%v%%", nickname)
	nickname = fmt.Sprintf("%%%s%%", nickname)
	babyid = fmt.Sprintf("%%%s%%", babyid)
	// conn.Debug().Where("nickname like ?", keyword).Order("missed_at desc").Limit(5).Find(&article)
	conn.Debug().Where("nickname like ? and babyid like ?", nickname, babyid).Order("missed_at desc").Limit(5).Find(&article)
	// conn.Last(&article)
	fmt.Println("%s", article)
	return
}

func GetArticles(page int) (articles []Article) {
	stepsize := 5
	offset := (page - 1) * stepsize
	conn.Order("missed_at desc").Offset(offset).Limit(stepsize).Find(&articles)
	fmt.Println("GetArticles: ", articles)
	return
}

func GetArticlesByCity(province, city string, page int) (articles []Article) {
	stepsize := 5
	offset := (page - 1) * stepsize
	conn.Where("missed_province = ? or missed_city = ?", province, city).Order("missed_at desc").Offset(offset).Limit(stepsize).Find(&articles)
	// fmt.Println("GetArticles: ", len(articles))
	if 0 == len(articles) {
		return GetArticles(page)
	}
	return
}

func IsExistsBabyid(babyid int64) (flag bool) {
	count := 0
	conn.Debug().Table("articles").Where("babyid = ?", babyid).Count(&count)
	fmt.Println("cound: ", count)
	if count > 0 {
		flag = true
	}
	return
}

func GetArticlesCount() (count int64) {
	conn.Table("articles").Where("deleted_at IS NULL").Count(&count)
	// fmt.Println("count ss:", count)
	return
}

func UpdateArticle(article Article) {
	// article.CreatedAt = time.Now()
	// article.UpdatedAt = time.Now()
	// article.MissedAt = time.Now()

	beego.Info(article)
	conn.Debug().Save(&article)
}

func DeleteArticle(uuid string) (flag bool) {
	beego.Info("delete uuid:", uuid)
	conn.Debug().Where("uuid = ?", uuid).Delete(Article{})

	return
}

func GetArticleSummary() {
	// var article Article
	// conn.First(&article).Preload("Summary").Related(&article.Summary)
	// var articleSummary ArticleSummary
	var overview ArticleOverview
	// rows, _ := conn.Debug().Table("articles").Joins("left join article_summaries on articles.babyid = article_summaries.babyid").Rows()
	rows, _ := conn.Debug().Raw("select * from articles, article_summaries where articles.babyid = article_summaries.babyid limit 1;").Rows()
	for rows.Next() {
		// rows.Scan(&name, &age, &email)
		// conn.ScanRows(rows, &article )
		// beego.Info("=======", article)

		conn.ScanRows(rows, &overview)
		beego.Info("=======", overview)
		break
	}
	// conn.Debug().Model(&article).Related(&articleSummary)
	// conn.Raw(select * from articles as a , article_summaries as b  where a.babyid = b.babyid").Scan(
	// conn.Preload("Article.Summary").First(&article)
	// conn.Debug().Preload("ArticleSummary").Where("babyid = ?", 319169).Find(&articleSummary)
	// conn.Debug().Preload("ArticleSummary").Where("babyid = ?", 319169).First(&article)
	// conn.Debug().Preload("Article").Where("babyid = ?", 319169).First(&articleSummary)
	// beego.Info("============GetArticleSummary=========:", article)
	// beego.Info("============GetArticleSummary=========:", articleSummary)
}

func GetArticleOverview() (overviews []ArticleOverview) {
	sqltext := "select * from articles, article_summaries where articles.babyid = article_summaries.babyid limit 5;"
	rows, _ := conn.Debug().Raw(sqltext).Rows()
	var overview ArticleOverview
	for rows.Next() {
		conn.ScanRows(rows, &overview)
		overviews = append(overviews, overview)
		beego.Info("=======", overview)
	}
	return
}

func GetArticleOverviewByPage(page int) (overviews []ArticleOverview) {
	step := 2
	sqltext := "select * from articles, article_summaries where articles.babyid = article_summaries.babyid"
	sqltext += " limit " + strconv.Itoa(step)
	sqltext += " offset " + strconv.Itoa(step * page)
	sqltext += " ;"
	rows, _ := conn.Debug().Raw(sqltext).Rows()
	var overview ArticleOverview
	for rows.Next() {
		conn.ScanRows(rows, &overview)
		overviews = append(overviews, overview)
		beego.Info("=======", overview)
	}
	return
}

// 更新总览
func UpdateArticleVisit(babyid int64) ( status int8) {
	db := conn.Table("article_summaries").Where("babyid = ?", babyid).UpdateColumn("Visit", gorm.Expr("visit + ?", 1))
	if db.Error != nil {
		// db有错误，则为-1
		status = -1
	}
	return
}

func UpdateArticleComment(babyid int64) (status int8) {
	db := conn.Table("article_summaries").Where("babyid = ?", babyid).UpdateColumn("Comment", gorm.Expr("comment + ?", 1))
	if db.Error != nil {
		// db有错误，则为-1
		status = -1
	}
	return
}

func UpdateArticleForward(babyid int64) (status int8) {
	db := conn.Table("article_summaries").Where("babyid = ?", babyid).UpdateColumn("Forward", gorm.Expr("forward + ?", 1))
	if db.Error != nil {
		// db有错误，则为-1
		status = -1
	}
	return
}

func UpdateArticleStatus(babyid int64) (status int8) {
	// 设置为1表示已找回。
	db := conn.Table("article_summaries").Where("babyid = ?", babyid).Update("Status", 1)
	if db.Error != nil {
		// db有错误，则为-1
		status = -1
	}
	return
}
