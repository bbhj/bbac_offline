package models

import (
	_ "fmt"
	_ "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

type Image struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	OpenID  string
	UnionID string
	UUID    string
	URL     string `gorm:"not null;unique"`
	Status  uint
}

func UploadImage(image Image) (openid string) {
	// update user info
	conn.Save(&image)
	return
}
