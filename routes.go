package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"TodoCreate",
		"POST",
		"/v1/todos",
		TodoCreate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/v1/todos/{todoId}",
		TodoDelete,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/v1/todos",
		GetTodos,
	},
	Route{
		"TodoShow",
		"GET",
		"/v1/todos/{todoId}",
		TodoShow,
	},
}
