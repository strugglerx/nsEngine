package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type form struct {
	Id         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	OpenId     string        `json:"openId,omitempty" bson:"openId,omitempty"`
	AuthId     string        `json:"authId,omitempty" bson:"authId,omitempty"`
	UpdateTime string        `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
	CreateTime int           `json:"createTime,omitempty" bson:"createTime,omitempty"`
	FormId     []string      `json:"formId" bson:"formId,omitempty"`
	NickName   string        `json:"nickName,omitempty" bson:"nickName,omitempty"`
	Name       string        `json:"name,omitempty" bson:"name,omitempty"`
	College    string        `json:"college,omitempty" bson:"college,omitempty"`
	Major      string        `json:"major,omitempty" bson:"major,omitempty"`
	AvatarUrl  string        `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	Class_     bool          `json:"class_,omitempty" bson:"class_,omitempty"`
}

type formSize struct {
	OpenId string `bson:"openId"`
	Count  int    `bson:"count"`
}

//all list
func FormList(page int) []form {
	result := []form{}
	FormId.
		Find(bson.M{}).
		Select(bson.M{"_id": 0}).
		Skip(page * 100).
		Limit(100).
		All(&result)
	return result
}

func FormCount() int {
	var count int
	count, err := FormId.Count()
	if err != nil {
		return 100 //默认显示一页数据100条
	}
	return count
}

//返回第一个FormId
func FormFirstId(openid string) string {
	var result form
	FormId.
		Find(bson.M{"openId": openid}).
		Select(bson.M{"formId": bson.M{"$slice": 1}, "_id": 0}).
		One(&result)
		//返回第一个formid
	if len(result.FormId) <= 0 {
		return ""
	}
	return result.FormId[0]
}

//返回最后一个FormId
func FormLastId(openid string) string {
	var result form
	FormId.
		Find(bson.M{"openId": openid}).
		Select(bson.M{"formId": bson.M{"$slice": -1}, "_id": 0}).
		One(&result)
		//返回第一个formid
	// fmt.Printf("%+v", result)
	if len(result.FormId) <= 0 {
		return ""
	}
	return result.FormId[0]
}

//删除第一个formId
func FormShift(openid string) bool {
	err := FormId.
		Update(bson.M{"openId": openid}, bson.M{"$pop": bson.M{"formId": -1}})
	if err != nil {
		return false
	}
	return true
}

//删除最后一个formId
func FormPop(openid string) bool {
	err := FormId.
		Update(bson.M{"openId": openid}, bson.M{"$pop": bson.M{"formId": 1}})
	if err != nil {
		return false
	}
	return true
}

func FormChangeClass_(authId string, _switch_ bool) bool {
	err := FormId.
		Update(bson.M{"authId": authId}, bson.M{"$set": bson.M{"class_": _switch_}})
	if err != nil {
		return false
	}
	return true
}

//通openId获取用户信息，给feedback用
func FormFindOpenId(openId string) (form, error) {
	var result form
	err := FormId.
		Find(bson.M{"openId": openId}).
		One(&result)
	if err != nil {
		return result, errors.New("not exist")
	}
	return result, nil
}

// $or查询

func FormFindKeyWord(key string) ([]form, error) {
	var result []form
	err := FormId.
		Find(bson.M{"$or": []bson.M{bson.M{"nickName": bson.M{"$regex": bson.RegEx{Pattern: key, Options: "im"}}},
			bson.M{"name": bson.M{"$regex": bson.RegEx{Pattern: key, Options: "im"}}},
			bson.M{"college": bson.M{"$regex": bson.RegEx{Pattern: key, Options: "im"}}},
			bson.M{"major": bson.M{"$regex": bson.RegEx{Pattern: key, Options: "im"}}}}}).
		All(&result)
	if err != nil {
		return result, errors.New("not exist")
	}
	return result, nil
}

//通过名字模糊查找
func FormFindName(name string) ([]form, error) {
	var result []form
	err := FormId.
		Find(bson.M{"name": bson.M{"$regex": bson.RegEx{Pattern: name, Options: "im"}}}).
		All(&result)
	if err != nil {
		return result, errors.New("not exist")
	}
	return result, nil
}

//通过学院模糊查找
func FormFindCollege(college string) ([]form, error) {
	var result []form
	err := FormId.
		Find(bson.M{"college": bson.M{"$regex": bson.RegEx{Pattern: college, Options: "im"}}}).
		All(&result)
	if err != nil {
		return result, errors.New("not exist")
	}
	return result, nil
}

//通过专业模糊查找
func FormFindMajor(major string) ([]form, error) {
	var result []form
	err := FormId.
		Find(bson.M{"major": bson.M{"$regex": bson.RegEx{Pattern: major, Options: "im"}}}).
		All(&result)
	if err != nil {
		return result, errors.New("not exist")
	}
	return result, nil
}

//数据去重复
func FormUnique(origin []form) []form {
	result := make([]form, 0)
	for _, value := range origin {
		//首先定义一个状态
		status := true
		for _, v := range result {
			if value.OpenId == v.OpenId {
				//如果有相同的值就改变状态值
				status = false
				break
			}
		}
		if status {
			result = append(result, value)
		}
	}
	return result
}

func FormClass_(authId string) (bool, error) {
	var result struct {
		Class_ bool `bson:"class_"`
	}
	m := []bson.M{
		{"$match": bson.M{"authId": authId}},
		{"$project": bson.M{"_id": 0, "class_": 1}},
	}
	err := FormId.
		Pipe(m).
		One(&result)
	// fmt.Printf("%+v", result)
	if err != nil {
		return false, errors.New("not exist")
	}
	return result.Class_, nil
}

func FormDel(openid string) bool {
	err := FormId.
		Remove(bson.M{"openId": openid})
	if err != nil {
		return false
	}
	return true
}

func FormExist(openid string) bool {
	var result []form
	err := FormId.
		Find(bson.M{"openId": openid}).All(&result)
	if err != nil || len(result) == 0 {
		return false
	}
	return true
}

//返回formId array 的长度
func FormSize(openid string) int {
	m := []bson.M{
		{"$match": bson.M{"openId": openid}},
		{"$project": bson.M{"_id": 0, "openId": 1, "count": bson.M{"$size": "$formId"}}},
	}
	var result formSize
	err := FormId.
		Pipe(m).
		One(&result)
	if err != nil {
		return 0
	}
	// fmt.Printf("%+v", result)
	return result.Count
}

func FormAdd(authid, openid, formid, nickName, name, college, major, avatarUrl string) (string, bool) {
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	formId := FormId
	err := formId.Update(bson.M{"openId": openid}, bson.M{"$push": bson.M{"formId": formid}, "$set": bson.M{"updateTime": updateTime, "nickName": nickName, "name": name, "college": college, "major": major, "avatarUrl": avatarUrl, "authId": authid}})
	if err != nil {
		return "fail", false
	}
	return "success", true
}

func FormInsert(authid, openid, formid, nickName, name, college, major, avatarUrl string) (string, bool) {
	now := int(time.Now().Unix())
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	formids := []string{formid}
	temp := form{
		OpenId:     openid,
		AuthId:     authid,
		FormId:     formids,
		UpdateTime: updateTime,
		CreateTime: now,
		NickName:   nickName,
		Name:       name,
		College:    college,
		Major:      major,
		AvatarUrl:  avatarUrl,
		Class_:     true,
	}
	err := FormId.
		Insert(temp)
	if err != nil {
		return "", false
	}
	return "success", true
}
