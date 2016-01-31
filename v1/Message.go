package main

import "time"

type Message struct {
	Id        int       `json:"id"`
	Text      string 		`json:"text"`
	EmailAddress      string    `json:"email"`
}

type Messages []Message
