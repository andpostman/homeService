package main

import (
	"homeService/handlers"
	"homeService/version"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Starting the webservice...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is not set")
	}
	router := handlers.Router(version.BuildTime, version.Commit, version.Release)
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
