package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CustomerLogin struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	// {"platform":"devtools","brand":"devtools","model":"iPhone 7 Plus"}
	IP          string
	ToUser      string
	FromUser    string
	CreateTime  int64
	MsgType     string
	Event       string
	SessionFrom string
}

func (cl *CustomerLogin) String() string {
	return fmt.Sprintf("IP: %s, OpenID: %s, FromUser: %s ,CreateTime: ", cl.IP, cl.FromUser, cl.ToUser, cl.CreateTime)
}

func AddCustomerLogin(cl CustomerLogin) {
	conn.Save(&cl)
	return
}
func IsFirstLoginCustomeService(openid string) (flag bool) {
	count := 0
	conn.Where("openid = ?", openid).Count(&count)
	if count > 0 {
		flag = true
	}
	return
}

func GetCustomerLogin() {
	var record CustomerLogin
	conn.First(&record, 3)
	fmt.Println("==============", record.String())

}
