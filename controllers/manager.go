package controllers

import (
	"fmt"
	"server/models"
	"server/utils"

	"github.com/astaxie/beego"
)

type ManagerController struct {
	beego.Controller
}

func (c *ManagerController) ManagerInfo() {
	sess := c.GetSession("user")
	c.Ctx.WriteString(ApiResponse(sess))
}

func (c *ManagerController) ManagerIndex() {
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
	Uid, ok := models.ArtInsert(title, author, content, date)
	if ok {
		models.IAInsert(path, "/pages/tools/article/article?ID="+Uid, title, Uid, date, author)
	}
	c.Ctx.WriteString(ApiResponse(ok))
}

//广告数据增加 AdInsert
func (c *ManagerController) AdInsert() {
	uuid := c.GetString("uuid")
	dateStart := c.GetString("dateStart")
	dateEnd := c.GetString("dateEnd")
	path := c.GetString("path")
	remark := c.GetString("remark")
	if uuid != "" && dateStart != "" && dateEnd != "" && path != "" && remark != "" {
		ok := models.AdInsert(uuid, path, dateStart, dateEnd, remark)
		if ok {
			c.Ctx.WriteString(ApiResponse(ok))
			return
		}
		c.Ctx.WriteString(ApiDbFail())
	}

}

//广告数据列表
func (c *ManagerController) AdList() {
	result := models.AdList()
	c.Ctx.WriteString(ApiResponse(result))
}

func (c *ManagerController) Update() {
	type_ := c.GetString("type")
	id_ := c.GetString("id")
	name_ := c.GetString("newName")
	switch type_ {
	case "article":
		title_ := c.GetString("title")
		author_ := c.GetString("author")
		content_ := c.GetString("content")
		err := models.ArtRename(id_, title_, author_, content_)
		if err == nil {
			err := models.IARename(id_, title_, author_)
			if err != nil {
				c.Ctx.WriteString(MsgDbFail("Something Wrong"))
				return
			}
		}
	case "advertisement":
		err := models.AdRename(id_, name_)
		if err != nil {
			c.Ctx.WriteString(MsgDbFail("Something Wrong"))
			return
		}
	case "keywords":
		err := models.KeywordRename(id_, name_)
		if err != nil {
			c.Ctx.WriteString(MsgDbFail("Something Wrong"))
			return
		}
	default:
		c.Ctx.WriteString(MsgFail("Params Wrong"))
		return
	}
	c.Ctx.WriteString(MsgResponse("Update Success"))
}

//删除数据
func (c *ManagerController) DbDelete() {
	type_ := c.GetString("type")
	id := c.GetString("id")
	var ok bool
	switch type_ {
	case "article":
		ok = models.ArtDel(id)
		if ok {
			models.IADel(id)
		}
		c.Ctx.WriteString(MsgResponse("success"))
	case "step":
		ok = models.StepDelete(id)
		c.Ctx.WriteString(MsgResponse("success"))
	case "ads":
		ok = models.AdDel(id)
		c.Ctx.WriteString(MsgResponse("success"))
	case "keyword":
		ok = models.KeywordDelete(id)
		c.Ctx.WriteString(MsgResponse("success"))
	case "feedback":
		_id := c.GetString("_id")
		fmt.Println(_id)
		ok = models.FeedBackDelete(_id)
		c.Ctx.WriteString(MsgResponse("success"))
	default:
		c.Ctx.WriteString(MsgDbFail("fail"))
	}
}

func (c *ManagerController) FeedBackList() {
	result := models.FeedbackList()
	c.Ctx.WriteString(ApiResponse(result))
}

func (c *ManagerController) FeedBackAction() {
	openid := c.GetString("openId")
	content := c.GetString("text")
	accesstoken := utils.WxGetAccessToken().Token
	result := utils.WxSendMsg(content, openid, accesstoken)
	if result == 0 {
		c.Ctx.WriteString(MsgResponse(result))
	} else {
		c.Ctx.WriteString(MsgDbFail(result))
	}

}

//option
func (c *ManagerController) Option() {
	id, _ := c.GetInt("id")
	status, _ := c.GetInt("status")
	ok := models.OptUpConf(id, status)
	if ok {
		c.Ctx.WriteString(MsgResponse("success"))
	} else {
		c.Ctx.WriteString(MsgDbFail("fail"))
	}
}

//改变密码
func (c *ManagerController) ChangePwd() {
	user := c.GetString("user")
	passwd := c.GetString("passwd")
	cryptoPwd := utils.Md5String(passwd)
	ok := models.UserUpdate(user, cryptoPwd)
	if ok {
		c.Ctx.WriteString(MsgResponse("success"))
	} else {
		c.Ctx.WriteString(MsgDbFail("fail"))
	}
}

//keywords
func (c *ManagerController) KeywordInsert() {
	keyword := c.GetString("keyword")
	content := c.GetString("content")
	ok := models.KeywordInsert(content, keyword)
	if ok {
		c.Ctx.WriteString(ApiResponse(true))
		return
	}
	c.Ctx.WriteString(ApiDbFail())
}

func (c *ManagerController) KeywordList() {
	result := models.KeywordList()
	if len(result) > 0 && result != nil {
		c.Ctx.WriteString(ApiResponse(result))
		return
	}
	c.Ctx.WriteString(ApiDbFail())
}

func (c *ManagerController) FormId() {
	_type := c.GetString("_type")
	openid := c.GetString("openId")
	page, _ := c.GetInt("p") //页数一页一百条数据
	type List struct {
		Count int         `json:"count,omitempty"`
		List  interface{} `json:"list,omitempty"`
	}
	var result interface{}
	//分页数据结构
	switch _type {
	case "search":
		key := c.GetString("name")
		/* 		nameList, err := models.FormFindName(key)
		   		if err != nil {
		   			c.Ctx.WriteString(MsgDbFail("fail"))
		   			return
		   		}
		   		collegeList, err := models.FormFindCollege(key)
		   		if err != nil {
		   			c.Ctx.WriteString(MsgDbFail("fail"))
		   			return
		   		}
		   		majorList, err := models.FormFindMajor(key)
		   		if err != nil {
		   			c.Ctx.WriteString(MsgDbFail("fail"))
		   			return
		   		}

		   		for _, item := range collegeList {
		   			nameList = append(nameList, item)
		   		}
		   		for _, item := range majorList {
		   			nameList = append(nameList, item)
		   		}
		   		result = models.FormUnique(nameList) */
		resultList, err := models.FormFindKeyWord(key)
		if err != nil {
			c.Ctx.WriteString(MsgDbFail("fail"))
			return
		}
		result = resultList
	case "one":
		result = models.FormFirstId(openid)
		models.FormPop(openid)
	case "del":
		result = models.FormDel(openid)
	default:
		var result List
		result.List = models.FormList(page)
		result.Count = models.FormCount()
		c.Ctx.WriteString(ApiResponse(result))
		return
	}
	c.Ctx.WriteString(ApiResponse(result))

}

func (c *ManagerController) PushToast() {
	body := c.Ctx.Input.RequestBody
	response, err := utils.WxPushTemplate(body)
	if err != nil {
		c.Ctx.WriteString(ApiFail())
		return
	}
	c.Ctx.WriteString(ApiResponse(response))
}
