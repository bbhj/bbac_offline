package controllers

import (
	_ "encoding/json"

	"github.com/bbhj/bbac/models"
	"github.com/astaxie/beego"
)

// Operations about images
type ImageController struct {
	beego.Controller
}

// @Title Get Image info List
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /getList [get]
func (u *ImageController) GetList() {
	babyid, _ := u.GetInt64("babyid")
	uuid := u.GetString("uuid")
	pictureInfoList := models.GetPictureInfoList(babyid, uuid)
	for i,_ := range pictureInfoList {
		pictureInfoList[i].Attachment = "https://attachment-10016990.file.myqcloud.com/forum/" + pictureInfoList[i].Attachment
	}
	u.Data["json"] = pictureInfoList
	u.ServeJSON()

}
