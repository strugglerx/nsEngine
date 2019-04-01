package models

import (
	"gopkg.in/mgo.v2/bson"
)

type feedback struct {
	Id         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Openid     string        `json:"openid" bson:"openid,omitempty"`
	Content    string        `json:"content" bson:"content,omitempty" `
	CreateTime int           `json:"createTime" bson:"createTime,omitempty" `
	NickName   string        `json:"nickName" bson:"nickName,omitempty"`
	Name       string        `json:"name" bson:"name,omitempty"`
	AvatarUrl  string        `json:"avatarUrl" bson:"avatarUrl,omitempty"`
	College    string        `json:"college" bson:"college,omitempty"`
}

func FeedbackList() []feedback {
	result := []feedback{}
	FeedBack.
		Find(bson.M{}).
		Sort("-createTime").
		All(&result)
	return result
}

func FeedBackDelete(_id string) bool {
	err := FeedBack.RemoveId(bson.ObjectIdHex(_id))
	if err != nil {
		return false
	}
	return true
}

func FeedbackInsert(nickName, name, avatarUrl, college, openid, content string, createTime int) bool {
	temp := feedback{
		Openid:     openid,
		Content:    content,
		CreateTime: createTime,
		College:    college,
		NickName:   nickName,
		Name:       name,
		AvatarUrl:  avatarUrl,
	}
	err := FeedBack.
		Insert(temp)
	if err != nil {
		return false
	}
	return true
}
