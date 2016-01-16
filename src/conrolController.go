package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piLights/dioder"
)

func controlController(w http.ResponseWriter, r *http.Request) {
	state := mux.Vars(r)["state"]

	if state == "off" {
		dioder.SetAll(0, 0, 0)
	}
}
