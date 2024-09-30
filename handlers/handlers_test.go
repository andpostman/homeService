package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := Router("", "", "")
	server := httptest.NewServer(r)
	defer server.Close()

	response, err := http.Get(server.URL + "/home")
	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("StatusCode for path /home wrong. Expected:%d, Got:%d", http.StatusOK, response.StatusCode)
	}
}
