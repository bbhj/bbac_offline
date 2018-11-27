package models

// import (
// 	"encoding/json"
// )

type (
	WechatServiceMessage struct {
		ToUser          string `json:"touser"`
		TemplateID      string `json:"template_id"`
		Page            string `json:"page"`
		FormID          string `json:"form_id"`
		EmphasisKeyword string `json:"emphasis_keyword"`
		Data struct {
			Keyword1 struct {
				Value string `json:"value"`
			} `json:"keyword1"`
			Keyword2 struct {
				Value string `json:"value"`
			} `json:"keyword2"`
			Keyword3 struct {
				Value string `json:"value"`
			} `json:"keyword3"`
			Keyword4 struct {
				Value string `json:"value"`
			} `json:"keyword4"`
			Keyword5 struct {
				Value string `json:"value"`
			} `json:"keyword5"`
		} `json:"data"`
	}
)
