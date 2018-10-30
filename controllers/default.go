package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"server/models"
	"server/utils"
)

type MainController struct {
	beego.Controller
}

type MainResponse struct {
	Code   string `json:"code"`
	Status int    `json:"status"`
	Msg interface{}    `json:"message,omitempty"`
}

func (res *MainResponse) JsonParse() string {
	datas, _ := json.Marshal(res)
	return string(datas)
}

func (c *MainController) Get() {
	c.TplName = "home.html"
}


func (c *MainController) Login() {
	sess :=c.GetSession("role")
	if sess ==nil{
		c.TplName = "login.html"
	}else{
		c.Redirect("/manager",302)
	}

}

func (c *MainController) LoginPost() {
	sess :=c.GetSession("role")
	if sess ==nil{
		user:=c.GetString("user")
		pwd:=c.GetString("passwd")

		if user!=""&&pwd!=""{
			//加密密码
			cryptoPwd:=utils.Md5String(pwd)
			verifyInfo:=models.UserVerify(user,cryptoPwd)
			if len(verifyInfo)==1 {
				c.SetSession("role",verifyInfo[0].Role)
				c.SetSession("user",verifyInfo[0].User)
				info := MainResponse{"1000", 0,"login success"}
				c.Ctx.WriteString(info.JsonParse())
			}else {
				info := MainResponse{"1002", -2,"login fail"}
				c.Ctx.WriteString(info.JsonParse())
			}
		}else {
			info := MainResponse{"1002", -2,"params can not empty!"}
			c.Ctx.WriteString(info.JsonParse())

		}

	}else{
		info := MainResponse{"1001", -1,"you already login"}
		c.Ctx.WriteString(info.JsonParse())
	}
}

func (c *MainController) Logout() {
	c.DelSession("role")
	info := MainResponse{"1000", 0,"logout success"}
	c.Ctx.WriteString(info.JsonParse())

}
