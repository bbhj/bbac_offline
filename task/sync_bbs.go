package task

import (
	"github.com/astaxie/beego"
	"github.com/bbhj/bbac/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type (
	PreForumPost struct {
		gorm.Model
		Message string `json:"message"`
		Subject string `json:"subject"`
		Useip   string `json:"useip"`
		Pid     int64  `json:"pid"`
		Tid     int64  `json:"tid"`
	}
)

// 去掉html中所有标签
func trimHtml(src string) string {

    //将HTML标签全转换成小写
    re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
    src = re.ReplaceAllStringFunc(src, strings.ToLower)

    // //去除STYLE
    re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
    src = re.ReplaceAllString(src, "")

    //去除SCRIPT
    re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
    src = re.ReplaceAllString(src, "")

    // [attach]
    // src = strings.Replace(src, "[attach]", "[attach]图片ID：", -1)
    src = strings.Replace(src, "&nbsp;", "", -1)

    // 去掉[]标签, 如 [color=#000000]
    re, _ = regexp.Compile("\\[[\\S\\s]+?\\]")
    src = re.ReplaceAllString(src, "")

    //去除所有尖括号内的HTML代码，并换成换行符
    re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
    // src = re.ReplaceAllString(src, "\n")
    src = re.ReplaceAllString(src, "")


    // //去除连续的换行符
    re, _ = regexp.Compile("\\s{4,}")
    src = re.ReplaceAllString(src, "\n")

    return strings.TrimSpace(src)
}

func formatTime(tstr string) (tm time.Time, err error) {
	// timeFormats := []string{"2006-01-02 15:04:05", "2006-01-0215:04:05", "2006-01", "2006-1-02", "2006-01-2", "2006-1-2", "2006-01-02", "2006-0102", "2006年1月2日", "2006年1月02日", "2006年01月2日", "2006年01月02日", "2006年01月**日", "2006年1月", "2006年"}
	timeFormats := []string{"2006-01-02 15:04:05", "2006-01-02150405", "2006-01-0215:04:05", "2006-01", "2006-1-02", "2006-01-2", "2006-1-2", "2006-01-02", "2006-0102", "2006--01-02", "2006年1月2日", "2006年1月02日", "2006年01月2日", "2006年01月02日", "2006年01月**日", "2006年1月", "2006年", "2006年01月02", "2006年01月02月", "2006月01月02日", "2006年01年02月", "2006年01年02日", "2006年1月02", "2006"}
	for _, timeFormat := range timeFormats {
		tm, err = time.ParseInLocation(timeFormat, tstr, time.Local)
		if err == nil {
			break
		}
	}
	return
}

func parseHtml(datafrom, title, msg string) (article models.Article) {
	reg := regexp.MustCompile(`\r`)
	msg = reg.ReplaceAllString(msg, "\n")

	// 剔除msg中无用的头尾
	// 登记信息/d ---.*-- /d站务电话
	reg = regexp.MustCompile(`(.*\n)*.*登记信息.*\n`)
	msg = reg.ReplaceAllString(msg, "")

	// reg = regexp.MustCompile(`站务电话(.*\n)*.*`)
	reg = regexp.MustCompile(`站务电话(.*\n)*.*|宝贝回家志愿者(.*\n)*.*`)
	//reg = regexp.MustCompile(`站务电话(.*\n)*.*`)
	msg = reg.ReplaceAllString(msg, "")


	msg = strings.Trim(msg, "\n")

	beego.Info("=======", datafrom, "==========")
	beego.Info("=======", title, "==========")

	// var article models.Article
	article.DataFrom = datafrom
	article.Title = title

	var err error
	var detailFlag bool

	infoList := strings.Split(msg, "\r")
	//infoList := strings.Split(msg, "\n")
	if (len(infoList) <= 1) {
		beego.Error("=======fail", datafrom, "==========")
		return
	}
	for _, info := range infoList {
		if ("" == info) {
			continue
		}

		exp := regexp.MustCompile(`：.*|:.*`)
        	valueList := exp.FindAllString(info, -1)
		value := ""
		if (len(valueList) == 1) {
			value = valueList[0]
		} else {
			exp := regexp.MustCompile(`:.*`)
        		valueList = exp.FindAllString(info, -1)
			if (len(valueList) == 1) {
				value = valueList[0]
			}
		}
		key := strings.Replace(info, value, "", -1)

		// 格式化key, value
 		// exp = regexp.MustCompile(`(\[.+?\])|(\^M)|：|:|(\s)|(\r)`)
 		exp = regexp.MustCompile(`(\[.+?\])|(\^M)|：|:|(\s)`)
		key = exp.ReplaceAllString(key, "")
		value = exp.ReplaceAllString(value, "")
		// beego.Error(key)

		// // 删除key中的非汉字和空格
		// exp = regexp.MustCompile(`[^\p{Han}]| `)
		// key = exp.ReplaceAllString(key, "")

		tmexp := regexp.MustCompile(`[（|(].*[）|)]|农历|阳历|不准确|大概|日期不确定|不确定|约|X日|份左右|期不详。|。|具体.+?|失踪|宋振彪|宋振邦2008年|.*身份证日期|左右|号|阴历|某天|夏.*|年底|腊月.*|九.*|八.*|天已记不清楚|冬月.*|正.*|初.*|下午1点半|大|生`)
		switch key {
			case "寻亲类别", "类别" :
				article.Category = value
			case "宝贝回家编号" :
				article.Babyid, err = strconv.ParseInt(value, 10, 64)
				if (err != nil) {
					beego.Error("宝贝回家编号", value, err)
				}
			case "姓名" :
				article.Nickname = value
			case "性别" :
				// /* 值为1时是男性，值为2时是女性，值为0时是未知 */
				if "女" == value {
					// 1 --> 2
					article.Gender = 2
				} else if "男" == value {
					// 0 -> 1
					article.Gender = 1
				}
			case "失踪时身高" :
				article.Height = value
			case "失踪人所在省", "籍贯" :
				article.Province = value
			case "失踪地点", "失踪地址", "地址" :
				article.Address = value
			case "失踪者特征描述" :
				article.Characters = value
			case "失踪人户籍所在地" :
				article.BirthedProvince = value
			case "采血情况" :
			case "出生日期" :
				value =	tmexp.ReplaceAllString(value, "")
				value = strings.Replace(value, "—", "", -1)
				article.BirthedAt, err = formatTime(value)
				if err != nil {
					beego.Error("出生日期", value, err)
				}
				// beego.Info(key, value, ", 格式化:", article.BirthedAt)
			case "失踪日期", "失踪时间" :
				value =	tmexp.ReplaceAllString(value, "")
				value = strings.Replace(value, "—", "", -1)
				article.MissedAt, err = formatTime(value)
				if err != nil {
					beego.Error("失踪日期", value, err)
				}
				// beego.Info(key, value, ", 格式化:", article.MissedAt)
			case "注册时间" :
				
			case "其他资料", "其他情况" :
				article.Details = value
				detailFlag = true
			default :
				if detailFlag {
					article.Details += value
				}
		}
	}
	beego.Error(msg)
	beego.Info("Babyid:", article.Babyid, ", 数据来源:", article.Details)
	return
}

func syncFrombbs() {
	// preForumPosts := models.GetLastBBSInfo()
	preForumPosts := models.GetAllBBSInfo()

	for _, preForumPost := range preForumPosts {
		datafrom := "https://bbs.baobeihuijia.com/thread-"
		datafrom += strconv.FormatInt(int64(preForumPost.Tid), 10) + "-1-1.html"
		// beego.Info("====", datafrom)

		// beego.Error(preForumPost.Tid, preForumPost.Pid)
		msg := trimHtml(preForumPost.Message)
		article := parseHtml(datafrom, preForumPost.Subject, msg)
		if article.Babyid == 0 {
			beego.Critical("this babyid is null.", article.DataFrom)
			continue
		}

		if beego.BConfig.RunMode == "prod" {
		article.UUID = uuid.New().String()
		models.AddArticle(article)
		models.SyncPictureFromBbs(preForumPost.Tid, preForumPost.Pid, article.Babyid, article.UUID)
		}
		return
	}
}

