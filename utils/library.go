package utils

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/asmcos/requests"
	"github.com/gin-gonic/gin/json"
	"regexp"
	"strings"
)

const eipDomain  = "http://libwx.imnu.edu.cn"
var headers = requests.Header{
"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Safari/537.36",
"Host": "libwx.imnu.edu.cn",
}

func EipLibrary(name,page string) string {
	p:=requests.Params{
		"q":name,
		"t":"any",
		"page":page,
	}
	resp,_ :=requests.Get(eipDomain+"/m/weixin/wsearch.action",p,headers)
	doc,err:=goquery.NewDocumentFromResponse(resp.R)
	if err != nil ||resp.R.StatusCode!=200 {
		return "-1"
	}
	type bookItem struct {
		Id string `json:"id"`
		Img string `json:"img"`
		Name string `json:"name"`
		Tag string `json:"tag"`
		Status string `json:"status"`
	}
	var array []bookItem
	doc.Find(".weui_panel_bd a").Each(func(i int, selection *goquery.Selection) {
		var item bookItem
		url,_ :=selection.Attr("href")
		item.Id= url[int(strings.LastIndex(url,"id="))+3:]
		img,_:=selection.Find("img").Attr("src")
		if ok,_:=regexp.Match("http",[]byte(img));!ok{
			img=""
		}
		item.Img=img
		item.Name=selection.Find("h4").Text()
		item.Tag=selection.Find("p").Text()
		item.Status=selection.Find("li").Text()
		array = append(array,item)
	})
	result,_ :=json.Marshal(array)
	return string(result)
}

func EipLibraryDetail(id string) string {
	p:=requests.Params{
		"id":id,
	}
	resp,_:=requests.Get(eipDomain+"/m/weixin/wdetail.action",p,headers)
	doc,err:=goquery.NewDocumentFromResponse(resp.R)
	if err != nil ||resp.R.StatusCode!=200 {
		return "-1"
	}
	type itemDetail struct {
		Name string `json:"name"`
		Tag string `json:"tag"`
		Status string `json:"status"`
	}
	var array []itemDetail
	doc.Find(`.weui_panel_bd:nth-child(2) div[class="weui_media_bd"]`).
		Each(func(i int, selection *goquery.Selection) {
			var item itemDetail
			item.Name=selection.Find("h4").Text()
			item.Tag=selection.Find("p").Text()
			item.Status=selection.Find("li").Text()
			array = append(array,item)
	})

	result,_:=json.Marshal(array)

	return  string(result)



}
