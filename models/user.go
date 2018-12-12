package models

import (
	_ "fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

type (
	WechatLogin struct {
		Code string `json:"code"`
		MobileInfo
		User
	}

	MobileInfo struct {
		OpenID      string  `gorm:"index"`
		Platform    string  `json:"platform"`
		System      string  `json:"system"`
		NetworkType string  `json:"networkType"`
		PhoneBrand  string  `json:"brand"`
		Model       string  `json:"model"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
	}

	// weblogin: Sex, Privilege, Headimgurl
	// mplogin: Gender, AvatarUrl
	User struct {
		gorm.Model
		Openid     string `json:"openid"`
		Unionid    string `json:"unionid"`
		Nickname   string `json:"nickName"`
		Gender     uint   `json:"gender"`
		Sex        uint   `json:"sex"`
		Language   string `json:"language"`
		City       string `json:"city"`
		Province   string `json:"province"`
		Country    string `json:"country"`
		Privilege  string `json:"privilege"`
		AvatarUrl  string `json:"avatarUrl"`
		Headimgurl string `json:"headimgurl"`
	}

	Auth struct {
		Openid      string `json:"openid"`
		Unionid     string `json:"unionid"`
		YourCity    string `json:"yourcity"`
		Token       string `json:"token"`
		AccessToken string `json:"accessToken"`
		IsAdmin     bool   `json:"isAdmin"`
		IsVolunteer bool   `json:"isVolunteer"`
		ServiceTime string `json:"serviceTime"`
		Rights      uint   `json:"rights"`
		WeCosUrl    string `json:"weCosUrl"`
	}

	Profile struct {
		gorm.Model
		Openid       string `gorm:"column:openid;size:64";json:"openid"`
		IsAdmin      bool   `json:"isAdmin"`
		IsVolunteer  bool   `json:"isVolunteer"`
		ServiceTime  int64  `json:"serviceTime"`
		Rights       uint   `json:"rights"`
		IsSubscribed bool   `json:"isSubscribed"`
		IsFirstLogin bool   `json:"isFirstLogin"`
	}

	ContanctnInfo struct {
		Openid  int `gorm:"index"` // 外键 (属于), tag `index`是为该列创建索引
		name    string
		Email   string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
		QQ      string `gorm:"type:varchar(20);unique_index"`
		Phone   string `gorm:"type:varchar(20);unique_index"`
		Address string
	}

	// iview-admin
        UserInfo struct {
                Avator string `json:"avator"`
                Name   string `json:"name"`
                UserId string `json:"user_id"`
                Access string `json:"access"`
        }

)

func AddUserInfo(user User) (openid string) {
	// aaa := conn.Model(&user).Where("openid = ?", user.OpenID).Updates(user)
	conn.Where("openid = ?", user.Openid).FirstOrCreate(&user)

	// update user info
	conn.Save(&user)
	return
}

func GetUserInfoList() (userInfoList []User) {
	page := 1
	stepsize := 5
	offset := (page - 1) * stepsize
	conn.Order("created_at desc").Offset(offset).Limit(stepsize).Find(&userInfoList)
	// fmt.Println("GetArticles: ", len(articles))
	// if 0 == len(articles) {
	//         return GetUserInfoList(page)
	// }
	return
}
func TestGetUserInfo() (userInfo *User) {
	user := &User{}
	logs.Info("TestGetUserInfo=====", conn)
	conn.Debug().First(&user)
	logs.Info("TestGetUserInfo=====", user)

	return user
}
