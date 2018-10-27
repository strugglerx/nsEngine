package main

import (
	"github.com/astaxie/beego"
	"server/controllers"
	"server/models/mymongo"
	_ "server/routers"
)


func main() {
	defer mymongo.CloseMgo()
	beego.BConfig.WebConfig.Session.SessionName="sessionId"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime=1000
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie=true
    //错误页处理
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetStaticPath("/static","static")
	beego.Run()
}
