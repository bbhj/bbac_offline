package models

import (
	_ "fmt"
	"github.com/astaxie/beego"
	_ "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

type (
	Volunteer struct {
		gorm.Model
		Openid  string `gorm:"column:openid;size:64";json:"openid"`
		Email   string 
		QQ      string
		Phone   string
		City	string
		Address string
		Referee string
	}
)

func AddVolunteer(user Volunteer) () {
	beego.Info("add volunteer, openid: ", user.Openid)
	conn.Save(&user)
	return
}

func GetVolunteer() (flag bool) {
        return
}
