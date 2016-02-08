package main

import (
	"time"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var (
	DatabaseInstance, _ = SingletonDatabase()
)

const (
	MessageCollection = "Messages"
)

type Message struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Text      string 		`json:"text"`
	EmailAddress      string    `json:"email"`
	Time       time.Time `json:"time"`
}

type Messages []Message

func CreateMessage(message Message) error {
	id := bson.NewObjectId()
	message["ID"] = id
	return getMessageCollection().Insert(message)
}

// Collection (Table) of messages
func getMessageCollection() *mgo.Collection {
	return DatabaseInstance.C(MessageCollection)
}

func FindAllMessages() {
	var messages []Messages
	err := getMessageCollection().find().Sort("-time").All(&messages)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("All Messages: ", &messages)
	}
}
