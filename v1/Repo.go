package main

import "fmt"

var currentId int

var messages Messages

// Give us some seed data
func init() {
	RepoCreateMessage(Message{Text: "Amazing article here: nytimes.com"})
}

func RepoFindMessage(id int) Message {
	for _, m := range messages {
		if m.Id == id {
			return m
		}
	}
	// return empty Message if not found
	return Message{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateMessage(t Message) Message {
	currentId += 1
	m.Id = currentId
	messages = append(messages, m)
	return m
}

func RepoDestroyMessage(id int) error {
	for i, m := range messages {
		if m.Id == id {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Message with id of %d to delete", id)
}
