package models

import (
	"gopkg.in/mgo.v2/bson"
)

type WebUser struct {
	Id     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User   string        `json:"user"`
	Passwd string        `json:"passwd"`
	Role   int           `json:"role"`
}

//用户验证
func UserVerify(user, passwd string) []WebUser {
	result := []WebUser{}
	StrUser.
		Find(bson.M{"user": user, "passwd": passwd}).
		Select(bson.M{"_id": 0}).
		All(&result)

	return result
}

//用户信息登录
func UserUpdate(user, passwd string) bool {
	err := StrUser.
		Update(bson.M{"user": user},
			bson.M{"$set": bson.M{"passwd": passwd}})
	if err != nil {
		return false
	}
	return true
}
