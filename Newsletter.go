package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	DatabaseInstance, _ = SingletonDatabase()
)

const (
	NewsletterCollection = "Newsletter"
)

type Newsletter struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name      string    `json:"name"`
}

type Newsletters []Newsletter

func CreateNewsletter(newsletter Newsletter) error {
	id := bson.NewObjectId()
	newsletter["ID"] = id
	return getNewsletterCollection().Insert(newsletter)
}

// Collection (Table) of newsletters
func getNewsletterCollection() *mgo.Collection {
	return DatabaseInstance.C(NewsletterCollection)
}
