package controllers

import (
	"encoding/json"
	"server/utils"

	"github.com/astaxie/beego"
)

type EipController struct {
	beego.Controller
}

//api响应生成结构体
type CustomResponse struct {
	Code   string `json:"code"`
	Status int64    `json:"status"`
	Data interface{}    `json:"data,omitempty"`
}

func (res *CustomResponse) JsonFormat() string {
	datas, _ := json.Marshal(res)
	return string(datas)
}

func (c *EipController) Entry() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	user := c.GetString("user")
	customtype := c.GetString("type")
	date := c.GetString("date")
	if user != "" && customtype != "" || len(date) == 10 {
		result := utils.EipEntry(user, customtype, date)
		if result != "-1" {
			var unresult interface{}
			json.Unmarshal([]byte(result),&unresult)
			info := CustomResponse{"1000", 0,unresult}
			c.Ctx.WriteString(info.JsonFormat())
		} else {
			info := CustomResponse{"1001", -1,nil}
			c.Ctx.WriteString(info.JsonFormat())
		}
	} else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}
}

func (c *EipController) Library() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	book := c.GetString("book")
	page := c.GetString("page")

	if book != "" && page != ""  {
		result := utils.EipLibrary(book,page)
		if result != "-1" {
			var unresult interface{}
			json.Unmarshal([]byte(result),&unresult)
			info := CustomResponse{"1000", 0,unresult}
			c.Ctx.WriteString(info.JsonFormat())
		} else {
			info := CustomResponse{"1001", -1,nil}
			c.Ctx.WriteString(info.JsonFormat())
		}
	} else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}
}

func (c *EipController) LibraryDetail() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	id := c.GetString("id")
	if id != ""  {
		result := utils.EipLibraryDetail(id)
		if result != "-1" {
			var unresult interface{}
			json.Unmarshal([]byte(result),&unresult)
			info := CustomResponse{"1000", 0,unresult}
			c.Ctx.WriteString(info.JsonFormat())
		} else {
			info := CustomResponse{"1001", -1,nil}
			c.Ctx.WriteString(info.JsonFormat())
		}
	} else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}
}



func (c *EipController) SportEntry() {
	//c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")
	user := c.GetString("user")
	pwd := c.GetString("pwd")
	if user != "" && pwd != "" {
		result := utils.SportCurl(user, pwd)
		if result != "-1" {
			info := CustomResponse{"1000", 0,result}
			c.Ctx.WriteString(info.JsonFormat())
		} else {
			info := CustomResponse{"1001", -1,nil}
			c.Ctx.WriteString(info.JsonFormat())
		}
	} else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
		// info := make(map[string]string)
		// info["code"] = "203"
		// info["status"] = "-1"
		// c.Data["json"] = &info
		// c.ServeJSON()
	}
}

