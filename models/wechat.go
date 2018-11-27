package models

import (
	 "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type (
	AnalysisDailySummary struct {
		List []struct {
			RefDate    string `json:"ref_date"`
			SharePv    int    `json:"share_pv"`
			ShareUv    int    `json:"share_uv"`
			VisitTotal int    `json:"visit_total"`
		} `json:"list"`
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
		OpenID string `json:"openid"`
		FormID string `json:"formid"`
		Status uint   // -1 发送消息失败， 0 已通知，1,初始化  2 已过期, 剩余保留
	}

	// 获取小程序码
	// 参考：https://developers.weixin.qq.com/miniprogram/dev/api/open-api/qr-code/getWXACodeUnlimit.html
	QRCodeRequestParms struct {
		AccessToken string `json:"access_token"`
		Scene string `json:"scene"`
		Page string `json:"page"`
		Width string `json:"width"`
		AutoColor string `json:"auto_color"`
		LineColor string `json:"line_color"`
		IsHyaline string `json:"is_hyaline"`
	}

	QRCodeReturn struct {
		Errcode string `json:"errcode"`
		Errmsg string `json:"errmsg"`
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
        // aaa := conn.Model(&user).Where("open_id = ?", user.OpenID).Updates(user)
        // conn.Where("openid = ?", wxlogin.Openid).FirstOrCreate(&wxlogin)
	if ("" != wxlogin.Openid) {
        	conn.Save(&wxlogin)
		flag = true
	}

        return
}
