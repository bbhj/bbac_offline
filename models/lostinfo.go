package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type LostBabyInfo struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	UUID        string
	AvatarUrl   string
	NickName    string
	Gender      uint
	Province    string
	City        string
	Country     string
	Address     string
	Age         int
	Missingtime time.Time
}

func init() {
}

func AddLostBabyInfo(lbi LostBabyInfo) {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&LostBabyInfo{})

	lbi.CreatedAt = time.Now()
	lbi.UpdatedAt = time.Now()
	fmt.Println("db=====", lbi)
	conn.Create(&lbi)
	// db.Save(&lbi)
}

func GetLostBabyInfo() (lbis []LostBabyInfo) {
	// var lbis []LostBabyInfo
	conn.Limit(5).Find(&lbis)

	fmt.Println("%s", lbis)
	return
}
