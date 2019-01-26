package models
import (
	"github.com/jinzhu/gorm"
)

type (
	PictureInfo struct {
		gorm.Model
		Babyid  string `json:"babyid"`
		UUID	string	`json:"uuid"`
		Filesize int64  `json:"filesize"`
                Width int64  `json:"width"`
                Attachment string `json:"attachment"`

	}
)

func GetPictureInfoList(babyid int64, uuid string) (pictureInfo []PictureInfo) {
        conn.Debug().Where("babyid = ? or uuid = ?", babyid, uuid).Find(&pictureInfo)
        return
}
