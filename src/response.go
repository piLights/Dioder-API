package main

import (
	"encoding/json"
)

func GenerateResponseMessage(v interface{}) (string, error) {
	message, error := json.Marshal(v)

	if error != nil {
		return "", error
	}

	return string(message), nil
}
