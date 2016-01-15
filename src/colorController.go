package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type colorResponse struct {
	Message	string
	Color	string
}

func ColorController(w http.ResponseWriter, r *http.Request) {
	color := mux.Vars(r)["color"]

	responseMessage := colorResponse {"Color successfully activated",color}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	fmt.Fprintf(w, message)

	switch color {
		case "red":
			SetRed(255)
		break;
		case "green":
			SetGreen(255)
		break;
		case "blue":
			SetBlue(255)
		break;

	}
}
