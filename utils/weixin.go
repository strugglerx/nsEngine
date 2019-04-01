/*
 * @Description:
 * @Author: Moqi
 * @Date: 2018-12-12 10:35:14
 * @Email: str@li.cm
 * @Github: https://github.com/strugglerx
 * @LastEditors: Moqi
 * @LastEditTime: 2019-03-09 11:48:14
 */

package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"server/models"
	"time"

	"github.com/asmcos/requests"
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
)

type WxUser struct {
	Session string `json:"session,omitempty"`
	Openid  string `json:"openid,omitempty"`
	Appid   string `json:"appid,omitempty"`
	Status  bool   `json:"status,omitempty"`
}

type wxSessResponse struct {
	Openid      string
	Session_key string
	Unionid     string
	Errcode     int
	ErrMsg      string
}

type WxTemplate struct {
	Touser           string      `json:"touser,omitempty"`
	Template_id      string      `json:"template_id,omitempty"`
	Page             string      `json:"page,omitempty"`
	Form_id          string      `json:"form_id,omitempty"`
	Data             interface{} `json:"data,omitempty"`
	Emphasis_keyword string      `json:"emphasis_keyword,omitempty"`
}

type WxToken struct {
	Token      string
	UpdateTime int
}

var WxToken_ *WxToken = &WxToken{"", 0}

func WxSession(code string) WxUser {
	appid := beego.AppConfig.String("appid")
	secret := beego.AppConfig.String("secret")
	p := requests.Params{
		"appid":      appid,
		"secret":     secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	resp, _ := requests.Get("https://api.weixin.qq.com/sns/jscode2session", p)
	var data wxSessResponse
	change := []byte(resp.Text())
	json.Unmarshal(change, &data)
	wxres := data
	if wxres.Errcode == 0 {
		result := WxUser{wxres.Session_key, wxres.Openid, appid, true}
		return result
	}
	return WxUser{"", "", "", false}

}

func WxGetAccessToken() *WxToken {
	appid := beego.AppConfig.String("appid")
	secret := beego.AppConfig.String("secret")
	p := requests.Params{
		"appid":      appid,
		"secret":     secret,
		"grant_type": "client_credential",
	}

	now := int(time.Now().Unix())

	if now-WxToken_.UpdateTime >= 7200 {
		resp, _ := requests.Get("https://api.weixin.qq.com/cgi-bin/token", p)
		WxToken_.Token = gjson.Get(resp.Text(), "access_token").String()
		WxToken_.UpdateTime = now
		return WxToken_
	} else {
		return WxToken_
	}

}

func WxSendMsg(content, openid, accessToken string) int64 {
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + accessToken

	type jsonBody struct {
		Touser  string `json:"touser"`
		Msgtype string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}
	var body jsonBody
	body.Touser = openid
	body.Msgtype = "text"
	body.Text.Content = content
	//解决json marshal 转义问题
	bb := WxJsonMarshal(body)
	//封装json post请求
	jres := JsonPost(url, string(bb))
	return gjson.Get(jres, "errcode").Int()
}

func WxJsonMarshal(t interface{}) []byte {
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(t)
	return buffer.Bytes()
}

// 示例数据
/* 	{
	"touser":"oYA_z0DMu9nlFQQNPlzWMZqpL-SY",
	"template_id":"8QsuHQf-Nd6mmHOmFIhP_aFGvuWkjXN56eYRE7fzS7s",
	"page":"/pages/index/index",
	"data":{
		"keyword1":{
			"value":"祁强强"
		},
		"keyword2":{
			"value":"2015110xxxx"
		},
		"keyword3":{
			"value":"明天的课程更新啦"
		}
	},
	"emphasis_keyword":"keyword1.DATA"
} */

func WxPushTemplate(body []byte) (interface{}, error) {
	var result WxTemplate
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("unmarshal fail")
	}
	result.Form_id = models.FormLastId(result.Touser)
	//清除formId
	if result.Form_id == "" {
		return nil, errors.New("formid is empty")
	}
	models.FormPop(result.Touser)
	//获取token
	token := WxGetAccessToken().Token

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", token)
	raw, _ := json.Marshal(result)
	text := JsonPost(url, string(raw))
	var res interface{}
	json.Unmarshal([]byte(text), &res)
	return res, nil
}
