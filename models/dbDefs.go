package models

import (
	"server/models/mymongo"
)

var DB = mymongo.GetDataBase()

var FormId = DB.C("formId")

var Article = DB.C("articles")

var FeedBack = DB.C("feedback")

var Advertisment = DB.C("advertisment")

var HeaderOptions = DB.C("header_options")

var Jobs = DB.C("jobs")

var Keywords = DB.C("keywords")

var CenterPoint = DB.C("centerPoint")

var MiniOptions = DB.C("mini_options")

var Run = DB.C("run")

var StrUser = DB.C("str_user")

var MapSigns = DB.C("mapSign")
