package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type advertisment struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	BgUrl     string        `json:"backgroundUrl,omitempty"`
	ID        string        `json:"ID" bson:"ID,omitempty"`
	DateStart string        `json:"dateStart,omitempty"`
	DateEnd   string        `json:"dateEnd,omitempty"`
	Remark    string        `json:"remark,omitempty"`
}

func AdList() []advertisment {
	result := []advertisment{}
	Advertisment.
		Find(bson.M{}).
		Sort("-datestart").
		All(&result)
	return result
}

func AdRename(id, remark string) error {
	err := Advertisment.
		Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"remark": remark}})
	if err != nil {
		return errors.New("fail")
	}
	return nil
}

func AdListLimit() []advertisment {
	db := Advertisment.
		Find(bson.M{}).
		Sort("-datestart").
		Limit(3).
		Select(bson.M{"_id": 0})

	result := []advertisment{}
	db.All(&result)
	return result
}

func AdDel(id string) bool {
	err := Advertisment.Remove(bson.M{"ID": id})
	if err != nil {
		return false
	}
	return true
}

func AdInsert(id, bgUrl, dateStart, dateEnd, remark string) bool {
	temp := advertisment{
		ID:        id,
		BgUrl:     bgUrl,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		Remark:    remark,
	}
	err := Advertisment.
		Insert(temp)
	if err != nil {
		return false
	}
	return true
}
