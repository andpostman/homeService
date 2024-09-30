package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the webservice...")
	http.HandleFunc("/home", func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(writer, "Request was processed!")
	})
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
