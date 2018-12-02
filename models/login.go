package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Login struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	// {"platform":"devtools","brand":"devtools","model":"iPhone 7 Plus"}
	IP          string
	Openid      string `json:"openid"`
	Unionid     string `json:"unionid"`
	Platform    string
	NetworkType string
	Brand       string
	Pmodel      string
	Longitude   float64
	Latitude    float64
}

func (lr *Login) String() string {
	return fmt.Sprintf("IP: %s, Openid: %s, Unionid: %s , Location: {%f, %f}", lr.IP, lr.Openid, lr.Unionid, lr.Longitude, lr.Latitude)
}

func init() {
}
func AddLogin(lr Login) {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Login{})
	id := uuid.New().String()
	fmt.Println(id)

	lr.CreatedAt = time.Now()
	lr.UpdatedAt = time.Now()
	conn.Save(&lr)

	return
}

func GetLogin() {
	var record Login
	conn.First(&record, 3)
	fmt.Println("==============", record.String())
}
