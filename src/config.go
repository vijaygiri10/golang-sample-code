package src

import (
	"encoding/json"
	"fmt"
	"os"
)

// ConfigStruct ...
type ConfigStruct struct {
	Spanner struct {
		Project  string `json:"project"`
		Instance string `json:"instance"`
		Database string `json:"database"`
	} `json:"spanner"`
	Service struct {
		Host   string `json:"host"`
		Port   string `json:"port"`
		Path   string `json:"path"`
		Method string `json:"Method"`
	} `json:"service"`
}

// ReadConfigFile ...
func ReadConfigFile(file string) ConfigStruct {
	var config ConfigStruct
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
