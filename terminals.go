package gopikacloud

import (
	"errors"
	"fmt"
)

// Terminal definition
type Terminal struct {
	Aid      string `json:"aid,omitempty"`
	Tid      string `json:"tid,omitempty"`
	Cid      string `json:"cid"`
	Token    string `json:"token,omitempty"`
	LastPing int    `json:"last_ping,omitempty"`
	Ready    bool   `json:"ready,omitempty"`
}

func terminalPath(aid string, tid string) string {
	if tid != "" {
		return fmt.Sprintf("run/agents/%s/docker/terminals/%s/", aid, tid)
	}
	return fmt.Sprintf("run/agents/%s/docker/terminals/", aid)
}

// Terminals list terminals for an agent and container id
func (client *Client) Terminals(aid string) ([]Terminal, error) {
	terminals := []Terminal{}
	if err := client.Get(terminalPath(aid, ""), &terminals); err != nil {
		return []Terminal{}, err
	}
	return terminals, nil
}

// Terminal retrieves a specific terminal
func (client *Client) Terminal(aid string, tid string) (Terminal, error) {
	res := Terminal{}
	if err := client.Get(terminalPath(aid, tid), &res); err != nil {
		return Terminal{}, err
	}
	return res, nil
}

// CreateTerminal creates a terminal
func (client *Client) CreateTerminal(aid string, terminal interface{}) (Terminal, error) {
	res := Terminal{}
	status, err := client.Post(terminalPath(aid, ""), terminal, &res)
	if err != nil {
		return Terminal{}, err
	}
	if status == 201 {
		return res, nil
	}
	return res, errors.New("Failed to create terminal")
}

// Delete a terminal
func (terminal *Terminal) Delete(client *Client) error {
	_, status, err := client.sendRequest("DELETE", terminalPath(terminal.Aid, terminal.Tid), nil)
	if err != nil {
		return err
	}
	if status == 204 {
		return nil
	}
	return errors.New("Failed to delete terminal")
}
