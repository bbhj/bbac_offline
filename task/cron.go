package task

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"github.com/bbhj/bbac/models"
	"github.com/esap/wechat"
	"github.com/imroc/req"
	"plugin"
)

func CronTask() {
	// syncFrombbs()
	// SendWechatTemplateMessage()
	// return

	timeStr := "0/5 * * * * *" //每隔3秒执行

	t1 := toolbox.NewTask("timeTask", timeStr, func() error {

		// todo do what you want
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
		"end_date":   "20180722",
	}
	a1, _ := req.Post(beego.AppConfig.String("wehcat_getAnalysisDailySummary")+wechat.GetAccessToken(), req.BodyJSON(param))
	var analysisDaily models.AnalysisDailySummary
	a1.ToJSON(&analysisDaily)
	logs.Info(analysisDaily)
	logs.Info("累计用户数: %d, 转发人数: %d, 转发次数: %d\n", analysisDaily.List[0].VisitTotal, analysisDaily.List[0].ShareUv, analysisDaily.List[0].SharePv)
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
