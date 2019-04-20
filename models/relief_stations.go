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
	ReliefStation struct {
		gorm.Model
		Name string `json:"name"`
		Province string `json:"province"`
		City string `json:"city"`
		Address string `json:"address"`
		OfficePhone string `json:"officephone"`
	}
)

func AddReliefStation(station ReliefStation) {
	// conn.FirstOrCreate(&station, ReliefStation{OfficePhone: station.OfficePhone})
	conn.Save(&station)
	beego.Info("add....")
	return
}


func GetReliefStations() (stations []ReliefStation) {
	page := 0 
        stepsize := 5
        offset := (page - 1) * stepsize
        conn.Offset(offset).Limit(stepsize).Find(&stations)
        return
}

