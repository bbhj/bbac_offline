package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	PreForumPost struct {
		gorm.Model
		Message string `json:"message"`
	}
)

func GetAllPreForumPost() (preForumPost PreForumPost) {
	conn.Limit(3).Find(&preForumPost)
	fmt.Println("aaaaa===")
	return
}
