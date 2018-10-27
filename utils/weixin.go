package utils

import (
	"encoding/json"
	"github.com/asmcos/requests"
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
)

type WxUser struct {
	Session string `json:"session,omitempty"`
	Openid string  `json:"openid,omitempty"`
	Appid string   `json:"appid,omitempty"`
	Status bool    `json:"status,omitempty"`
}

type wxSessResponse struct {
	Openid string
	Session_key string
	Unionid string
	Errcode int
	ErrMsg string
}

func WxSession(code string) WxUser {
	appid:=beego.AppConfig.String("appid")
	secret:=beego.AppConfig.String("secret")
	p :=requests.Params{
		"appid": appid,
		"secret": secret,
		"js_code": code,
		"grant_type": "authorization_code",
	}
	resp,_:=requests.Get("https://api.weixin.qq.com/sns/jscode2session",p)
	var data wxSessResponse
	change := []byte(resp.Text())
	json.Unmarshal(change,&data)
	wxres :=data
	if wxres.Errcode==0{
		result :=WxUser{wxres.Session_key,wxres.Openid,appid,true}
		return result
	}
	return WxUser{"","","",false}

}

func WxGetAccessToken() string  {
	appid:=beego.AppConfig.String("appid")
	secret:=beego.AppConfig.String("secret")
	p :=requests.Params{
		"appid": appid,
		"secret": secret,
		"grant_type": "client_credential",
	}
	resp,_:=requests.Get("https://api.weixin.qq.com/cgi-bin/token",p)
	return gjson.Get(resp.Text(),"access_token").String()
}

func WxSendMsg(content,openid,accessToken string) int64{
	url:="https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="+accessToken

	type jsonBody struct {
		Touser string `json:"touser"`
		Msgtype string `json:"msgtype"`
		Text struct{
			Content string `json:"content"`
		} `json:"text"`
	}
	var body jsonBody
	body.Touser=openid
	body.Msgtype="text"
	body.Text.Content=content
	bb,_:=json.Marshal(body)
	//封装json请求
	jres:=JsonPost(url,string(bb))
	return gjson.Get(jres,"errcode").Int()
}
