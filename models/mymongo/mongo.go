package mymongo

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var database *mgo.Database

func GetMgo() *mgo.Session {
	return session.Copy()
}
func CloseMgo() {
	session.Close()
}

func GetDataBase() *mgo.Database {
	return database
}

func init() {

	dial := beego.AppConfig.String("mongodb::url")
	session, err := mgo.Dial(dial)
	if err != nil {
		// 连接数据库
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	database = session.DB("struggler")
	//defer session.Clone()
}
