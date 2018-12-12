package test

import (
	"github.com/bbhj/bbac/models"
	_ "github.com/bbhj/bbac/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	"github.com/esap/wechat"
	. "github.com/smartystreets/goconvey/convey"
)

var apiurlMap = map[string]string{
	"/lastest/api/robot/qq": "GET",
	"/lastest/webapi/robot/qq": "GET",
	"/lastest/webapi/user/get_drag_list": "GET",
	"/lastest/webapi/user/get_info": "GET",
	"/lastest/webapi/user/login": "POST",
	"/lastest/webapi/user/logout": "POST",
}

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../"+string(filepath.Separator))))

	models.Connect()
	// defer models.Close()

	wechat.Set(beego.AppConfig.String("wechat_token"), beego.AppConfig.String("wechat_appid"), beego.AppConfig.String("wechat_secret"), beego.AppConfig.String("wechat_aeskey"))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	apiurl := "/lastest/api/robot/wechat"
	r, _ := http.NewRequest("GET", apiurl, nil)
	r.Header.Add("Cookie", "token=test")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestBeego is a sample to run an endpoint test
func TestBeegoUser(t *testing.T) {
	for apiurl, method := range apiurlMap {
		// apiurl := "/lastest/api/robot/qq"
		r, _ := http.NewRequest(method, apiurl, nil)
		r.Header.Add("Cookie", "token=test")
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
		beego.Info("request method:", method, "apiurl:", apiurl)

		Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	}
}
