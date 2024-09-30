package handlers

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

var tests = []struct {
	name     string
	value    *atomic.Value
	expected int
}{
	{
		name:     "valid_test",
		value:    &atomic.Value{},
		expected: http.StatusOK,
	},
	{
		name:     "error_test",
		value:    nil,
		expected: http.StatusServiceUnavailable,
	},
}

func TestReadyz(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			if tt.name == "valid_test" {
				tt.value.Store(true)
			}
			r := readyz(tt.value)
			r(writer, nil)
			response := writer.Result()
			if response.StatusCode != tt.expected {
				t.Errorf("StatusCode is wrong. Expected:%d, Got:%d", http.StatusOK, response.StatusCode)
			}
		})
	}
}
