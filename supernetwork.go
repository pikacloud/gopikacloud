package gopikacloud

// SuperNetwok defines a pikacloud supernetwork
type SuperNetwork struct {
	User int64  `json:"user"`
	Key  string `json:"key"`
}

// SuperNetwork returns current supernetwork
func (client *Client) SuperNetwork() (SuperNetwork, error) {
	sn := SuperNetwork{}
	if err := client.Get("run/supernetwork", &sn); err != nil {
		return SuperNetwork{}, err
	}
	return sn, nil
}
