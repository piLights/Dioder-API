package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piLights/dioder"
)

func rgbColorController(w http.ResponseWriter, r *http.Request) {
	redValue := mux.Vars(r)["red"]
	greenValue := mux.Vars(r)["green"]
	blueValue := mux.Vars(r)["blue"]

	responseColorString := redValue + ":" + greenValue + ":" + blueValue
	responseMessage := colorResponse{"Color successfully activated", responseColorString}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	redUint, _ := strconv.Atoi(redValue)
	greenUint, _ := strconv.Atoi(greenValue)
	blueUint, _ := strconv.Atoi(blueValue)

	dioder.SetAll(uint8(redUint), uint8(greenUint), uint8(blueUint))
	fmt.Fprintf(w, message)
}
