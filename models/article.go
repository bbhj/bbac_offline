package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Article struct {
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

	MissedProvince string
	MissedCity     string
	MissedCountry  string
	MissedAddress  string
	MissedAt       time.Time
	Handler        string
	Babyid         int64 `gorm:"type:int64;unique_index"`
	Category       string
	Height         string
}

func AddArticle(article Article) (uuid string) {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Article{})

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	article.MissedAt = time.Now()

	fmt.Println("db=====", article)
	err := conn.Create(&article)
	if err.Error != nil {
		fmt.Println("xxxxx, db", err.Error)
	} else {
		uuid = article.UUID
	}
	// db.Save(&lbi)
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

func UpdateArticle() {
}
func DeleteArticle(uuid string) (flag bool) {
	// var lbis []LostBabyInfo
	fmt.Println("delete uuid", uuid)
	conn.Debug().Where("uuid = ?", uuid).Delete(Article{})

	return
}
