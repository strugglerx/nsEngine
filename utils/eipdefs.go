package utils

import "github.com/asmcos/requests"

const EipDomain = "http://210.31.182.24"

var DefaultHeader = requests.Header{
	"Host":       "eip.imnu.edu.cn",
	"Origin":     "http://eip.imnu.edu.cn",
	"Connection": "keep-alive",
	"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_3 like Mac OS X) AppleWebKit/603.3.8 (KHTML, like Gecko) Mobile/14G60 MicroMessenger/6.6.7 NetType/WIFI Language/zh_CN",
	"Referer":    "http://eip.imnu.edu.cn/EIP/weixin/weui/chengjichaxun.html",
}

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

//图书馆

type BookItem struct {
	Id     string `json:"id"`
	Img    string `json:"img"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	Status string `json:"status"`
}

type ItemDetail struct {
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	Status string `json:"status"`
}
