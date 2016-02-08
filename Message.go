package main

import (
	"time"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func FindAllMessages() error {
	var messages []Messages
	err := getMessageCollection().find().Sort("-time").All(&results)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("All Messages: ", &results)
	}
}
