package main

import (
	mgo "github.com/AndrewUdell/GroupMail/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

var (
	session  *mgo.Session
	database *mgo.Database
)

const (
	DatabaseName = "GroupMailDatabase"
)

// Singleton Pattern: get existing session (or if it doesn't exist, start a new session)
func singletonSession() (*mgo.Session, error) {
	if session == nil {
		session, err := mgo.Dial("127.0.0.1")
		return session, err
	} else {
		return session, nil
	}
}

// Single Database: get existing database (or create new one) if session exists
func SingletonDatabase() (*mgo.Database, error) {
	if database == nil {
		session, err := singletonSession()
		if err != nil {
			return nil, err
		} else {
			return session.DB(DatabaseName), nil
		}
	} else {
		return database, nil
	}
}
