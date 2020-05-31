package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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
func Post(url string, data *url.Values) (ResponseMessage, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return ResponseMessage{}, err
	}
	defer request.Body.Close()

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3100.0 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		return ResponseMessage{}, err
	}

	result, _ := ioutil.ReadAll(resp.Body)

	message := ResponseMessage{}
	err = json.Unmarshal(result, &message)
	if err != nil {
		return ResponseMessage{}, err
	}
	return message, nil
}
