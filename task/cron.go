package task

import (
	"github.com/astaxie/beego/logs"
	"github.com/bbhj/bbac/models"
	"github.com/astaxie/beego/toolbox"
	"github.com/imroc/req"
	"github.com/esap/wechat"
	"github.com/astaxie/beego"
	"plugin"
)

func CronTask() {

	timeStr := "0/5 * * * * *" //每隔3秒执行

	t1 := toolbox.NewTask("timeTask", timeStr, func() error {

		// todo do what you want
		if beego.BConfig.RunMode == "dev" {
			syncFrombbs()
		}
		// msg := models.Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
		/// logs.Info("this is crontab task....")
		// weixinData()
		// broadcast <- msg

		loadPlugins()
		return nil
	})

	toolbox.AddTask("tk1", t1)
	toolbox.StartTask()
}

func weixinData() {
	// 20180722
        param := req.Param{
                "begin_date": "20180722",
                "end_date":  "20180722",
        }
	a1, _ := req.Post(beego.AppConfig.String("wehcat_getAnalysisDailySummary") + wechat.GetAccessToken(), req.BodyJSON(param))
	var analysisDaily models.AnalysisDailySummary
	a1.ToJSON(&analysisDaily)
	logs.Info(analysisDaily)
        logs.Info("累计用户数: %d, 转发人数: %d, 转发次数: %d\n", analysisDaily.List[0].VisitTotal, analysisDaily.List[0].ShareUv, analysisDaily.List[0].SharePv)
}

func WechatSendTemplateMessage() {
        var tm models.TemplateMessage
        tm.Touser  = "oPPbr0M2h0geV-jgzUPve9g3x3jg"
        tm.TemplateID = "nYL5mnIUq5aybu8uZWaF1qwWPp_6Up4EhhUmrbHXWmc"
        tm.Page = "/pages/home/main"
        tm.FormID = "2fc6e039e32e2c77035d7c7311920ab1"
        tm.EmphasisKeyword = ""
        tm.Data.Keyword1.Value = "帮忙寻找xxx"
        tm.Data.Keyword2.Value = "阿正_Dean"
        tm.Data.Keyword3.Value = "常州市新北区河海东路200号8栋2单元301"
        tm.Data.Keyword4.Value = "这是备注"
        tm.Data.Keyword5.Value = "已有500人参与提供帮助"

        a2, _ := req.Post(beego.AppConfig.String("wechatApiSendTemplateMessage") + wechat.GetAccessToken(), req.BodyJSON(tm))
        logs.Debug( "send template message: ", a2 )
}


func loadPlugins() {
	pluginName := "plugins/" + "print.so"
	//打开动态库
	pdll, err := plugin.Open(pluginName)
	if err != nil {
		//...
		return
	}
	//获取动态库方法
	funcPrint, err := pdll.Lookup("PrintTest")
	if err != nil {
		//...
		return
	}

	//动态库方法调用
	funcPrint.(func(string))("hello go plugin")
	return
}

