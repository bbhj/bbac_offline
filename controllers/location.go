package controllers

import (
	"encoding/json"
	"github.com/bbhj/bbac/models"
	_ "strings"

	"fmt"
	"github.com/astaxie/beego"
	_ "strconv"
	_ "time"
)

// Operations about Articles
type LocationController struct {
	beego.Controller
}

// @Title CreateArticle
// @Description create users
// @Param	body		body 	models.Article	true		"body for user content"
// @Success 200 {int} models.Article.Id
// @Failure 403 body is empty
// @router / [post]
func (u *LocationController) Post() {
	u.Data["json"] = map[string]string{"markers": ""}
	u.ServeJSON()
	return
}

// @Title Markers
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router /markers [get]
func (u *LocationController) Get() {
	userLatitude := u.GetString("latitude")
	userLongitude := u.GetString("longitude")
	fmt.Println("localtion get controller: ", userLatitude, userLongitude)
	data := `{  "markers": [{"iconPath":  "/images/map.png","id": 0,"title": "dean","label": "455112c2-cd5b-4bd5-9bfa-3687c0802da6","latitude":` + userLatitude + `,  "longitude": ` + userLongitude + `,"width": 50,"height": 50}, {"iconPath":  "/images/map.png","id": 1,"title": "dean","label": "455112c2-cd5b-4bd5-9bfa-3687c0802da6","latitude":` + userLatitude + `, "longitude": ` + userLongitude + `,"width": 50,"height": 50 }], "circles": [{"latitude":23.129163, "longitude": 113.264435, "color": "#FF0000DD", "fillColor": "#7cb5ec88", "radius": 3000, "strokeWidth": 1}] }`
	var markers models.LocaltionMarkers
	err := json.Unmarshal([]byte(data), &markers)
	if err != nil {
	}
	// u.Data["json"] = map[string]string{"markers": [{"iconPath":  "/images/map.png"}]}
	// markers.Markers[0].IconPath = "/images/map.png"
	// markers.Markers[0].ID = 0
	// markers.Markers[0].Title = "dean title"
	// markers.Markers[0].Label = "455112c2-cd5b-4bd5-9bfa-3687c0802da6"
	// markers.Markers[0].Latitude = 0
	// markers.Markers[0].Longitude = 0
	// markers.Markers[0].Width = 50
	// markers.Markers[0].Height = 50

	u.Data["json"] = markers
	u.ServeJSON()
}
