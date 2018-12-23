package models

import ()

type RetMsg struct {
	Status int8 `json:"status"` // 0 是成功，非 0 为失败，且返回 Errmsg
	Msg    string `json:"msg"` // 当 Status = 0 时，反回成功信息，非 0 时，此字段为空
	Errmsg string `json:"errmsg"` // 当 Status 非 0 时，返回错误信息
	Data   string `json:"data"`  // 返回用户所需数据
}
