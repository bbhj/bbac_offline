package models

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

func Connect() (db *gorm.DB, err error) {
	conn, err = gorm.Open("mysql", beego.AppConfig.String("gdbc"))

	return
}

func Init() {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Article{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Comment{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Contact{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Login{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&User{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Summary{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Image{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&CustomerLogin{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&CustomerTextMsg{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Ipip{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&BaiduLBS{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&TemplateFormID{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&ServiceTime{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Volunteer{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Profile{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&OpenWeixinAccessToken{})
}
