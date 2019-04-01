package models

import (
	"gopkg.in/mgo.v2/bson"
)

type StepUser struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	AvatarUrl string        `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	Nickname  string        `json:"nickName" bson:"nickName,omitempty"`
	Step      interface{}   `json:"step,omitempty"`
}

func StepList() []StepUser {
	result := []StepUser{}
	Run.
		Find(bson.M{}).
		Select(bson.M{"_id": 0}).
		Sort("-step").
		All(&result)
	return result
}

func StepFindOne(nickName string) bool {
	db := Run.
		Find(bson.M{"nickName": nickName}).
		Select(bson.M{"_id": 0})
	result := []StepUser{}
	db.All(&result)
	if len(result) == 0 {
		return false
	}
	return true
}

func StepUpdate(avatarUrl, nickName string, step int64) bool {
	err := Run.
		Update(bson.M{"nickName": nickName},
			bson.M{"$set": bson.M{"avatarUrl": avatarUrl, "step": step}})
	if err != nil {
		return false
	}
	return true
}

func StepDelete(id string) bool {
	err := Run.
		Remove(bson.M{"_id": id})
	if err != nil {
		return false
	}
	return true
}

func StepInsert(avatarUrl, nickName string, step int64) bool {
	temp := StepUser{
		AvatarUrl: avatarUrl,
		Nickname:  nickName,
		Step:      step,
	}
	err := Run.
		Insert(temp)
	if err != nil {
		return false
	}
	return true
}
