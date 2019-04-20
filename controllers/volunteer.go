package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/bbhj/bbac/models"
)

type VolunteerController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param       body            body    models.User     true            "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [Post]
func (u *VolunteerController) Get() {
	var user models.Volunteer
	
        json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	// user.Openid = u.GetString("openid")
	// user.City = u.GetString("city")
	// user.Address = u.GetString("address")
	// user.Phone = u.GetString("phone")
	// user.Email = u.GetString("email")
	// user.Referee= u.GetString("referee")

	models.AddVolunteer(user)
	beego.Info("======", user)

	var ret models.RetMsg
	ret.Status = 0
	ret.Msg = "add volunteer successfully."

	u.Data["json"] = ret
	u.ServeJSON()
	return
}

// @Title Get Relief Stations
// @Description Relief Stations
// @Param       body            body    models.User     true            "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /check [get]
func (u *VolunteerController) CheckVolunteer() {

	var ret models.RetMsg
	ret.Status = 0
	ret.Msg = "add volunteer successfully."
	u.Data["json"] = ret
	u.ServeJSON()
	return
}
