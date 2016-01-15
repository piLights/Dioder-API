package main

import (
		"log"
		"net/http"
		"fmt"
)

func main() {
	fmt.Printf("Starting Dioder-API...\n")

	fmt.Printf("Registering available routes\n")
	colorAPI := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", colorAPI))
}
