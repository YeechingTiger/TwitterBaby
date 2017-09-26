package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"model"
)

func DBDrop() {
	session := DBConnect()
	err := session.DB("se_avengers").DropDatabase()
	if err != nil {
		panic(err)
	}

	session.Close()
}

func DBInsert() {
	session := DBConnect()

	userC := session.DB("se_avengers").C("users")

	users := []model.User {
		model.User{ID: bson.NewObjectId(), FirstName: "Jason", LastName: "Ho", Password: "test1", Email: "hojason117@gmail.com", Followers: nil, Followed: nil, 
			Bio: "Hi everyone, this is Jason Ho.", Token: "", UserIDdev: "JasonHo"},
		model.User{ID: bson.NewObjectId(), FirstName: "Chih-Yin", LastName: "Lee", Password: "test2", Email: "c788678867886@gmail.com", Followers: nil, Followed: nil, 
			Bio: "Hi everyone, this is Mars Lee.", Token: "", UserIDdev: "MarsLee"},
		model.User{ID: bson.NewObjectId(), FirstName: "Jason", LastName: "He", Password: "test3", Email: "hexing_h@hotmail.com", Followers: nil, Followed: nil, 
			Bio: "Hi everyone, this is Jason He.", Token: "", UserIDdev: "JasonHe"},
		model.User{ID: bson.NewObjectId(), FirstName: "Diane", LastName: "Lin", Password: "test4", Email: "diane@gmail.com", Followers: nil, Followed: nil, 
			Bio: "Hi everyone, this is Diane Lin.", Token: "", UserIDdev: "DianeLin"}, 
		model.User{ID: bson.NewObjectId(), FirstName: "Tom", LastName: "Riddle", Password: "test5", Email: "triddle@gmail.com", Followers: nil, Followed: nil, 
			Bio: "Hi everyone, this is Lord Voldemort.", Token: "", UserIDdev: "TomRiddle"}}
	
	for _, u:= range users {
		err := userC.Insert(u)
		if err != nil {
			panic(err)
		}
	}

	tweetC := session.DB("se_avengers").C("tweets")

	tweets := []model.Tweet {
		model.Tweet{ID: bson.NewObjectId(), From: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), From: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()}}

	for _, t := range tweets {
		err := tweetC.Insert(t)
		if err != nil {
			panic(err)
		}
	}

	session.Close()
}

func DBFind() {
	session := DBConnect()
	
	/*collect := session.DB("se_avangers").C("users")

	userIDdev := "JasonHo"
	var result model.User
	err := collect.Find(bson.M{"useriddev": userIDdev}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.ID)*/

	collect := session.DB("se_avengers").C("tweets")

	from := "JasonHo"
	var result model.User
	err := collect.Find(bson.M{"from": from}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.ID)

	session.Close()
}

func DBConnect() *mgo.Session {
	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		panic(err)
	}

	return session
}