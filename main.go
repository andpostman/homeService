package main

import (
	"homeService/handlers"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the webservice...")
	router := handlers.Route()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
