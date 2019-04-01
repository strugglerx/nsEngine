package controllers

import (
	"server/models"
	"server/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.TplName = "home.html"
}

func (c *MainController) Login() {
	sess := c.GetSession("role")
	if sess == nil {
		c.TplName = "login.html"
	} else {
		c.Redirect("/manager", 302)
	}

}

func (c *MainController) LoginPost() {
	sess := c.GetSession("role")
	if sess == nil {
		user := c.GetString("user")
		pwd := c.GetString("passwd")

		if user != "" && pwd != "" {
			//加密密码
			cryptoPwd := utils.Md5String(pwd)
			verifyInfo := models.UserVerify(user, cryptoPwd)
			if len(verifyInfo) == 1 {
				c.SetSession("role", verifyInfo[0].Role)
				c.SetSession("user", verifyInfo[0].User)
				c.Ctx.WriteString(MsgResponse("login success"))
			} else {
				c.Ctx.WriteString(MsgFail("login fail"))
			}
		} else {
			c.Ctx.WriteString(MsgFail("params can not empty!"))

		}

	} else {
		c.Ctx.WriteString(MsgDbFail("you already login"))
	}
}

func (c *MainController) Logout() {
	c.DelSession("role")
	c.Ctx.WriteString(MsgResponse("logout success"))

}
