package main

import "time"

type Message struct {
	Id        int       `json:"id"`
	Text      string 		`json:"text"`
	EmailAddress      string    `json:"email"`
	Time       time.Time `json:"time"`

}

type Messages []Message
