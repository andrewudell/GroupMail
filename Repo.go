package main

import "fmt"

var currentId int

var messages Messages

// Give us some seed data
func init() {
	RepoCreateMessage(Message{Text: "Amazing article here: nytimes.com"})
}

func RepoFindMessage(id int) Message {
	for _, t := range messages {
		if t.Id == id {
			return t
		}
	}
	// return empty Message if not found
	return Message{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateMessage(t Message) Message {
	currentId += 1
	t.Id = currentId
	messages = append(messages, t)
	return t
}

func RepoDestroyMessage(id int) error {
	for i, t := range messages {
		if t.Id == id {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Message with id of %d to delete", id)
}
