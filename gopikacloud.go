package gopikacloud

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	defaultBaseURL = "http://localhost:8000/api/"
	apiVersion     = "v1"
)

type gopikacloudClient struct {
	APIToken   string
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(apiToken string) *gopikacloudClient {
	return &gopikacloudClient{APIToken: apiToken, HTTPClient: &http.Client{}, BaseURL: defaultBaseURL}
}

func (client *gopikacloudClient) makeRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := client.BaseURL + fmt.Sprintf("%s/%s", apiVersion, path)
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Authorization: Token", fmt.Sprintf("%s", client.APIToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *gopikacloudClient) get(path string, val interface{}) error {
	body, _, err := client.sendRequest("GET", path, nil)
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(body), &val); err != nil {
		return err
	}

	return nil
}

func (client *gopikacloudClient) sendRequest(method, path string, body io.Reader) (string, int, error) {
	req, err := client.makeRequest(method, path, body)
	if err != nil {
		return "", 0, err
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	return string(responseBytes), resp.StatusCode, nil
}
