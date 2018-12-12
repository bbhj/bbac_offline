package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB
var bbsconn *gorm.DB

func Connect() (db *gorm.DB, err error) {
	conn, err = gorm.Open("mysql", beego.AppConfig.String("gdbc"))
	if ( err != nil ) {
		beego.Error("Can't connect database, db info: ", beego.AppConfig.String("gdbc"))
	}
	
	bbsconn, err = gorm.Open("mysql", beego.AppConfig.String("bbs_gdbc"))
	if ( err != nil ) {
		beego.Error("Can't connect database, db info: ", beego.AppConfig.String("gdbc"))
	}

	return
}

func Close() {
	conn.Close()
	bbsconn.Close()
}

func Init() (flag bool, errmsg string) {
	
	var db *gorm.DB
	db = conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Article{})
	db = conn.Debug().Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&ArticleSummary{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Contact{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Login{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&User{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Summary{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Image{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&CustomerLogin{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&CustomerTextMsg{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Ipip{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&BaiduLBS{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&TemplateFormID{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&ServiceTime{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Volunteer{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Profile{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&OpenWeixinAccessToken{})
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Comment{})

	if (db.Error == nil){
		flag = true
	} else {
		errmsg = fmt.Sprintf("%s", db.Error)
	}
	beego.Info("models init info, status:", flag, ", errmsg:", db.Error)
	return
}
