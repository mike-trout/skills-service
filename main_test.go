// main_test.go

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllSkills(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/skills", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	body := strings.TrimSpace(resp.Body.String())
	jsonBytes, _ := json.Marshal(Skills)
	json := strings.TrimSpace(string(jsonBytes))
	if body != json {
		t.Errorf("Expected %s. Got %s", body, json)
	}
}

func TestGetSingleSkill(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/skills/1", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	body := strings.TrimSpace(resp.Body.String())
	jsonBytes, _ := json.Marshal(Skills[0])
	json := strings.TrimSpace(string(jsonBytes))
	if body != json {
		t.Fatalf("Expected %s. Got %s", body, json)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a := App{}
	a.Initialise()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
