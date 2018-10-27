package models

import (
	"gopkg.in/mgo.v2/bson"
	"server/models/mymongo"
)

type indexArticle struct {
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ID string        `json:"ID" bson:"ID,omitempty"`
	Title string        `json:"title,omitempty"`
	Path string        `json:"path"bson:"path,omitempty"`
	Url string        `json:"url,omitempty"`
	Date string        `json:"date,omitempty"`
	Author string        `json:"author,omitempty"`
}
//模块配置列表
func IAList() []indexArticle {
	database:=mymongo.GetDataBase()
	db:=database.C("header_options").Find(bson.M{}).Select(bson.M{"_id":0})
	result := []indexArticle{}
	db.All(&result)
	return result
}

//增加
func IAInsert(path,url,title,id,date,author string) bool {
	temp:=indexArticle{
		Path:path,
		Url:url,
		Title:title,
		ID:id,
		Date:date,
		Author:author,
	}
	database:=mymongo.GetDataBase()
	err:=database.C("header_options").
		Insert(temp)
	if err!=nil{
		return  false
	}
	return true
}

//删除
func IADel(id string) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("header_options").
		Remove(bson.M{"ID":id})
	if err!=nil{
		return  false
	}
	return true
}

