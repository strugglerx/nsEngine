package models

import (
	"server/models/mymongo"

	"gopkg.in/mgo.v2/bson"
)

const collections="keywords"

type keywords struct {
	Id      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string        `json:"content,omitempty"`
	Keyword string        `json:"keyword" bson:"keyword,omitempty"`
}

func KeywordInsert(content, keyword string) bool {
	data := keywords{
		Keyword: keyword,
		Content: content,
	}
	database := mymongo.GetDataBase()
	err := database.C(collections).Insert(data)
	if err != nil {
		return false
	}
	return true
}

func KeywordDelete(id string) bool {
	database := mymongo.GetDataBase()
	err := database.
		C(collections).
		Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return false
	}
	return true
}

func KeywordFind(keyword string) keywords {
	database := mymongo.GetDataBase()
	var result keywords
	err := database.
		C(collections).
		Find(bson.M{"keyword": bson.M{"$regex": keyword}}).
		One(&result)
	if err != nil {
		return keywords{}
	}
	return result
}

func KeywordList() []keywords {
	database := mymongo.GetDataBase()
	var result []keywords
	err := database.
		C(collections).
		Find(bson.M{}).
		Sort("-_id").
		All(&result)
	if err != nil {
		return nil
	}
	return result
}
