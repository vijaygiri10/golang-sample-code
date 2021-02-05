package src

import "net/http"

// SetHeaders ...
func (config ConfigStruct) SetHeaders(w http.ResponseWriter, req *http.Request) {
	// Set Content Type header to respond with a json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Set CORS Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", config.Service.Method)
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
