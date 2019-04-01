package models

import (
	"gopkg.in/mgo.v2/bson"
)

type indexOption struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty"`
	ID        int           `json:"ID" bson:"ID,omitempty"`
	Option_id int           `json:"option_id,omitempty"`
	Status    int           `json:"status" bson:"status,omitempty"`
}

//模块配置列表
func OptList() []indexOption {
	result := []indexOption{}
	MiniOptions.
		Find(bson.M{}).
		Sort("option_id").
		Select(bson.M{"_id": 0}).
		All(&result)

	return result
}

//开关状态更改
func OptUpConf(id, status int) bool {
	err := MiniOptions.
		Update(bson.M{"option_id": id},
			bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return false
	}
	return true
}

//增加配置
func OptInsert(name string, id int) bool {
	temp := indexOption{
		Name:      name,
		Option_id: id,
		ID:        id,
		Status:    1,
	}
	err := MiniOptions.
		Insert(temp)
	if err != nil {
		return false
	}
	return true
}
