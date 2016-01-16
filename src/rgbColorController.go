package main

/*import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piLights/dioder"
)

func rgbColorController(w http.ResponseWriter, r *http.Request) {
	redValue := mux.Vars(r)["red"]
	greenValue := mux.Vars(r)["green"]
	blueValue := mux.Vars(r)["blue"]

	responseColorString := r + ":" + g + ":" + b
	responseMessage := colorResponse{"Color successfully activated", responseColorString}

	message, error := GenerateResponseMessage(responseMessage)

	if error != nil {
		fmt.Fprintf(w, "There was an error")

		return
	}

	dioder.SetAll(uint8(redValue), uint8(greenValue), uint8(blueValue))
	fmt.Fprintf(w, message)
}
*/
