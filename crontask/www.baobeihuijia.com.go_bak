package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	AvatarUrl string
	Nickname  string
	Babyid    int64
	Category  string
	Height    string
	// 值为1时是男性，值为2时是女性，值为0时是未知
	Gender          uint
	BirthedProvince string
	BirthedCity     string
	BirthedCountry  string
	BirthedAddress  string
	BirthedAt       time.Time

	MissedProvince string
	MissedCity     string
	MissedCountry  string
	MissedAddress  string
	MissedAt       time.Time
	Handler        string
	// Title      string `gorm:"size:60"`
	Title      string
	Arcid      string // 规档编号
	Age        int
	Characters string
	Details    string
	DataFrom   string
	UUID       string
}

func ExampleScrape(id string) {
	// Request the HTML page.
	url := "http://www.baobeihuijia.com/view.aspx?id=" + id
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var article Article
	article.DataFrom = "www.baobeihuijia.com"
	// Find the review items
	// doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
	doc.Find("div#table_1_normaldivl").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		url, _ := s.Find("img").Attr("src")
		url = "http://www.baobeihuijia.com" + url
		article.AvatarUrl = url
		fmt.Printf("logo: %v\n", url)

	})
	doc.Find("div#table_1_normaldivr").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			// title := s.Find("span").Text()
			// fmt.Printf("title : %v\n", title)
			// fmt.Printf("%v\n", s.Eq(0).Text())
			msg := strings.Split(s.Eq(0).Text(), "：")
			// fmt.Println("xxxx:", msg)
			if "寻亲类别" == msg[0] {
				article.Category = msg[1]
			}
			if "寻亲编号" == msg[0] {
				article.Babyid, _ = strconv.ParseInt(msg[1], 10, 64)
			}
			if "姓    名" == msg[0] {
				article.Nickname = msg[1]
			}
			if "性    别" == msg[0] {
				// /* 值为1时是男性，值为2时是女性，值为0时是未知 */
				if "女" == msg[1] {
					// 1 --> 2
					article.Gender = 2
				} else if "男" == msg[1] {
					// 0 -> 1
					article.Gender = 1
				}
			}
			if "出生日期" == msg[0] {
				// article.BirthedAt = msg[1]
				article.BirthedAt, _ = time.ParseInLocation("2006年01月02日", msg[1], time.Local)
				// fmt.Println("xxxx", p)
			}
			if "失踪时身高" == msg[0] {
				article.Height = msg[1]
			}
			if "失踪时间" == msg[0] {
				fmt.Println("xxxxx", msg[1])
				article.MissedAt, _ = time.ParseInLocation("2006年01月02日", msg[1], time.Local)
				// article.MissedAt = msg[1]
			}
			if "失踪人所在地" == msg[0] {
				if "" != msg[1] {
					if strings.Contains(msg[1], ",") {
						article.BirthedProvince = strings.Split(msg[1], ",")[0]
						article.BirthedCity = strings.Split(msg[1], ",")[1]
						article.BirthedAddress = strings.Split(msg[1], ",")[2]
					} else {
						article.BirthedAddress = msg[1]
					}
				}
			}
			if "失踪地点" == msg[0] {
				if "" != msg[1] {
					if strings.Contains(msg[1], ",") {
						article.MissedProvince = strings.Split(msg[1], ",")[0]
						article.MissedCity = strings.Split(msg[1], ",")[1]
						article.MissedAddress = strings.Split(msg[1], ",")[2]
					} else {
						article.MissedAddress = msg[1]
					}
				}
			}
			if "寻亲者特征描述" == msg[0] {
				article.Characters = msg[1]
			}
			if "其他资料" == msg[0] {
				article.Details = msg[1]
			}
			if "跟进志愿者" == msg[0] {
				article.Handler = msg[1]
			}

			// msg := s.Text()
			// fmt.Printf("logo: %v\n", msg)
		})
	})

	auto(article)
}

func auto(article Article) {
	fmt.Println("article: ", article)
	bytesData, err := json.Marshal(article)
	if err != nil {
		fmt.Println("error:", err)
	}
	reader := bytes.NewReader(bytesData)
	url := "http://127.0.0.1:8000/lastest/wechatapi/small/article/new"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

}

func main() {
	idstring := os.Args[1]
	ExampleScrape(idstring)
}
