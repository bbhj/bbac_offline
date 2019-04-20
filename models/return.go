package models

import ()

type RetMsg struct {
	Status int8   `json:"status"` // 本系统内部执行状态码， 0 表示正常
	Msg    string `json:"msg"` // 当 Errcode 非 0 时，反回成功信息, 为 0 时，此字段为空
	Errcode   int    `json:"errcode"` // 其他系统返回错误码
	Errmsg string `json:"errmsg"` // 当 Status 非 0 时，返回错误信息
	Data   string `json:"data"`  // 返回用户所需数据
}
