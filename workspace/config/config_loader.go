package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"workspace/types"
)

func LoadConfig(filename string) (*types.Config, error) {
	// Read configuration file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse configuration JSON into Config struct
	var config types.Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	fmt.Println("config is: ", config)
	return &config, nil
}
