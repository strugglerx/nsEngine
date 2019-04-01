package models

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

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
	err := Keywords.
		Insert(data)
	if err != nil {
		return false
	}
	return true
}

func KeywordRename(id, content string) error {
	err := Keywords.
		Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"content": content}})
	if err != nil {
		fmt.Println(err)
		return errors.New("fail")
	}
	return nil
}

func KeywordDelete(id string) bool {
	err := Keywords.
		Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return false
	}
	return true
}

func KeywordFind(keyword string) keywords {
	var result keywords
	err := Keywords.
		Find(bson.M{"keyword": bson.M{"$regex": keyword}}).
		One(&result)
	if err != nil {
		return keywords{}
	}
	return result
}

func KeywordList() []keywords {
	var result []keywords
	err := Keywords.
		Find(bson.M{}).
		Sort("-_id").
		All(&result)
	if err != nil {
		return nil
	}
	return result
}
