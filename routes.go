package main

import (
	"net/http"
	user "populi/user"
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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserGET",
		"GET",
		"/user/{id}",
		user.Get,
	},
	Route{
		"UserCREATE",
		"POST",
		"/user/create",
		user.Create,
	},
	Route{
		"UserUPDATE",
		"PUT",
		"/user/{id}",
		user.Update,
	},
	Route{
		"UserDELETE",
		"DELETE",
		"/user/{id}",
		user.Delete,
	},
}
