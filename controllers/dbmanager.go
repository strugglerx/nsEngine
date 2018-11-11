package controllers

import (
	"github.com/astaxie/beego"
	"server/models"
	"server/utils"
)

type ManagerController struct {
	beego.Controller
}

func (c *ManagerController) ManagerInfo(){
	sess :=c.GetSession("user")
	info := CustomResponse{"1000", 0,sess}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ManagerController) ManagerIndex(){
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.TplName = "manager.html"
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
//广告数据增加 AdInsert
func (c *ManagerController) AdInsert() {
	uuid := c.GetString("uuid")
	dateStart := c.GetString("dateStart")
	dateEnd := c.GetString("dateEnd")
	path := c.GetString("path")
	remark := c.GetString("remark")
	if (uuid!=""&&dateStart!=""&&dateEnd!=""&&path!=""&&remark!=""){
		ok:=models.AdInsert(uuid,path,dateStart,dateEnd,remark)
		if ok{
			info := CustomResponse{"1000", 0,ok}
			c.Ctx.WriteString(info.JsonFormat())
			return
		}
		info := CustomResponse{"1001", -1,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}

}
//广告数据列表
func (c *ManagerController) AdList() {
	result:=models.AdList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}
//删除数据
func (c *ManagerController) DbDelete() {
	type_ := c.GetString("type")
	id := c.GetString("id")
	var ok bool
	switch type_ {
	case"article":
		ok =models.ArtDel(id)
		if ok{
			models.IADel(id)
		}
		info := MainResponse{"1000", 0,"success"}
		c.Ctx.WriteString(info.JsonParse())
	case "step":
		ok =models.StepDelete(id)
		info := MainResponse{"1000", 0,"success"}
		c.Ctx.WriteString(info.JsonParse())
	case "ads":
		ok =models.AdDel(id)
		info := MainResponse{"1000", 0,"success"}
		c.Ctx.WriteString(info.JsonParse())
	case "keyword":
		ok =models.KeywordDelete(id)
		info := MainResponse{"1000", 0,"success"}
		c.Ctx.WriteString(info.JsonParse())
	default:
		info := MainResponse{"1001", -1,"fail"}
		c.Ctx.WriteString(info.JsonParse())
	}
}

func (c *ManagerController) FeedBackList(){
	result:=models.FeedbackList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ManagerController) FeedBackSendMsg(){
	openid:=c.GetString("openId")
	content:=c.GetString("text")
	accesstoken:=utils.WxGetAccessToken()
	result:=utils.WxSendMsg(content,openid,accesstoken)
	if result==0{
		info := MainResponse{"1000", 0,result}
		c.Ctx.WriteString(info.JsonParse())
	}else{
		info := MainResponse{"1001", -1,result}
		c.Ctx.WriteString(info.JsonParse())
	}

}
//option
func (c *ManagerController) Option(){
	id,_ :=c.GetInt("id")
	status,_:=c.GetInt("status")
	ok:=models.OptUpConf(id,status)
	if ok{
		info := MainResponse{"1000", 0,"success"}
		c.Ctx.WriteString(info.JsonParse())
	}else {
		info := MainResponse{"1001", -1,"fail"}
		c.Ctx.WriteString(info.JsonParse())
	}
}
//改变密码
func (c *ManagerController) ChangePwd(){
	user:=c.GetString("user")
	passwd:=c.GetString("passwd")
	cryptoPwd:=utils.Md5String(passwd)
	ok:=models.UserUpdate(user,cryptoPwd)
	if ok{
		info := MainResponse{"1000", 0,"success"}
		c.Ctx.WriteString(info.JsonParse())
	}else {
		info := MainResponse{"1001", -1,"fail"}
		c.Ctx.WriteString(info.JsonParse())
	}
}
//keywords
func (c *ManagerController) KeywordInsert() {
	keyword :=c.GetString("keyword")
	content :=c.GetString("content")
	ok:=models.KeywordInsert(content,keyword)
	if ok{
		info := CustomResponse{"1000", 0,true}
		c.Ctx.WriteString(info.JsonFormat())
		return
	}
	info := CustomResponse{"1001", -1,nil}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ManagerController) KeywordList() {
	result:=models.KeywordList()
	if len(result)>0&&result!=nil{
		info := CustomResponse{"1000", 0,result}
		c.Ctx.WriteString(info.JsonFormat())
		return
	}
	info := CustomResponse{"1001", -1,nil}
	c.Ctx.WriteString(info.JsonFormat())
}


