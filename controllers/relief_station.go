package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bbhj/bbac/models"
)

type ReliefStationController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param       body            body    models.User     true            "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [get]
func (u *ReliefStationController) Get() {
	var station models.ReliefStation
	
	station.Name  = u.GetString("name")
	station.Province = u.GetString("province")
	station.City = u.GetString("city")
	station.Address = u.GetString("address")
	station.OfficePhone = u.GetString("officephone")

	models.AddReliefStation(station)
	beego.Info("======", station)
	msg := "hello"
	u.Ctx.WriteString(msg)
	return
}
