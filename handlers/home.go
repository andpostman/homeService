package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(buildTime, commit, release string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		info := struct {
			BuildTime string `json:"build_time"`
			Commit    string `json:"commit"`
			Release   string `json:"release"`
		}{
			buildTime,
			commit,
			release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Couldn't encode info data: %v", err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(body)
	}
}
