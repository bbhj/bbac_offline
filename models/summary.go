package models

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Summary struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	UUID     string
	OpenID   string
	Category string
	PicUrl   string
	Title    string
	Status   int // 0 发布成功; 1 草稿箱; -1 平台屏蔽
	Like     int64
	Blow     int64
}

func init() {
}

func AddSummary(summary Summary) {

	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Summary{})
	summary.CreatedAt = time.Now()
	summary.UpdatedAt = time.Now()
	conn.Create(&summary)
	// db.Save(&lbi)
}

func GetSummary() (summary Summary) {
	// var lbis []LostBabyInfo
	conn.First(&summary, 1)

	fmt.Println("%s", summary)
	return
}

func GetSummaries() (summaries []Summary) {
	// var lbis []LostBabyInfo
	conn.Order("id desc").Limit(4).Find(&summaries)

	fmt.Println("%s", summaries)
	return
}
func GetSummariesByOpenID(openid string, status int) (summaries []Summary) {
	conn.Where("open_id = ? and status = ?", openid, status).Find(&summaries)

	fmt.Println("%s", summaries)
	return
}
func DeleteSummary(uuid string) (flag bool) {

	fmt.Println("delete", uuid)
	// var lbis []LostBabyInfo
	// db.Table("summaries").Where("uuid = ?", uuid).Delete()
	conn.Where("uuid = ?", uuid).Delete(Summary{})

	return
}
