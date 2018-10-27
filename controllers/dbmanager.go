package controllers

import (
	"github.com/astaxie/beego"
	"server/models"
)

type ManagerController struct {
	beego.Controller
}

//增加信息
func (c *ManagerController) ArtInsert() {
	title := c.GetString("title")
	date := c.GetString("date")
	author := c.GetString("author")
	content := c.GetString("content")
	path := c.GetString("path")
	Uid,ok:=models.ArtInsert(title,author,content,date)
	if ok {
		models.IAInsert(path,"/pages/tools/article/article?ID="+Uid,title,Uid,date,author)
	}
	info := CustomResponse{"1000", 0,ok}
	c.Ctx.WriteString(info.JsonFormat())
}
//删除
func (c *ManagerController) DbDelete() {
	type_ := c.GetString("type")
	id := c.GetString("id")
	var ok bool
	if type_ =="article"{
		ok =models.ArtDel(id)
		if ok{
			models.IADel(id)
		}
	}
	info := CustomResponse{"1000", 0,ok}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ManagerController) FeedBackList(){
	result:=models.FeedbackList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}
