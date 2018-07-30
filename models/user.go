package models

import (
	_ "fmt"
	_ "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

type User struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	OpenID    string `gorm:"index"`
	UnionID   string
	Nickname  string
	Gender    uint
	Language  string
	City      string
	Province  string
	Country   string
	AvatarUrl string
}

func AddUserInfo(user User) (openid string) {
	// aaa := conn.Model(&user).Where("open_id = ?", user.OpenID).Updates(user)
	conn.Where("open_id = ?", user.OpenID).FirstOrCreate(&user)

	// update user info
	conn.Save(&user)
	return
}
