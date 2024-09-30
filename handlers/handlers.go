package handlers

import "github.com/gorilla/mux"

func Route() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home)
	return r
}
