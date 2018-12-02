package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego/logs"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	PreForumPost struct {
		// gorm.Model
		Message string `json:"message"`
		Subject string `json:"subject"`
		Useip   string `json:"useip"`
		Pid     int64  `json:"pid"`
		Tid     int64  `json:"tid"`
	}
	QQRobotReciveMessage struct {
		gorm.Model
		QQ  string `json:"qq"`
		Groupid string `json:"group"`
		Command  string `json:"command"`
		Message string `json:"message"`
	}
)

func GetAllPreForumPostList() (preForumPost []PreForumPost) {
	err := bbsconn.Table("pre_forum_post").Where("subject != ?", "").Order("tid desc").Limit(5).Find(&preForumPost)
	// bbsconn.Table("pre_forum_post").Where("subject != ?", "").Order("tid desc").Limit(5).Find(preForumPost)
	logs.Error("======bbs models====",err, preForumPost)
	return
}

// 由于subject 包含了姓名和编号，所以只需要这一个接口即可
func GetAllPreForumPostByKeyword(keyword string) (preForumPost []PreForumPost) {
	keyword = "%" + keyword +"%"
	bbsconn.Table("pre_forum_post").Where("subject like ?", keyword).Order("tid desc").Limit(5).Find(&preForumPost)
	return
}
