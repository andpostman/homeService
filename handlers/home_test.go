package handlers

import (
	"encoding/json"
	"homeService/version"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHome(t *testing.T) {
	writer := httptest.NewRecorder()
	buildTime := time.Now().Format("20060102_03:04:05")
	commit := "test hash"
	release := "0.0.1"
	h := home(buildTime, commit, release)
	h(writer, nil)

	response := writer.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("StatusCode is wrong. Expected:%d, Got:%d", http.StatusOK, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		t.Error(err)
	}
	info := struct {
		BuildTime string `json:"build_time"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		t.Error(err)
	}
	if info.BuildTime != version.BuildTime {
		t.Errorf("BuildTime is wrong. Expected:%s, Got:%s", version.BuildTime, info.BuildTime)
	}
	if info.Commit != version.Commit {
		t.Errorf("Commit is wrong. Expected:%s, Got:%s", version.Commit, info.Commit)
	}
	if info.Release != version.Release {
		t.Errorf("Release is wrong. Expected:%s, Got:%s", version.Release, info.Release)
	}
}
