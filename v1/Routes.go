package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MessageIndex",
		"GET",
		"/messages",
		MessageIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/messages",
		MessageCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/messages/{messageId}",
		MessageShow,
	},
}
