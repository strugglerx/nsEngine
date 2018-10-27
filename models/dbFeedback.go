package models

import (
	"gopkg.in/mgo.v2/bson"
	"server/models/mymongo"
)

type feedback struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Openid string `json:"openid" bson:"openid,omitempty"`
	Content string `json:"content" bson:"content,omitempty" `
	CreateTime string `json:"createTime" bson:"createTime,omitempty" `
}


func FeedbackList() []feedback {
	database:=mymongo.GetDataBase()
	db:=database.C("feedback").
		Find(bson.M{}).
		Select(bson.M{"_id":0}).
		Sort("-createTime")
	result := []feedback{}
	db.All(&result)
	return result
}

func FeedbackInsert(openid,content,createTime string) bool {
	database:=mymongo.GetDataBase()
	temp:=feedback{
		Openid:openid,
	    Content:content,
		CreateTime:createTime,
	}
	err:=database.C("feedback").
		Insert(temp)
	if err!=nil{
		return false
	}
	return true
}