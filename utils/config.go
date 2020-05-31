package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigParam struct {
	Port int64  `json:"port"`
	Host string `json:"host"`
	Key  string `json:"key"`
}

func ConfigExist(configPath string) bool {
	_, err := os.Stat(configPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
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

func ReplaceConfig(fileName string, conf ConfigParam) (bool, error) {
	if ConfigExist(fileName) {
		return writeConfig(fileName, conf)
	} else {
		return createConfig(fileName, conf)
	}
}

func createConfig(fileName string, conf ConfigParam) (bool, error) {
	err := os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		return false, fmt.Errorf("fail to create dir: %s", err)
	}

	f, err := os.Create(fileName)
	if err != nil {
		// 创建文件失败处理
		return false, err
	}
	defer f.Close()

	confBytes, _ := json.MarshalIndent(conf, "", "    ")

	_, err = f.Write(confBytes)
	if err != nil {
		return false, fmt.Errorf("write file err in create: %s", err)
	}

	return true, nil
}

func writeConfig(fileName string, conf ConfigParam) (bool, error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	}
	if f == nil {
		return false, err
	}
	defer f.Close()

	confBytes, _ := json.MarshalIndent(conf, "", "    ")
	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt(confBytes, n)
	return true, nil
}
