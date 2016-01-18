package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piLights/dioder"
)

func colorController(w http.ResponseWriter, r *http.Request) {
	color := mux.Vars(r)["color"]

	responseMessage := colorResponse{"Color successfully activated", color}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	fmt.Fprintf(w, message)

	switch color {
	case "red":
		dioder.SetAll(255, 0, 0)
		break
	case "green":
		dioder.SetAll(0, 255, 0)
		break
	case "blue":
		dioder.SetAll(0, 0, 255)
		break
	}
}
