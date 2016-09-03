package gopikacloud

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	defaultBaseURL = "https://pikacloud.com/api/"
	apiVersion     = "v1"
)

// Client manages communication with Pikacloud API.
type Client struct {
	// API Token for authenticating
	APIToken string
	// HTTP client used to communicate with the Pikacloud API.
	HTTPClient *http.Client
	// Base URL for API requests.
	BaseURL string
}

// NewClient users
func NewClient(apiToken string) *Client {
	return &Client{APIToken: apiToken, HTTPClient: &http.Client{}, BaseURL: defaultBaseURL}
}

func (client *Client) makeRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := client.BaseURL + fmt.Sprintf("%s/%s", apiVersion, path)
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", fmt.Sprintf("Token: %s", client.APIToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *Client) get(path string, val interface{}) error {
	body, _, err := client.sendRequest("GET", path, nil)
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(body), &val); err != nil {
		return err
	}

	return nil
}

func (client *Client) delete(path string, val interface{}) error {
	_, _, err := client.sendRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) sendRequest(method, path string, body io.Reader) (string, int, error) {
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
