package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigParam struct {
	Port int64  `json:"port"`
	Host string `json:"host"`
	Key  string `json:"key"`
}

func LoadConfig(filepath string) (ConfigParam, error) {
	filePtr, err := os.Open(filepath)
	if err != nil {
		return ConfigParam{}, fmt.Errorf("open file failed [Err :%s]", err.Error())
	}
	defer filePtr.Close()

	var person ConfigParam

	// Create json decoder
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&person)
	if err != nil {
		return ConfigParam{}, fmt.Errorf("create decoder failed, err :%s", err.Error())
	}

	return person, nil
}
