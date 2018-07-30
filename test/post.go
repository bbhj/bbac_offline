package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

type Article struct {
	AvatarUrl string
	Nickname  string
	Gender    uint
	Province  string
	City      string
	Country   string
	Address   string
	// Title      string `gorm:"size:60"`
	Title      string
	Arcid      string // 规档编号
	Age        int
	Characters string
	Details    string
	DataFrom   string
	UUID       string
}

func main() {
	var article Article

	article.Nickname = "dean"
	article.Title = "title dean"
	article.UUID = "title-dean"

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
