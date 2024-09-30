package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	h := healthz()
	writer := httptest.NewRecorder()
	h(writer, nil)
	response := writer.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("StatusCode is wrong. Expected:%d, Got:%d", http.StatusOK, response.StatusCode)
	}
}
