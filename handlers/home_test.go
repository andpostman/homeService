package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expectedBody = "Request was processed!"

func TestHome(t *testing.T) {
	writer := httptest.NewRecorder()
	home(writer, nil)
	response := writer.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("StatusCode is wrong. Expected:%d, Got:%d", http.StatusOK, response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if string(body) != expectedBody {
		t.Errorf("Body content is wrong. Expected:%s, Got:%s", expectedBody, string(body))
	}
}
