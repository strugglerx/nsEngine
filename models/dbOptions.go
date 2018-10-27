package models

import (
	"gopkg.in/mgo.v2/bson"
	"server/models/mymongo"
)

type indexOption struct {
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty"`
	ID int        `json:"ID" bson:"ID,omitempty"`
	Option_id int        `json:"option_id,omitempty"`
	Status int        `json:"status" bson:"status,omitempty"`
}
//模块配置列表
func OptList() []indexOption {
	database:=mymongo.GetDataBase()
	db:=database.C("mini_options").
		Find(bson.M{}).
		Sort("option_id").
		Select(bson.M{"_id":0})
	result := []indexOption{}
	db.All(&result)
	return result
}
//开关状态更改
func OptUpConf(id,status int) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("mini_options").Update(bson.M{"option_id":id},
		bson.M{"$set":bson.M{"status":status}})
	if err!=nil{
		return  false
	}
	return true
}

//增加配置
func OptInsert(name string, id int) bool {
	temp:=indexOption{
		Name:name,
		Option_id:id,
		ID:id,
		Status:1,
	}
	database:=mymongo.GetDataBase()
	err:=database.C("mini_options").
		Insert(temp)
	if err!=nil{
		return  false
	}
	return true
}
