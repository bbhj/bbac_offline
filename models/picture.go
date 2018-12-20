package models
import (
	"github.com/jinzhu/gorm"
)

type (
	PictureInfo struct {
		gorm.Model
		Babyid int64	`json:"babyid"`
		UUID	string	`json:"uuid"`
		Filesize int64  `json:"filesize"`
                Width int64  `json:"width"`
                Attachment string `json:"attachment"`

	}
)
