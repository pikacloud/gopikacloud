package gopikacloud

import "fmt"

// SuperNetwok defines a pikacloud supernetwork
type SuperNetwork struct {
	User int64  `json:"user"`
	Key  string `json:"key"`
}

// SuperNet returns current supernetwork
func (client *Client) SuperNetwork(aid string) (SuperNetwork, error) {
	aidURI := fmt.Sprintf("run/supernetwork/?aid=%s", aid)
	sn := SuperNetwork{}
	if err := client.Get(aidURI, &sn); err != nil {
		return SuperNetwork{}, err
	}
	return sn, nil
}
