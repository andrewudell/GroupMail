package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	DatabaseInstance, _ = SingletonDatabase()
)

const (
	UserCollection = "Users"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	EmailAddress      string    `json:"email"`
	Name      string    `json:"name"`
}

type Users []User

func CreateUser(user User) error {
	id := bson.NewObjectId()
	user["ID"] = id
	return getUserCollection().Insert(user)
}

// Collection (Table) of users
func getUserCollection() *mgo.Collection {
	return DatabaseInstance.C(UserCollection)
}
