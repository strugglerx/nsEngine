/*
 * @Description:
 * @Author: Moqi
 * @Date: 2018-12-12 10:36:40
 * @Email: str@li.cm
 * @Github: https://github.com/strugglerx
 * @LastEditors: Moqi
 * @LastEditTime: 2018-12-12 10:36:42
 */

package main

import (
	"server/controllers"
	"server/models/mymongo"
	_ "server/routers"

	"github.com/astaxie/beego"
)

func main() {
	defer mymongo.CloseMgo()
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	//错误页处理
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
