package models

import (
	"gopkg.in/mgo.v2/bson"
	"server/models/mymongo"
)

type WebUser struct {
	Id    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User string      `json:"user"`
	Passwd string    `json:"passwd"`
	Role int      `json:"role"`
}
//用户验证
func UserVerify(user, passwd string) []WebUser {
	database:=mymongo.GetDataBase()
	db:=database.C("str_user").
		Find(bson.M{"user":user,"passwd":passwd}).
		Select(bson.M{"_id":0})
	result := []WebUser{}
	db.All(&result)
	return result
}
//用户信息登录
func UserUpdate(user, passwd string) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("str_user").
		Update(bson.M{"user":user},
			bson.M{"$set":bson.M{"passwd":passwd}})
	if err!=nil{
		return  false
	}
	return true
}

