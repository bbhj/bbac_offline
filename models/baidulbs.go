package models

// import (
// 	"encoding/json"
// )

// 参考地址：http://lbsyun.baidu.com/index.php?title=lbscloud/cloudgc/explain
// http://api.map.baidu.com/cloudgc/v1?geotable_id=******&address=北京市海淀区营业网点&city=北京&ak=你的ak //GET请求
type BaiduLBS struct {
	City                      string // input
	Address                   string // input
	GeotableId                int    // input (reserve)
	Source                    string
	Latitude                  float64
	Longitude                 float64
	Bound                     string
	FormattedAddress          string
	CustomAddress             string
	AddressComponentsProvince string
	AddressComponentsCity     string
	AddressComponentsDistrict string
	AddressComponentsStreet   string
	AddressComponentsLevel    string  //解析结果的地址级别，省、城市、区县，如果解析到POI或门牌号，则level为POI类型、道路等
	Precise                   float64 // 1.x表示精确坐标，0.x表示模糊坐标
}

func InsertBaiduLBS(lbs BaiduLBS) {
	conn.Save(&lbs)
	return
}
