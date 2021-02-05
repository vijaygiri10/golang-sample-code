// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"service/src"
)

func main() {
	configFilePath := "./config/development.json"
	if len(os.Args) > 1 {
		configFilePath = os.Args[1]
	}

	// Read configuration from the file path
	config := src.ReadConfigFile(configFilePath)

	// Create a new Spanner client
	src.NewSpannerClient(config)
	defer src.DBClient.Client.Close()

	http.HandleFunc(config.Service.Path, config.Router) // Load all the routes

	fmt.Println("Server is up:", config)
	log.Fatal(http.ListenAndServe(config.Service.Host+":"+config.Service.Port, nil)) // Start the API server
}
