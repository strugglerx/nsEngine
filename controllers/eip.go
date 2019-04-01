package controllers

import (
	"encoding/json"
	"server/utils"

	"github.com/astaxie/beego"
)

type EipController struct {
	beego.Controller
}

func (c *EipController) Entry() {
	var user string
	customtype := c.GetString("type")
	user = c.Ctx.Request.Header.Get("auth")
	//限制只有type等于info才接受user的值
	if user == "" && customtype == "info" {
		user = c.GetString("user")
	}
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	date := c.GetString("date")
	if user != "" && customtype != "" || len(date) == 10 {
		var cryptoUser string
		if customtype == "info" {
			cryptoUser = utils.CustomAesEncrypt(user)
			result, err := utils.EipEntry(user, customtype, date)
			if err == nil {
				var unresult interface{}
				json.Unmarshal([]byte(result), &unresult)
				res := CustomEipResponse{"1000", 0, unresult, cryptoUser}
				c.Ctx.WriteString(res.JsonFormat())
			} else {
				c.Ctx.WriteString(ApiDbFail())
			}
		} else if len(user) == 24 {
			cryptoUser = utils.CustomAesDecrypt(user)
			result, err := utils.EipEntry(cryptoUser, customtype, date)
			if err == nil {
				var unresult interface{}
				json.Unmarshal([]byte(result), &unresult)
				c.Ctx.WriteString(ApiResponse(unresult))
			} else {
				c.Ctx.WriteString(ApiDbFail())
			}
		} else {
			c.Ctx.WriteString(ApiFail())
		}
	} else {
		c.Ctx.WriteString(ApiFail())
	}
}

func (c *EipController) Library() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	book := c.GetString("book")
	page := c.GetString("page")

	if book != "" && page != "" {
		result := utils.EipLibrary(book, page)
		if result != "-1" {
			var unresult interface{}
			json.Unmarshal([]byte(result), &unresult)
			c.Ctx.WriteString(ApiResponse(unresult))
		} else {
			c.Ctx.WriteString(ApiDbFail())
		}
	} else {
		c.Ctx.WriteString(ApiFail())
	}
}

func (c *EipController) LibraryDetail() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	id := c.GetString("id")
	if id != "" {
		result := utils.EipLibraryDetail(id)
		if result != "-1" {
			var unresult interface{}
			json.Unmarshal([]byte(result), &unresult)
			c.Ctx.WriteString(ApiResponse(unresult))
		} else {
			c.Ctx.WriteString(ApiDbFail())
		}
	} else {
		c.Ctx.WriteString(ApiFail())
	}
}

func (c *EipController) SportEntry() {
	//c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")
	user := c.GetString("user")
	pwd := c.GetString("pwd")
	if user != "" && pwd != "" {
		result := utils.SportCurl(user, pwd)
		if result != "-1" {
			c.Ctx.WriteString(ApiResponse(result))
		} else {
			c.Ctx.WriteString(ApiDbFail())
		}
	} else {
		c.Ctx.WriteString(ApiFail())
		// info := make(map[string]string)
		// info["code"] = "203"
		// info["status"] = "-1"
		// c.Data["json"] = &info
		// c.ServeJSON()
	}
}
