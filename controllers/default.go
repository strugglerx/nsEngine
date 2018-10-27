package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
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
	sess :=c.GetSession("role")
	if sess =="admin"{
		c.Redirect("/admin",302)
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) LoginGet() {
	sess :=c.GetSession("role")
	if sess =="admin"{
		c.Redirect("/admin",302)
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) Login() {
	sess :=c.GetSession("role")
	if sess ==nil{
		c.SetSession("role","admin")
		c.Data["Website"] = "second access"
		c.Data["Email"] = "second@gmail.com"
	}else{
		c.Data["Website"] = sess
		c.Data["Email"] = sess
	}
	c.TplName = "index.html"
}

func (c *MainController) Logout() {
	c.DelSession("role")
	info := MainResponse{"1000", 0,"logout success"}
	c.Ctx.WriteString(info.JsonParse())

}
