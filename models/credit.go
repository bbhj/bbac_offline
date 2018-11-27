package models

import (
	_ "fmt"
	_ "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

type (
	ServiceTime struct {
		gorm.Model
		Openid     string `gorm:"column:openid;size:64";json:"openid"`
		Type	   string 	// 奖励类型
		Duration   int		// 以分钟为单位，默认每次登录奖励3分钟
	}

	Donate struct {
		gorm.Model
		Openid     string `gorm:"column:openid;size:64";json:"openid"`
		Type	   string 	// 捐赠类型
		From	   string	// 捐赠来源
		Money      float64	// 捐赠金额
	}

)

func AddServiceTime(st ServiceTime) {
	// conn.Where("open_id = ?", user.Openid).FirstOrCreate(&user)

	conn.Save(&st)
	return
}
