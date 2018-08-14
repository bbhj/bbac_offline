package models

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "time"
)

type Ipip struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	// {"platform":"devtools","brand":"devtools","model":"iPhone 7 Plus"}
	IP            string
	Anycast       bool
	BaseStation   string
	City          string
	CityCode      string
	ContinentCode string
	Country       string
	CountryCode   string
	CountryCode3  string
	CurrencyCode  string
	CurrencyName  string
	EuropeanUnion bool
	IDC           string
	ISP           string
	Latitude      string
	Longitude     string
	Organization  string
	PhonePrefix   string
	Province      string
	TimeZone      string
	TimeZone2     string
}

func Insert() {
}
func InsertIpip(ipip Ipip) {
	// ipip.CreatedAt = time.Now()
	// lr.UpdatedAt = time.Now()
	conn.Save(&ipip)

	return
}
