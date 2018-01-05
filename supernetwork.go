package gopikacloud

import "fmt"

// SuperNetwork defines a pikacloud supernetwork
type SuperNetwork struct {
	Key       string `json:"key"`
	Cidr      string `json:"cidr"`
	DNSDomain string `json:"dns_domain"`
}

// SuperNetwork returns current supernetwork
func (client *Client) SuperNetwork(aid string) (SuperNetwork, error) {
	sn := SuperNetwork{}
	uri := fmt.Sprintf("run/agents/%s/supernetwork/", aid)
	if err := client.Get(uri, &sn); err != nil {
		return SuperNetwork{}, err
	}
	return sn, nil
}
