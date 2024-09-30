package main

import (
	"context"
	"homeService/handlers"
	"homeService/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Print("The service is ready to listen and serve.")
	killsignal := <-interrupt
	switch killsignal {
	case os.Interrupt:
		log.Print("GOT SIGINT...")
	case syscall.SIGTERM:
		log.Print("GOT SIGTERM...")
	}
	log.Print("The service is shutting down...")
	server.Shutdown(context.Background())
	log.Print("Done")
}
