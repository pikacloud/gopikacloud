package gopikacloud

// ContainerConfig definition
type ContainerConfig struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	AlwaysPull bool   `json:"always_pull"`
}
