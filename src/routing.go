package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

type Route struct {
	Name	string
	Method	string
	Pattern	string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		fmt.Printf("Registering route %s (%s)", route.Name, route.Pattern)
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logging(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
		fmt.Printf(" - OK\n")
	}

	return router
}

var routes = Routes {
	Route {
		"Color",
		"GET",
		"/color/{color}",
		ColorController,
	},
	/*Route {
		"Hex",
		"GET",
		"/hex/{hex}",
		HexColorController,
	},
	Route {
		"RGB",
		"GET",
		"/rgb/{rgb}",
		RGBColorController,
	},*/
}
