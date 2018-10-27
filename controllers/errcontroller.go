package controllers

import (
	"github.com/astaxie/beego"
)
type ErrorController struct {
	beego.Controller
}
func (c *ErrorController) Error404() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	//c.Data["content"] = "page not found"
	c.TplName = "404.html"
}
func (c *ErrorController) Error405() {
	//c.Data["content"] = "page not found"
	c.TplName = "404.html"

}
func (c *ErrorController) Error501() {
	//c.Data["content"] = "server error"
	c.TplName = "404.html"
}
func (c *ErrorController) ErrorDb() {
	//c.Data["content"] = "database is now down"
	c.TplName = "404.html"
}