package utils

import (
	"github.com/astaxie/beego"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asmcos/requests" // "fmt"
	// "reflect"
)

func SportCurl(user string, pwd string) interface{} {
	headers := requests.Header{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	}
	// get参数
	/* 	params := requests.Params{
		"page":"2",
	} */
	data := requests.Datas{
		"txtuser":           user,
		"txtpwd":            pwd,
		"__EVENTTARGET":     "",
		"__EVENTARGUMENT":   "",
		"__LASTFOCUS":       "",
		"__VIEWSTATE":       "/wEPDwUKLTM4NjY5Mzc1Ng9kFgJmD2QWCmYPEGRkFgFmZAICDw8WAh4HVmlzaWJsZWhkZAIDDw8WAh4EVGV4dAXfDjxzcGFuIHN0eWxlPSJGT05ULUZBTUlMWTog5a6L5L2TOyBCQUNLR1JPVU5EOiB3aGl0ZTsgQ09MT1I6IHJlZDsgRk9OVC1TSVpFOiAxMy41cHQ7IG1zby1oYW5zaS1mb250LWZhbWlseTogQXJpYWw7IG1zby1iaWRpLWZvbnQtZmFtaWx5OiBBcmlhbDsgbXNvLWFzY2lpLWZvbnQtZmFtaWx5OiBBcmlhbCI+PHNwYW4gc3R5bGU9IkZPTlQtRkFNSUxZOiDlrovkvZM7IEZPTlQtU0laRTogMTJwdCIgbGFuZz0iRU4tVVMiPjxvOnA+DQo8cCBhbGlnbj0iY2VudGVyIj48c3Ryb25nPjxmb250IGNvbG9yPSIjMDAwMDAwIiBzaXplPSI1Ij48Zm9udCBzdHlsZT0iQkFDS0dST1VORC1DT0xPUjogI2UwZGI5ZSIgZmFjZT0iQXJpYWwiPjxmb250IHN0eWxlPSJCQUNLR1JPVU5ELUNPTE9SOiAjZTBkYjllIiBmYWNlPSJBcmlhbCI+PHNwYW4gc3R5bGU9IkZPTlQtRkFNSUxZOiDlrovkvZM7IEJBQ0tHUk9VTkQ6IHdoaXRlOyBDT0xPUjogcmVkOyBGT05ULVNJWkU6IDEzLjVwdDsgbXNvLWhhbnNpLWZvbnQtZmFtaWx5OiBBcmlhbDsgbXNvLWJpZGktZm9udC1mYW1pbHk6IEFyaWFsOyBtc28tYXNjaWktZm9udC1mYW1pbHk6IEFyaWFsIj48c3BhbiBzdHlsZT0iRk9OVC1GQU1JTFk6IOWui+S9kzsgRk9OVC1TSVpFOiAxMnB0IiBsYW5nPSJFTi1VUyI+PG86cD4mbmJzcDsgPC9vOnA+PC9zcGFuPjwvc3Bhbj48L2ZvbnQ+PC9mb250PjwvZm9udD48L3N0cm9uZz48L3A+DQo8cCBhbGlnbj0iY2VudGVyIj48c3Ryb25nPjxmb250IGNvbG9yPSIjMDAwMDAwIiBzaXplPSI1Ij7pgJrnn6U8L2ZvbnQ+PC9zdHJvbmc+PC9wPg0KPHAgYWxpZ249ImNlbnRlciI+PGZvbnQgY29sb3I9IiMwMDAwMDAiPiZuYnNwOzwvZm9udD48L3A+DQo8cD48Zm9udCBjb2xvcj0iIzAwMDAwMCIgc2l6ZT0iNCI+PC9mb250PjwvcD4NCjxwPjxmb250IGNvbG9yPSIjMDAwMDAwIiBzaXplPSI0Ij48L2ZvbnQ+PC9wPg0KPHA+PGZvbnQgc2l6ZT0iNCI+PGZvbnQgY29sb3I9IiMwMDAwMDAiPjEuMjAxNS0yMDE25a2m5bm056ysMuWtpuacn+S9k+iCsuivvuaIkOe7qeW3suabtOaWsO+8jOi/m+WFpeafpeivouezu+e7n+WQjuivt+eCueWHuyZsZHF1bzvljoblj7LkvZPogrLmiJDnu6nkv6Hmga8mcmRxdW876L+b6KGM5p+l6K+i77yBPC9mb250PiZuYnNwOzwvZm9udD48L3A+DQo8cD48Zm9udCBzaXplPSI0Ij4yLuWtpueUn+eZu+W9lei0puWPt+S4uuWtpuWPt++8jOWvhueggeS4uueUn+aXpeWFq+S9jeOAgjwvZm9udD48L3A+DQo8cD48Zm9udCBjb2xvcj0iIzAwMDAwMCIgc2l6ZT0iNCI+PC9mb250PjwvcD4NCjxwPjxmb250IGNvbG9yPSIjMDAwMDAwIiBzaXplPSI0Ij48L2ZvbnQ+PC9wPg0KPHA+PGZvbnQgY29sb3I9IiMwMDAwMDAiIHNpemU9IjQiPjwvZm9udD48L3A+DQo8cD48Zm9udCBjb2xvcj0iIzAwMDAwMCIgc2l6ZT0iNCI+Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7IOWFseS9k+mDqDwvZm9udD48L3A+DQo8cD48Zm9udCBjb2xvcj0iIzAwMDAwMCIgc2l6ZT0iNCI+Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7Jm5ic3A7MjAxNuW5tDnmnIgxMuaXpTwvZm9udD48L3A+DQo8L286cD48L3NwYW4+PC9zcGFuPmRkAgQPDxYCHwEFEjIwMTYtMi0yMyAxODowMDowMGRkAgUPDxYCHwEFEjIwMjAtMy0yNiAyMzowNDo0MGRkGAEFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYBBQVidG5va5XoZBKrg3o8ecM6KR5Hk5FVdEvu",
		"__EVENTVALIDATION": "/wEWCALdr4DxCQKBwaG/AQLMrvvQDQLd8tGoBALWwdnoCALB2tiCDgKd+7q4BwL9kpmqCiqkts8uGgg92XCVJce3KG7zZAef",
		"dlljs":             "st",
		"btnok.x":           "4",
		"btnok.y":           "12",
	}
	sportDomain:=beego.AppConfig.String("sport::url")

	resp, err := requests.Post(sportDomain, headers, data)
	if err!=nil{
		return "-1"
	}
	doc, err := goquery.NewDocumentFromResponse(resp.R)
	if err != nil ||resp.R.StatusCode!=200 {
		return "-1"
	}
	//数据初始化
	type AllData struct {
		Name         string      `json:"name"`
		ClassName    string      `json:"className"`
		ClassDetail  interface{} `json:"classDetail"`
		SportsDetail interface{} `json:"sportsDetail"`
	}
	//提前扩展十个长度
	citList:=make([]map[string]string,10) //体育班信息
	var scoreList []string    //体育班成绩过度列表
	var list []map[string]string    //体测信息

	Result := new(AllData)
	//名字
	stuName := doc.Find("#lblname").Text()
	reg := regexp.MustCompile(`[\s]`)
	res := reg.ReplaceAllString(stuName, "")
	//判断是否登录成功
	if res == "学号" {
		return "-1"
	}
	Result.Name = res
	//体育班
	doc.Find("#pAll tr:nth-child(9) table tr td").Each(func(i int, item *goquery.Selection) {
		reg := regexp.MustCompile(`体育班名称：|[\s]| `)
		res := reg.ReplaceAllString(item.Text(), "")
		if i == 23 || i == 17 || i == 22 {
			return
		} else if i == 0 {
			Result.ClassName = strings.TrimSpace(res)
			return
		}
		scoreList = append(scoreList, res)
	})
	//初始化
	num :=0
	for i, single := range scoreList {
		citem := map[string]string{}

		if i%2 == 0 {
			citem["item"] = single
			citem["value"] = scoreList[i+1]
			citList[num] =citem
			num++
			continue
		}
	}
	Result.ClassDetail = citList
	//体测
	doc.Find("#pAll tr:nth-child(14) table tr").Each(func(i int, item *goquery.Selection) {
		key := [5]string{"name", "score", "single", "assess"}
		temp := map[string]string{}
		item.Find("td").Each(func(index int, item *goquery.Selection) {
			// fmt.Println(reflect.TypeOf(index))
			if index > 0 && index < 5 {
				pat := `[\s]+`
				reg, _ := regexp.Compile(pat)
				kname := key[index-1]
				temp[kname] = reg.ReplaceAllString(item.Text(), "")
			}
		})
		list = append(list, temp)
	})
	Result.SportsDetail = list
	// datas type []bytes
	//datas, _ := json.Marshal(Result)
	// fmt.Println(resp.Cookies())
	//return string(datas)
	//fmt.Printf("%+t",Result)
	return Result
}
