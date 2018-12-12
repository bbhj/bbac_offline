package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CustomerTextMsg struct {
	gorm.Model
	IP         string
	ToUser     string
	FromUser   string
	CreateTime int64
	MsgType    string
	Content    string
	MsgId      int64
}

func (ctm *CustomerTextMsg) String() string {
	return fmt.Sprintf("IP: %s, OpenID: %s, FromUser: %s ,CreateTime: ", ctm.IP, ctm.FromUser, ctm.ToUser, ctm.CreateTime)
}

func AddCustomerTextMsg(cTextMsg CustomerTextMsg) {
	conn.Save(&cTextMsg)
	return
}
func GetCustomerTextMsg(cTextMsg CustomerTextMsg) {
	page := 1
	stepsize := 5
	offset := (page - 1) * stepsize
	conn.Order("created_at desc").Offset(offset).Limit(stepsize).Find(&cTextMsg)
	return
}
