package main

import (
	"fmt"
	"net/http"

	"code.google.com/p/sadbox/color"
	"github.com/gorilla/mux"
	"github.com/piLights/dioder"
)

func hexColorController(w http.ResponseWriter, r *http.Request) {
	hexColor := mux.Vars(r)["hex"]

	responseMessage := colorResponse{"Color successfully activated", hexColor}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	dioder.SetAll(color.HexToRGB(color.Hex(hexColor)))
	fmt.Fprintf(w, message)
}
