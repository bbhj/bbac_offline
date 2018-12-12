package controllers

import (
	"encoding/json"
	"github.com/bbhj/bbac/models"
	_ "strings"

	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	_ "strconv"
	_ "time"
)

type BaiduLBSResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		AddressComponents struct {
			City     string `json:"city"`
			District string `json:"district"`
			Level    string `json:"level"`
			Province string `json:"province"`
			Street   string `json:"street"`
		} `json:"address_components"`
		Bound            string `json:"bound"`
		CustomAddress    string `json:"custom_address"`
		FormattedAddress string `json:"formatted_address"`
		Location         struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Precise float64 `json:"precise"`
		Source  string  `json:"source"`
	} `json:"result"`
}

// Operations about Articles
type BaiduLBSController struct {
	beego.Controller
}

// @Title Markers
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :uid is empty
// @router / [get]
func (u *BaiduLBSController) Get() {
	city := u.GetString("city")
	address := u.GetString("address")
	if "" != city || "" != address {
		lbsapi := fmt.Sprintf("%s&city=%s&address=%s", beego.AppConfig.String("baidulbsapi"), city, address)
		fmt.Println(lbsapi)
		resp, err := http.Get(lbsapi)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		}

		var baiduResp BaiduLBSResp
		err = json.Unmarshal(body, &baiduResp)
		if err == nil {
			if 0 == baiduResp.Status {
				for _, res := range baiduResp.Result {
					var lbs models.BaiduLBS
					lbs.City = city
					lbs.Address = address
					lbs.Source = res.Source
					lbs.Latitude = res.Location.Lat
					lbs.Longitude = res.Location.Lng
					lbs.Bound = res.Bound
					lbs.FormattedAddress = res.FormattedAddress
					lbs.AddressComponentsProvince = res.AddressComponents.Province
					lbs.AddressComponentsCity = res.AddressComponents.City
					lbs.AddressComponentsDistrict = res.AddressComponents.District
					lbs.AddressComponentsStreet = res.AddressComponents.Street
					lbs.AddressComponentsLevel = res.AddressComponents.Level
					fmt.Println(lbs)
					models.InsertBaiduLBS(lbs)
				}

			}
		}

	}
	u.Data["json"] = "update successfully."
	u.ServeJSON()
}
