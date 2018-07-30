package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Comment struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	UUID      string
	AvatarUrl string
	Nickname  string
	Content   string
	Like      int64
	Blow      int
}

func init() {
}

func AddComment(comment Comment) {
	// conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Comment{})

	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	conn.Create(&comment)
	// db.Save(&lbi)
}

func GetComment() (comment Comment) {
	// var lbis []LostBabyInfo
	conn.Last(&comment, 1)

	fmt.Println("%s", comment)
	return
}

func GetCommentsByUUID(uuid string) (comments []Comment) {
	// var lbis []LostBabyInfo
	fmt.Println("GetCommentsByUUID===", uuid)
	conn.Where("uuid = ?", uuid).Order("id desc").Limit(5).Find(&comments)

	// fmt.Println("%s", comments)
	return
}
