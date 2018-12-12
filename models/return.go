package models

import ()

type RetMsg struct {
	Status int8 `json:"status"`
	Msg    string `json:"msg"`
	Errmsg string `json:"errmsg"`
	Data   string `json:"data"`
}
