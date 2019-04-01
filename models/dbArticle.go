package models

import (
	"errors"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type article struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string        `json:"title,omitempty"`
	ID       string        `json:"ID" bson:"ID,omitempty"`
	Date     string        `json:"date,omitempty"`
	Author   string        `json:"author,omitempty"`
	Content  string        `json:"content,omitempty"`
	LikeUser []string      `json:"likeUser,omitempty" bson:"likeUser,omitempty"`
	Like     int           `json:"like,omitempty"`
	View     int           `json:"view,omitempty"`
}

func ArtList() []article {
	result := []article{}
	Article.
		Find(bson.M{}).
		Select(bson.M{"content": 0, "_id": 0, "likeUser": 0}).
		Sort("-date").
		All(&result)
	return result
}

func ArtRename(id, title, author, content string) error {
	err := Article.
		Update(bson.M{"ID": id}, bson.M{"$set": bson.M{"title": title, "author": author, "content": content}})
	if err != nil {
		return errors.New("fail")
	}
	return nil
}

func ArtDetail(id string) []article {
	db := Article.Find(bson.M{"ID": id}).Select(bson.M{"_id": 0, "likeUser": 0})
	result := []article{}
	db.All(&result)
	return result
}

func ArtUpView(id string) bool {
	err := Article.Update(bson.M{"ID": id},
		bson.M{"$inc": bson.M{"view": 1}})
	if err != nil {
		return false
	}
	return true
}

func ArtUpLike(id string, num int) bool {
	err := Article.Update(bson.M{"ID": id},
		bson.M{"$inc": bson.M{"like": num}})
	if err != nil {
		return false
	}
	return true
}
func ArtFindLike(id string, name string) bool {
	db := Article.Find(bson.M{"ID": id, "likeUser": name})
	result := []article{}

	db.All(&result)
	if len(result) == 0 {
		err := Article.Update(bson.M{"ID": id},
			bson.M{"$push": bson.M{"likeUser": name}})
		if err != nil {
			return false
		}
		return true
	} else {
		err := Article.Update(bson.M{"ID": id},
			bson.M{"$pull": bson.M{"likeUser": name}})
		if err != nil {
			return false
		}
	}
	return false
}

func ArtDel(id string) bool {
	err := Article.Remove(bson.M{"ID": id})
	if err != nil {
		return false
	}
	return true
}

func ArtInsert(title, author, content, date string) (string, bool) {
	//生成uuid
	id, _ := uuid.NewV4()
	temp := article{
		ID:      id.String(),
		Title:   title,
		Date:    date,
		Author:  author,
		Content: content,
		View:    0,
		Like:    0,
	}
	err := Article.
		Insert(temp)
	if err != nil {
		return "", false
	}
	return id.String(), true
}
