package main

import (
	"homeService/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting the webservice...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is not set")
	}
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
