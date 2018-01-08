package gopikacloud

import (
	"errors"
	"fmt"
)

// SuperNetwork defines a pikacloud supernetwork
type SuperNetwork struct {
	Key       string `json:"key,omitempty"`
	Cidr      string `json:"cidr"`
	DNSDomain string `json:"dns_domain"`
}

// SuperNetwork return a supernetwork
func (client *Client) SuperNetwork(aid string) (SuperNetwork, error) {
	sn := SuperNetwork{}
	uri := fmt.Sprintf("run/agents/%s/supernetwork/", aid)
	if err := client.Get(uri, &sn); err != nil {
		return SuperNetwork{}, err
	}
	return sn, nil
}

// CreateSuperNetwork create a supernetwork
func (client *Client) CreateSuperNetwork(supernetwork interface{}, aid string) (SuperNetwork, error) {
	sn := SuperNetwork{}
	uri := fmt.Sprintf("run/agents/%s/supernetwork/", aid)
	status, err := client.Post(uri, supernetwork, &sn)
	if err != nil {
		return SuperNetwork{}, err
	}
	if status == 201 || status == 200 {
		return sn, nil
	}
	return sn, errors.New("Failed to create superNetwork")
}
