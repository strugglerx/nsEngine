package utils

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson" // "github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	// // "reflect"
)

const EipDomain = "http://210.31.182.24"

//自定义Headers里的Cookie
func normalHeader(session string) requests.Header {
	normalHeader := requests.Header{
		"Host":       "eip.imnu.edu.cn",
		"Origin":     "http://eip.imnu.edu.cn",
		"Connection": "keep-alive",
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_3 like Mac OS X) AppleWebKit/603.3.8 (KHTML, like Gecko) Mobile/14G60 MicroMessenger/6.6.7 NetType/WIFI Language/zh_CN",
		"Referer":    "http://eip.imnu.edu.cn/EIP/weixin/weui/chengjichaxun.html",
		"Cookie":     session,
	}
	return normalHeader
}

//网费
func netlist(cookie *http.Cookie) string {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	//获取配置里的代理链接
	proxy:=beego.AppConfig.String("proxy::url")
	reqs.Proxy(proxy)
	headers := normalHeader("")
	req, _ := reqs.Post(EipDomain+"/EIP/edu/wangfei/queryUsrBindProduct.htm", headers)
	status := gjson.Get(req.Text(), "#").Bool()
	if status {
		return gjson.Get(req.Text(), "0.otherData").String()
	} else {
		return "-1"
	}

}

//校园卡余额 慢且经常不能用
func cardremain(cookie *http.Cookie) {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	reqs.Debug = 1
	reqs.Proxy("http://140.143.96.216:80")
	headers := normalHeader("")
	req, _ := reqs.Post(EipDomain+"/EIP/queryservice/query.htm?snumber=QRY_BAL&xh=20151105822", headers)
	fmt.Println(req.Text())
}

//校园卡消费列表
func cardlist(cookie *http.Cookie) string {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	headers := normalHeader("")
	req, _ := reqs.Post(EipDomain+"/EIP/edu/ykt_tongji.htm", headers)
	// fmt.Println(req.Text())
	return gjson.Get(req.Text(), "tongji.0.CDATE").String()
}

//校园卡消费明细
func carddetail(date string, cookie *http.Cookie) string {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	headers := normalHeader("")
	data := requests.Datas{
		"date": date,
	}
	req, _ := reqs.Post(EipDomain+"/EIP/edu/ykt_mingxi.htm", headers, data)
	// fmt.Println(req.Text())
	pop := fmt.Sprintf("%d", gjson.Get(req.Text(), "mingxi.#").Int()-1)
	// s := strconv.Itoa(i) int转string
	remain := gjson.Get(req.Text(), "mingxi."+string(pop)+".ACCOST").String()
	RemainJson := make(map[string]string)
	RemainJson["CARDBAL"] = remain
	r, _ := json.Marshal(RemainJson)
	return string(r)
}

//student information
func info(cookie *http.Cookie) string {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	headers := normalHeader("")
	req, _ := reqs.Post(EipDomain+"/EIP/edu/xueji.htm", headers)
	return req.Text()
}

//class
func class_(date string, cookie *http.Cookie) string {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	headers := normalHeader("")
	data := requests.Datas{
		"monday_": date,
	}
	req, _ := reqs.Post(EipDomain+"/EIP/qiandao/kebiao/queryKebiaoByUserId.htm", headers, data)
	if len(req.Text()) > 0 {
		return req.Text()
	}
	return "-1"

}

//scores
func score(cookie *http.Cookie) string {
	var reqs = requests.Requests()
	reqs.SetCookie(cookie)
	headers := normalHeader("")
	req, _ := reqs.Post(EipDomain+"/EIP/edu/chengji.htm", headers)
	// fmt.Println(req.Text())
	change := []byte(req.Text())
	//单个成绩数据结构
	type EachItem struct {
		JXJHH string
		XQ    string
		XH    string
		KCM   string
		XF    float64
		CJ    string
		XKLX  string
		_id   struct {
			inc        int64
			machine    int64
			new        bool
			time       int64
			timeSecond int64
		}
	}
	//原始api返回数据结构
	type EipStr struct {
		JXJHH string
		XQ    string
		CJ    []EachItem
	}
	//修改后的数据结构
	type CustomItem struct {
		Course     string      `json:"course"`
		Credit     interface{} `json:"credit"`
		Attributes string      `json:"attributes"`
		Score      string      `json:"score"`
	}
	type Custom struct {
		Term   string       `json:"term"`
		Values []CustomItem `json:"values"`
	}

	var eipformat []EipStr
	json.Unmarshal(change, &eipformat)
	// fmt.Printf("%+v", eipfor[0].CJ[0])
	var customformat []Custom
	for _, v := range eipformat {
		var tempitem []CustomItem
		for _, o := range v.CJ {
			item := CustomItem{o.KCM, o.XF, o.XKLX, o.CJ}
			tempitem = append(tempitem, item)
		}
		tempset := Custom{v.XQ, tempitem}
		customformat = append(customformat, tempset)
	}
	// fmt.Printf("%+v", customformat)

	customlen :=len(customformat) //获取长度
	recustom:=make([]Custom,customlen) //开辟空间
	for i,v:=range(customformat){
		//反序列
		recustom[customlen-i-1]=v
	}

	result, _ := json.Marshal(recustom)
	return string(result)
}

//登录
func login(user string) (bool, *http.Cookie) {
	// reqs.Debug = 1
	headers := requests.Header{
		"Host":                      "eip.imnu.edu.cn",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"Cookie":                    "EIPUserId=" + user,
		"User-Agent":                "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_3 like Mac OS X) AppleWebKit/603.3.8 (KHTML, like Gecko) Mobile/14G60 MicroMessenger/6.6.7 NetType/WIFI Language/zh_CN",
		"Accept-Language":           "zh-cn",
		"Accept-Encoding":           "gzip, deflate",
	}
	resp, err := requests.Get(EipDomain+"/EIP/weixinEnterprise/cookie.htm?url=/weixin/weui/jiugongge.htmlSPT623e8474576a41e5959ec47f0505109e", headers)
	cookie := &http.Cookie{}
	if err != nil ||resp.R.StatusCode!=200 {
		return false, cookie
	}
	//fmt.Printf("%+v", resp.R)
	head := resp.R.Header["Content-Type"][0]
	reg := regexp.MustCompile(`gbk`)
	if  reg.MatchString(head) {
		return false, resp.Cookies()[0]
	}
	return true, resp.Cookies()[0]
	// fmt.Println(resp.Cookies())
	//[]*http.Cookie
}

func EipEntry(user string, type_ string, date string) string {
	status, cookie := login(user)
	if status {
		switch type_ {
		case "score":
			return score(cookie)
		case "info":
			return info(cookie)
		case "card":
			date := cardlist(cookie)
			if len(date) == 8 {
				result := carddetail(date, cookie)
				return result
			} else {
				return "-1"
			}
		case "net":
			return netlist(cookie)
		case "class_":
			return class_(date, cookie)
		default:
			return "-1"
		}
	}
	return "-1"
}
