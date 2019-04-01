package controllers

import (
	"encoding/json"
	"server/models"
	"server/utils"

	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
	"github.com/xlstudio/wxbizdatacrypt"
)

type ApiController struct {
	beego.Controller
}

//apiIndex

func (c *ApiController) ApiIndex() {
	c.Ctx.WriteString(MsgResponse("Path Not Found"))
}

//信息列表
func (c *ApiController) ArtList() {
	result := models.ArtList()
	c.Ctx.WriteString(ApiResponse(result))
}

//查看详情
func (c *ApiController) ArtDetail() {
	p := c.GetString("p")
	r := models.ArtDetail(p)
	if len(r) == 1 {
		models.ArtUpView(p)
		c.Ctx.WriteString(ApiResponse(r[0]))
	} else {
		c.Ctx.WriteString(ApiDbFail())
	}

}

//like接口
func (c *ApiController) ArtUplike() {
	p := c.GetString("p")
	u := c.GetString("user")

	likeSataus := models.ArtFindLike(p, u)
	if likeSataus {
		models.ArtUpLike(p, 1)
	} else {
		models.ArtUpLike(p, -1)
	}
	c.Ctx.WriteString(ApiResponse(likeSataus))
}

//运动圈
func (c *ApiController) StepList() {
	result := models.StepList()
	c.Ctx.WriteString(ApiResponse(result))
}

func (c *ApiController) StepUpdate() {
	code := c.GetString("js_code")
	encry := c.GetString("encryptedData")
	iv := c.GetString("iv")
	nickName := c.GetString("nickName")
	avatarUrl := c.GetString("avatarUrl")
	wxsession := utils.WxSession(code)
	if wxsession.Status {
		one := wxbizdatacrypt.WxBizDataCrypt{
			AppID:      wxsession.Appid,
			SessionKey: wxsession.Session}
		result, _ := one.Decrypt(encry, iv, true)
		userStep := gjson.Get(result.(string), "stepInfoList.30.step").Int()
		verify := models.StepFindOne(nickName)
		if verify {
			models.StepUpdate(avatarUrl, nickName, userStep)
		} else {
			models.StepInsert(avatarUrl, nickName, userStep)
		}
		var response models.StepUser
		response.Step = userStep
		response.Nickname = nickName
		response.AvatarUrl = avatarUrl
		c.Ctx.WriteString(ApiResponse(response))
	}
}

//招聘信息流

func (c *ApiController) JobList() {
	page, err := c.GetInt("p")
	if err != nil || page <= 0 {
		c.Ctx.WriteString(ApiFail())
	} else {
		result := models.JobList(page - 1)
		if len(result) != 0 {
			c.Ctx.WriteString(ApiResponse(result))
		} else {
			c.Ctx.WriteString(ApiDbFail())
		}

	}
}

func (c *ApiController) JobDetail() {
	title := c.GetString("title")
	date := c.GetString("date")
	if title != "" && date != "" {
		models.JobUpView(title, date)
		result := models.JobDetail(title, date)
		if len(result) != 0 {
			c.Ctx.WriteString(ApiResponse(result[0]))
		} else {
			c.Ctx.WriteString(ApiDbFail())
		}
	} else {
		c.Ctx.WriteString(ApiFail())
	}
}

//index页面

func (c *ApiController) IndexSwiper() {
	result := models.IAList()
	c.Ctx.WriteString(ApiResponse(result))
}

func (c *ApiController) IndexConfig() {
	result := models.OptList()
	c.Ctx.WriteString(ApiResponse(result))
}

//地图标记

func (c *ApiController) PointIndex() {
	type object struct {
		CenterPoint struct {
			Saihan  []interface{} `json:"saihan,omitempty"`
			Shengle []interface{} `json:"shengle,omitempty"`
		} `json:"centerPoint,omitempty"`
		MapSign struct {
			Saihan  []interface{} `json:"saihan,omitempty"`
			Shengle []interface{} `json:"shengle,omitempty"`
		} `json:"mapSign,omitempty"`
	}
	type pointList []interface{}
	var data object
	data.CenterPoint.Saihan = models.PointList("saihan")[0].Point
	data.CenterPoint.Shengle = models.PointList("shengle")[0].Point
	var saihanPoint pointList
	for _, v := range models.SignList("saihan") {
		saihanPoint = append(saihanPoint, v.Point)
	}
	data.MapSign.Saihan = saihanPoint
	var shenglePoint pointList
	for _, v := range models.SignList("shengle") {
		shenglePoint = append(shenglePoint, v.Point)
	}
	data.MapSign.Shengle = shenglePoint
	c.Ctx.WriteString(ApiResponse(data))
}

func (c *ApiController) PointPost() {
	digest := c.GetString("digest")
	id, _ := c.GetInt("id")
	type_ := c.GetString("type")
	con := c.GetString("content")
	lat := c.GetString("latitude")
	lon := c.GetString("longitude")
	if digest == "" || digest == "push" {
		ok := models.SignPush(type_, con, lat, lon, id)
		if ok {
			c.Ctx.WriteString(MsgResponse("Push Success"))
		} else {
			c.Ctx.WriteString(MsgDbFail("Push Failed"))
		}
	} else if digest == "pull" {
		ok := models.SignPull(type_, con, lon, id)
		if ok {
			c.Ctx.WriteString(MsgResponse("Pull Success"))
		}
	} else if digest == "set" {
		newCon := c.GetString("newName")

		ok := models.SignSet(type_, con, lon, newCon, id)
		if ok {
			c.Ctx.WriteString(MsgResponse("Set Success"))
		}
	} else {
		c.Ctx.WriteString(ApiFail())
	}
}

//客服接口
func (c *ApiController) Msg() {
	echostr := c.GetString("echostr")
	if echostr != "" {
		c.Ctx.WriteString(echostr)

	} else {
		c.Ctx.WriteString(ApiFail())
	}

}

func (c *ApiController) MsgPost() {
	//微信发送的json格式
	type sendMsg struct {
		ToUserName   string
		FromUserName string
		CreateTime   int
		MsgType      string
		Content      string
		MsgId        int
	}
	var result sendMsg
	token := "strugglerno1"
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &result)
	if err != nil {
		c.Ctx.WriteString(ApiFail())
		return
	}
	content := result.Content
	openid := result.FromUserName
	createTime := result.CreateTime
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	verifySha1 := utils.VerifySha1(token, timestamp, nonce)
	//verifySha1==signature
	if verifySha1 == signature {
		accesstoken := utils.WxGetAccessToken().Token
		if content == "" {
			result := utils.WxSendMsg(`内师助手欢迎你,如果你现在有问题在下面留言即可,管理员看到会立刻回复你的!`, openid, accesstoken)
			c.Ctx.WriteString(ApiResponse(result))
			return
		} else {
			//留言收集
			if findRestut := models.KeywordFind(string(content)); findRestut.Content != "" {
				utils.WxSendMsg(findRestut.Content, openid, accesstoken)
				c.Ctx.WriteString(ApiResponse(true))
				return
			} else {
				utils.WxSendMsg("留言内容:["+content+"]已收到!等待管理员回复中,请耐心等待哦！也可以直接加开发者微信strongdreams帮你解决问题", openid, accesstoken)
				info, _ := models.FormFindOpenId(openid)
				models.FeedbackInsert(info.NickName, info.Name, info.AvatarUrl, info.College, openid, content, createTime)
				c.Ctx.WriteString(ApiResponse(true))
				return
			}
		}
	} else {
		c.Ctx.WriteString(ApiFail())
		return
	}
}

//首页广告列表
//广告数据列表
func (c *ApiController) AdList() {
	result := models.AdListLimit()
	c.Ctx.WriteString(ApiResponse(result))
}

//收集FormId

func (c *ApiController) CollectFormId() {
	type Formid struct {
		FormId string `json:"formid"`
		Openid string `json:"openid"`
	}
	FormId := c.GetString("formId")
	Code := c.GetString("code")
	nickName := c.GetString("nickName")
	name := c.GetString("name")
	college := c.GetString("college")
	major := c.GetString("major")
	avatarUrl := c.GetString("avatarUrl")
	authid := c.GetString("auth")
	wxsession := utils.WxSession(Code)
	if wxsession.Status {
		res := Formid{FormId, wxsession.Openid}
		// fmt.Printf("%+v", wxsession)
		exist := models.FormExist(res.Openid)
		if exist {
			formIdCount := models.FormSize(res.Openid)
			if formIdCount < 14 {
				models.FormAdd(authid, res.Openid, res.FormId, nickName, name, college, major, avatarUrl)
			} else if formIdCount == 14 {
				models.FormShift(res.Openid)
				models.FormAdd(authid, res.Openid, res.FormId, nickName, name, college, major, avatarUrl)
			}

		} else {
			models.FormInsert(authid, res.Openid, res.FormId, nickName, name, college, major, avatarUrl)
		}
		// c.Data["json"] = &res
		// c.ServeJSON()
		c.Ctx.WriteString(ApiResponse(true))
	} else {
		// c.Data["json"] = nil
		// c.ServeJSON()
		c.Ctx.WriteString(ApiDbFail())
	}

}

func (c *ApiController) Classswitch() {
	type_ := c.GetString("type_")
	authId := c.GetString("authId")
	if type_ == "switch" {
		_switch_, err := models.FormClass_(authId)
		if err != nil {
			c.Ctx.WriteString(ApiDbFail())
			return
		}
		models.FormChangeClass_(authId, !_switch_)
		c.Ctx.WriteString(ApiResponse(!_switch_))
		return
	} else if type_ == "verify" {
		_switch_, err := models.FormClass_(authId)
		if err != nil {
			c.Ctx.WriteString(ApiDbFail())
			return
		}
		c.Ctx.WriteString(ApiResponse(_switch_))
		return
	}
	c.Ctx.WriteString(ApiFail())
}
