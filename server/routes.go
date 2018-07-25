package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AppRoute struct {
	Name            string
	Method          string
	Pattern         string
	FunctionHandler http.HandlerFunc
}

type AppRoutes []AppRoute

func NewAppRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range appRoutes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.FunctionHandler)

	}
	return router
}

var appRoutes = AppRoutes{
	AppRoute{
		"UserList",
		"GET",
		"/users",
		ListUsers,
	},
	AppRoute{
		"UserDetails",
		"GET",
		"/users/{id}",
		GetUser,
	},
}
