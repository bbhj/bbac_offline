package models

// import (
// 	"encoding/json"
// )

type (
	// WxMsg 混合用户消息，业务判断的主体
	WxMsg struct {
		ToUserName   string `json:"ToUserName"`
		FromUserName string `json:"FromUserName"`
		CreateTime   int64  `json:"CreateTime"`
		MsgId        int64  `json:"MsgId"`
		MsgType      string `json:"MsgType"`
		Content      string `json:"Content"`
	}

	JsCode struct {
		SesssionKey string `json:"session_key"`
		OpenID      string `json:"openid"`
		UnionID     string `json:"unionid"`
	}
)
