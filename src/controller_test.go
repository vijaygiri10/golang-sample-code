package src

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController(test *testing.T) {
	config := ReadConfigFile("./../config/development.json")

	req, err := http.NewRequest(config.Service.Method, config.Service.Path, nil)
	if err != nil {
		test.Errorf("http.NewRequest = %d; want nil", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(config.Router)
	handler.ServeHTTP(rec, req)

	// TEST: Should return a 200 OK
	if status := rec.Code; status != http.StatusOK {
		test.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var records []Record

	// TEST: Should return a valid JSON Array with Account Objects
	if err := json.Unmarshal(rec.Body.Bytes(), &records); err != nil {
		test.Errorf("Json unmarshal raised an error %v", err)
	}

	// TEST: Should return all the accounts for current user
	if len(records) != 5 {
		test.Errorf("Number of accounts found %d; should be equal to 5", len(records))
	}
}

func TestInvalidRequest(test *testing.T) {
	config := ReadConfigFile("./../config/development.json")

	req, err := http.NewRequest("POST", config.Service.Path, nil)
	if err != nil {
		test.Errorf("http.NewRequest = %d; want nil", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(config.Router)
	handler.ServeHTTP(rec, req)

	// TEST: Should return a 400 bad request error
	if status := rec.Code; status != http.StatusBadRequest {
		test.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestSetCORS(test *testing.T) {
	config := ReadConfigFile("./../config/development.json")

	req, err := http.NewRequest("OPTIONS", config.Service.Path, nil)
	if err != nil {
		test.Errorf("http.NewRequest = %d; want nil", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(config.Router)
	handler.ServeHTTP(rec, req)

	// TEST: Should return a 200 OK
	if status := rec.Code; status != http.StatusOK {
		test.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
