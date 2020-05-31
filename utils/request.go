package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	contentType = "application/json; charset=utf-8"
)

type ResponseMessage struct {
	Code int64 `json:"code"`
	Data string `json:"data"`
	Message string `json:"message"`
}

// Get request
func Get(url string) (ResponseMessage, error) {
	// timeout：5s
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return ResponseMessage{}, err
	}
	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return ResponseMessage{}, err
		}
	}

	message := ResponseMessage{}
	err = json.Unmarshal(result.Bytes(), &message)
	if err != nil {
		return ResponseMessage{}, err
	}

	return message, nil
}

// Post request
func Post(url string, data interface{}) (ResponseMessage, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ResponseMessage{}, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	message := ResponseMessage{}
	err = json.Unmarshal(result, &message)
	if err != nil {
		return ResponseMessage{}, err
	}

	return message, nil
}
