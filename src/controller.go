package src

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ErrorResponse ...
// Example: {errors: {name: ["can't be blank"], email: ["can't be blank", "already taken"]}}
type ErrorResponse struct {
	Errors map[string][]string `json:"errors"`
}

func writeErrorResponse(w http.ResponseWriter, err error, status int) {
	log.Println(err)

	resp := ErrorResponse{
		Errors: map[string][]string{
			"base": {err.Error()},
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Println("ErrorResponse json.Marshal Error:", err)
	}

	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, string(jsonResp))
}

// Controller ...
func Controller(w http.ResponseWriter, req *http.Request) {
	params := getParams(req)

	// Get data from the database
	records, err := params.ExecuteQuery()
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// Convert data to JSON
	recordsJSON, err := json.Marshal(records)
	if err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
	}

	// Write data on successful query and json Marshal
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(recordsJSON))
}
