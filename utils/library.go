/*
 * @Description:
 * @Author: Moqi
 * @Date: 2018-12-12 10:34:46
 * @Email: str@li.cm
 * @Github: https://github.com/strugglerx
 * @LastEditors: Moqi
 * @LastEditTime: 2019-01-28 19:11:01
 */

package utils

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asmcos/requests"
)

const eipDomain = "http://libwx.imnu.edu.cn"

var headers = requests.Header{
	"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Safari/537.36",
	"Host":       "libwx.imnu.edu.cn",
}

func EipLibrary(name, page string) string {
	p := requests.Params{
		"q":    name,
		"t":    "any",
		"page": page,
	}
	resp, _ := requests.Get(eipDomain+"/m/weixin/wsearch.action", p, headers)
	doc, err := goquery.NewDocumentFromResponse(resp.R)
	if err != nil || resp.R.StatusCode != 200 {
		return "-1"
	}
	var array []BookItem
	doc.Find(".weui_panel_bd a").Each(func(i int, selection *goquery.Selection) {
		var item BookItem
		url, _ := selection.Attr("href")
		item.Id = url[int(strings.LastIndex(url, "id="))+3:]
		img, _ := selection.Find("img").Attr("src")
		if ok, _ := regexp.Match("http", []byte(img)); !ok {
			img = ""
		}
		item.Img = img
		item.Name = selection.Find("h4").Text()
		item.Tag = selection.Find("p").Text()
		item.Status = selection.Find("li").Text()
		array = append(array, item)
	})
	result, _ := json.Marshal(array)
	return string(result)
}

func EipLibraryDetail(id string) string {
	p := requests.Params{
		"id": id,
	}
	resp, _ := requests.Get(eipDomain+"/m/weixin/wdetail.action", p, headers)
	doc, err := goquery.NewDocumentFromResponse(resp.R)
	if err != nil || resp.R.StatusCode != 200 {
		return "-1"
	}

	var array []ItemDetail
	doc.Find(`.weui_panel_bd:nth-child(2) div[class="weui_media_bd"]`).
		Each(func(i int, selection *goquery.Selection) {
			var item ItemDetail
			item.Name = selection.Find("h4").Text()
			item.Tag = selection.Find("p").Text()
			item.Status = selection.Find("li").Text()
			array = append(array, item)
		})

	result, _ := json.Marshal(array)

	return string(result)

}
