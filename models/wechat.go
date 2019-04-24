package models

import (
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type (
	WechatAuth struct {
		gorm.Model
		Code        string `json:"code"`
		State       string `json:"code"`
		Scene       string `json:"scene"`
		Path        string `json:"path"`
		ShareTicket string `json:"shareTicket"`
	}

	AnalysisDailySummary struct {
		List []struct {
			RefDate    string `json:"ref_date"`
			SharePv    int    `json:"share_pv"`
			ShareUv    int    `json:"share_uv"`
			VisitTotal int    `json:"visit_total"`
		} `json:"list"`
	}

	WeixinDailySummary struct {
		gorm.Model
		RefDate    string `json:"ref_date"`
		SharePv    int    `json:"share_pv"`
		ShareUv    int    `json:"share_uv"`
		VisitTotal int    `json:"visit_total"`
	}

	TemplateMessage struct {
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
		EmphasisKeyword string `json:"emphasis_keyword"`
		FormID          string `json:"form_id"`
		Page            string `json:"page"`
		TemplateID      string `json:"template_id"`
		Touser          string `json:"touser"`
	}

	// 发送模板消息
	TemplateFormID struct {
		gorm.Model
		Openid string `json:"openid"`
		Formid string `json:"formid"`
		Status uint   // -1 发送消息失败， 0 已通知，1,初始化  2 已过期, 剩余保留
	}

	// 获取小程序码
	// 参考：https://developers.weixin.qq.com/miniprogram/dev/api/open-api/qr-code/getWXACodeUnlimit.html
	QRCodeRequestParms struct {
		AccessToken string `json:"access_token"`
		Scene       string `json:"scene"`
		Page        string `json:"page"`
		Width       string `json:"width"`
		AutoColor   string `json:"auto_color"`
		LineColor   string `json:"line_color"`
		IsHyaline   string `json:"is_hyaline"`
	}

	QRCodeReturn struct {
		Errcode int `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	WechatErrorReturn struct {
		Errcode int `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	// 网页登录 access_token
	OpenWeixinAccessToken struct {
		gorm.Model
		Openid       string `json:"openid"`
		Unionid      string `json:"unionid"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		SessionKey   string `json:"session_key"` // 小程序登录session_key
	}

	WechatID struct {
		Openid	     string `json:"openid"`	
		Unionid      string `json:"unionid"`
	}

	WechatLoginScene struct {
		gorm.Model
		Openid	     string `json:"openid"`
		Scene        int    `json:"scene"`
		Path         string `json:"path"`	
		ShareTicket  string `json:"shareTicket"`
	}
)

func AddTemplateFormID(tformid TemplateFormID) {
	tformid.CreatedAt = time.Now()
	tformid.UpdatedAt = time.Now()
	conn.Save(&tformid)
	return
}

func ValidateAccessToken(accessToken string) (flag bool) {
	if "dean" == accessToken {
		flag = true
	}
	flag = true
	return
}

func AddOpenWeixinAccessToken(wxlogin OpenWeixinAccessToken) (flag bool) {
	// aaa := conn.Model(&user).Where("openid = ?", user.Openid).Updates(user)
	// conn.Where("openid = ?", wxlogin.Openid).FirstOrCreate(&wxlogin)
	if "" != wxlogin.Openid {
		conn.Save(&wxlogin)
		flag = true
	}

	return
}

func GetFormByOpenid(openid string) (formid string) {
	var form TemplateFormID
	// conn.Debug().Where("open_id = ?", openid).Where("created_at > ?", "2018-11-20 23:40:46").First(&form)
	// db := conn.Debug().Where("openid = ?", openid).Where("created_at >= DATE_SUB(NOW(), INTERVAL 26 DAY)" ).First(&form)
	// 只有 7 天内的 formid 可以使用
	// db := conn.Debug().Where("openid = ?", openid).Where("created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)" ).First(&form)
	// status = 1 表示初始化
	db := conn.Debug().Not("formid", "the formId is a mock one").Where("status = 1 and openid = ?", openid).Where("created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)" ).First(&form)
	if db.Error != nil {
		beego.Error(db.Error)
	} else {
		formid = form.Formid
		beego.Info(form.CreatedAt, formid)
	}
	return
}

func UpdateFromidStatus(formid string) {
	var form TemplateFormID
	form.Formid = formid
	form.Status = 0
	conn.Debug().Model(&form).Where("formid = ?", formid).UpdateColumn("status", form.Status)
}


func AddWechatLoginScene (scene WechatLoginScene) (flag bool) {
        conn.Save(&scene)
	flag = true
        return
}
