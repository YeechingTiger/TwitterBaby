package model

import "gopkg.in/mgo.v2/bson"

// User : Data structure that holds a single user.
type User struct {
	ID        	bson.ObjectId 	`json:"id" bson:"_id"`
	FirstName 	string			`json:"firstname" bson:"firstname"`
	LastName	string			`json:"lastname,omitempty" bson:"lastname,omitempty"`
	Password 	string			`json:"password,omitempty" bson:"password"`
	Email 		string			`json:"email" bson:"email"`
	Followers	[]string		`json:"followers,omitempty" bson:"followers,omitempty"`
	Followed	[]string		`json:"followed,omitempty" bson:"followed,omitempty"`
	Bio			string			`json:"bio,omitempty" bson:"bio,omitempty"`
	Token		string			`json:"token,omitempty" bson:"-"`
	
	UserIDdev	string			`json:"useriddev,omitempty" bson:"useriddev,omitempty"`
}
