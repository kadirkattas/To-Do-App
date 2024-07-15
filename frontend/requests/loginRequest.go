package requests

import (
	"bytes"
	"fmt"
	"net/http"
)

func PostRequest(apiURL string, jsonData []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error while creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending HTTP request: %v", err)
	}

	return resp, nil
}
