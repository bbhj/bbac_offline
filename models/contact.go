package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Contact struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	UUID    string
	Name    string
	Wechat  string
	Phone   string
	Mail    string
	Address string
	Remark  string
}

func init() {
}
func AddContact(contact Contact) {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Contact{})

	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	fmt.Println("db=====", contact)
	conn.Create(&contact)
	// db.Save(&lbi)
}

func GetContact() (contact Contact) {
	// var lbis []LostBabyInfo
	conn.First(&contact, 1)

	fmt.Println("%s", contact)
	return
}

func GetContacts() (contacts []Contact) {
	// var lbis []LostBabyInfo
	conn.Limit(5).Find(&contacts)

	fmt.Println("%s", contacts)
	return
}
