package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
	"github.com/xlstudio/wxbizdatacrypt"
	"server/models"
	"server/utils"
)

type ApiController struct {
	beego.Controller
}
//apiIndex

func (c *ApiController) ApiIndex() {
	info := MainResponse{"1000", 0,"Path Not Found"}
	c.Ctx.WriteString(info.JsonParse())
}

//信息列表
func (c *ApiController) Get() {
	result:=models.ArtList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}

//查看详情
func (c *ApiController) ArtDetail() {
	p := c.GetString("p")
	models.ArtUpView(p)
	r :=models.ArtDetail(p)
	info := CustomResponse{"1000", 0,r[0]}
	c.Ctx.WriteString(info.JsonFormat())
}
//like接口
func (c *ApiController) ArtUplike() {
	p := c.GetString("p")
	u := c.GetString("user")

	likeSataus:=models.ArtFindLike(p,u)
	if likeSataus{
		models.ArtUpLike(p,1)
	}else {
		models.ArtUpLike(p,-1)
	}
	info := CustomResponse{"1000", 0,likeSataus}
	c.Ctx.WriteString(info.JsonFormat())
}
//运动圈
func (c *ApiController) StepList(){
	result:=models.StepList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ApiController) StepUpdate(){
	code := c.GetString("js_code")
	encry := c.GetString("encryptedData")
	iv := c.GetString("iv")
	nickName := c.GetString("nickName")
	avatarUrl := c.GetString("avatarUrl")
	wxsession:=utils.WxSession(code)
	if wxsession.Status{
		one :=wxbizdatacrypt.WxBizDataCrypt{wxsession.Appid,wxsession.Session}
		result,_ :=one.Decrypt(encry,iv,true)
		userStep :=gjson.Get(result.(string),"stepInfoList.30.step").Int()
		fmt.Printf("%+v",userStep)
		verify:=models.StepFindOne(nickName)
		if verify{
			models.StepUpdate(avatarUrl,nickName,userStep)
		}else {
			models.StepInsert(avatarUrl,nickName,userStep)
		}
		var response models.StepUser
		response.Step = userStep
		response.Nickname = nickName
		response.AvatarUrl = avatarUrl

		info := CustomResponse{"1000", 0,response}
		c.Ctx.WriteString(info.JsonFormat())
	}
}

//招聘信息流

func (c *ApiController) JobList(){
	page,err:=c.GetInt("p")
	if err!=nil||page<=0{
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}else{
		result:=models.JobList(page-1)
		if len(result)!=0{
			info := CustomResponse{"1000", 0,result}
			c.Ctx.WriteString(info.JsonFormat())
		}else {
			info := CustomResponse{"1001", -1,nil}
			c.Ctx.WriteString(info.JsonFormat())
		}

	}
}

func (c *ApiController) JobDetail(){
	title:=c.GetString("title")
	date:=c.GetString("date")
	if title!=""&&date!=""{
		models.JobUpView(title,date)
		result:=models.JobDetail(title,date)
		if len(result)!=0{
			info := CustomResponse{"1000", 0,result[0]}
			c.Ctx.WriteString(info.JsonFormat())
		}else {
			info := CustomResponse{"1001", -1,nil}
			c.Ctx.WriteString(info.JsonFormat())
		}
	}else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}
}

//index页面

func (c *ApiController) IndexSwiper(){
	result:=models.IAList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ApiController) IndexConfig(){
	result:=models.OptList()
	info := CustomResponse{"1000", 0,result}
	c.Ctx.WriteString(info.JsonFormat())
}

//地图标记

func (c *ApiController) PointIndex(){
	type object struct {
		CenterPoint struct{
			Saihan []interface{}  `json:"saihan,omitempty"`
			Shengle []interface{} `json:"shengle,omitempty"`
		} `json:"centerPoint,omitempty"`
		MapSign struct{
			Saihan []interface{} `json:"saihan,omitempty"`
			Shengle []interface{} `json:"shengle,omitempty"`
		}  `json:"mapSign,omitempty"`

	}
	type pointList []interface{}
	var data object
	data.CenterPoint.Saihan = models.PointList("saihan")[0].Point
	data.CenterPoint.Shengle = models.PointList("shengle")[0].Point
	var saihanPoint pointList
	for _,v:=range(models.SignList("saihan")){
		saihanPoint=append(saihanPoint,v.Point)
	}
	data.MapSign.Saihan = saihanPoint
	var shenglePoint pointList
	for _,v:=range(models.SignList("shengle")){
		shenglePoint=append(shenglePoint,v.Point)
	}
	data.MapSign.Shengle = shenglePoint
	info := CustomResponse{"1000", 0,data}
	c.Ctx.WriteString(info.JsonFormat())
}

func (c *ApiController) PointPost(){
	digest := c.GetString("digest")
	id ,_:= c.GetInt("id")
	type_ := c.GetString("type")
	con := c.GetString("content")
	lat := c.GetString("latitude")
	lon := c.GetString("longitude")
	if digest==""||digest=="push"{
		ok := models.SignPush(type_,con,lat,lon,id)
		if ok{
			info := MainResponse{"1000", 0,"Push Success"}
			c.Ctx.WriteString(info.JsonParse())
		}else {
			info := MainResponse{"1001", -1,"Push Failed"}
			c.Ctx.WriteString(info.JsonParse())
		}
	}else if digest=="pull"{
		ok := models.SignPull(type_,con,lat,lon,id)
		if ok{
			info := MainResponse{"1000", 0,"Pull Success"}
			c.Ctx.WriteString(info.JsonParse())
		}
	}else if digest=="set"{
		newCon := c.GetString("newContent")

		ok := models.SignSet(type_,con,lon,newCon,id)
		if ok{
			info := MainResponse{"1000", 0,"Set Success"}
			c.Ctx.WriteString(info.JsonParse())
		}
	}else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}
}

//客服接口
func (c *ApiController) Msg(){
	echostr:=c.GetString("echostr")
	if echostr!=""{
		c.Ctx.WriteString(echostr)

	}else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}

}

func (c *ApiController) MsgPost(){
	token:="strugglerno1"
	content:=c.GetString("Content")
	openid:=c.GetString("FromUserName")
    signature:= c.GetString("signature")
	timestamp:= c.GetString("timestamp")
	nonce:= c.GetString("nonce")
	verifySha1:=utils.VerifySha1(token,timestamp,nonce)
	if verifySha1==signature{
		accesstoken:=utils.WxGetAccessToken()
		result:=utils.WxSendMsg(content,openid,accesstoken)
		info := CustomResponse{"1000", result,true}
		c.Ctx.WriteString(info.JsonFormat())
	}else {
		info := CustomResponse{"1002", -2,nil}
		c.Ctx.WriteString(info.JsonFormat())
	}
}





