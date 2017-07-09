package gopikacloud

import (
	"errors"
	"fmt"
	"strconv"
)

// Instance definition
type Instance struct {
	ID          int    `json:"id,omitempty"`
	RegionName  string `json:"region"`
	Sticky      bool   `json:"sticky"`
	Certificate int    `json:"certificate,omitempty"`
}

func instanceIdentifier(value interface{}) string {
	switch value := value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case string:
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%d", valueInt)
	case Instance:
		return fmt.Sprintf("%d", value.ID)
	}
	return ""
}

func instancePath(instance interface{}) string {
	if instance != nil {
		return fmt.Sprintf("loadbalancers/instances/%s/", instanceIdentifier(instance))
	}
	return "loadbalancers/instances/"
}

// Instances lists load balancers you own
func (client *Client) Instances() ([]Instance, error) {
	instances := []Instance{}
	if err := client.Get(instancePath(nil), &instances); err != nil {
		return []Instance{}, err
	}
	return instances, nil
}

// Instance retrieve a specific instance
func (client *Client) Instance(instance interface{}) (Instance, error) {
	res := Instance{}
	if err := client.Get(instancePath(instance), &res); err != nil {
		return Instance{}, err
	}
	return res, nil
}

// CreateInstance create a instance
func (client *Client) CreateInstance(instance interface{}) (Instance, error) {
	res := Instance{}
	status, err := client.Post(instancePath(nil), instance, &res)
	if err != nil {
		return Instance{}, err
	}
	if status == 201 {
		return res, nil
	}
	return res, errors.New("Failed to create instance")
}

// Delete a Instance
func (instance *Instance) Delete(client *Client) error {

	_, status, err := client.sendRequest("DELETE", instancePath(instance.ID), nil)
	if err != nil {
		return err
	}
	if status == 204 {
		return nil
	}
	return errors.New("Failed to delete instance")
}
