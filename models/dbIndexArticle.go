package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type indexArticle struct {
	Id     bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ID     string        `json:"ID" bson:"ID,omitempty"`
	Title  string        `json:"title,omitempty"`
	Path   string        `json:"path" bson:"path,omitempty"`
	Url    string        `json:"url,omitempty"`
	Date   string        `json:"date,omitempty"`
	Author string        `json:"author,omitempty"`
}

//模块配置列表
func IAList() []indexArticle {
	result := []indexArticle{}
	HeaderOptions.
		Find(bson.M{}).
		Select(bson.M{"_id": 0}).
		All(&result)
	return result
}

func IARename(id, title, author string) error {
	err := HeaderOptions.
		Update(bson.M{"ID": id}, bson.M{"$set": bson.M{"title": title, "author": author}})
	if err != nil {
		return errors.New("fail")
	}
	return nil
}

//增加
func IAInsert(path, url, title, id, date, author string) bool {
	temp := indexArticle{
		Path:   path,
		Url:    url,
		Title:  title,
		ID:     id,
		Date:   date,
		Author: author,
	}
	err := HeaderOptions.
		Insert(temp)
	if err != nil {
		return false
	}
	return true
}

//删除
func IADel(id string) bool {
	err := HeaderOptions.
		Remove(bson.M{"ID": id})
	if err != nil {
		return false
	}
	return true
}
