package models

import (
	"gopkg.in/mgo.v2/bson"
	"server/models/mymongo"
)

type advertisment struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	BgUrl  string        `json:"backgroundUrl,omitempty"`
	ID string        `json:"ID" bson:"ID,omitempty"`
	DateStart string `json:"dateStart,omitempty"`
	DateEnd string `json:"dateEnd,omitempty"`
	Remark string `json:"remark,omitempty"`
}

func AdList() []advertisment {
	database:=mymongo.GetDataBase()
	db:=database.C("advertisment").
		Find(bson.M{}).
		Sort("-datestart").
		Select(bson.M{"_id":0})
	result := []advertisment{}
	db.All(&result)
	return result
}

func AdListLimit() []advertisment {
	database:=mymongo.GetDataBase()
	db:=database.C("advertisment").
		Find(bson.M{}).
		Sort("-datestart").
		Limit(3).
		Select(bson.M{"_id":0})

	result := []advertisment{}
	db.All(&result)
	return result
}



func AdDel(id string) bool {
	database:=mymongo.GetDataBase()
	err:=database.C("advertisment").Remove(bson.M{"ID":id})
	if err!=nil{
		return  false
	}
	return true
}

func AdInsert(id,bgUrl,dateStart,dateEnd,remark string) (bool) {
	temp:=advertisment{
		ID:id,
		BgUrl:bgUrl,
		DateStart:dateStart,
		DateEnd:dateEnd,
		Remark:remark,
	}
	database:=mymongo.GetDataBase()
	err:=database.C("advertisment").
		Insert(temp)
	if err!=nil{
		return false
	}
	return true
}
