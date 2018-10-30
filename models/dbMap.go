package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"server/models/mymongo"
)
//中心点
type CenterPotint struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Point []interface{} `json:"point" bson:"point,omitempty"`
	Type_ string `json:"type" bson:"type,omitempty" `
}

//标记点
type MapSign struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Point []interface{} `json:"point" bson:"point,omitempty"`
	Type_ string `json:"type,omitempty" bson:"type,omitempty" `
	ID int `json:"id,omitempty" bson:"id,omitempty" `
}

//中心点
func PointList(type_ string) []CenterPotint {
	database:=mymongo.GetDataBase()
	db:=database.C("centerPoint").
		Find(bson.M{"type":type_}).
		Select(bson.M{"_id":0})
	result := []CenterPotint{}
	db.All(&result)
	return result
}
//-----------------------------

func SignList(type_ string) []MapSign {
	database:=mymongo.GetDataBase()
	db:=database.C("mapSign").
		Find(bson.M{"type":type_}).
		Select(bson.M{"_id":0,"id":0,"type":0}).
		Sort("id")
	result := []MapSign{}
	db.All(&result)
	return result
}



func SignSet(type_, content,longitude, newName string,id int) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("mapSign").
		Update(bson.M{"id":id,"type":type_,"point.content":content,"point.longitude":longitude},
		bson.M{"$set":bson.M{"point.$.content":newName}})
	if err!=nil{
		return  false
	}
	return true
}

//"content" : "音乐学院", "latitude" : 40.802803, "longitude" : 111.69801

func SignPull(type_, content ,longitude string,id int  ) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("mapSign").
		Update(bson.M{"id":id,"type":type_},
			bson.M{"$pull":bson.M{"point":bson.M{"content" : content, "longitude" : longitude}}})
	fmt.Println(err)
	if err!=nil{
		return  false
	}
	return true
}

func SignPush(type_,content ,latitude ,longitude string,id int  ) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("mapSign").
		Update(bson.M{"id":id,"type":type_},
			bson.M{"$push":bson.M{"point":bson.M{"content" : content, "latitude" : latitude, "longitude" : longitude}}})
	if err!=nil{
		fmt.Println(err)
		return  false
	}
	return true
}

