package gopikacloud

import (
	"errors"
	"fmt"
	"strconv"
)

// Zone definition
type Zone struct {
	ID         int    `json:"id,omitempty"`
	DomainName string `json:"domain_name"`
	CreatedAt  string `json:"created_at,omitempty"`
	Serial     int    `json:"serial,omitempty"`
}

func zoneIdentifier(value interface{}) string {
	switch value := value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case string:
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%d", valueInt)
	case Zone:
		return fmt.Sprintf("%d", value.ID)
	}
	return ""
}

func zonePath(zone interface{}) string {
	if zone != nil {
		return fmt.Sprintf("zones/%s/", zoneIdentifier(zone))
	}
	return "zones/"
}

// Zones lists DNS zones you own
func (client *Client) Zones() ([]Zone, error) {
	zones := []Zone{}
	if err := client.get(zonePath(nil), &zones); err != nil {
		return []Zone{}, err
	}
	return zones, nil
}

// Zone retrieve a specific zone
func (client *Client) Zone(zone interface{}) (Zone, error) {
	res := Zone{}
	if err := client.get(zonePath(zone), &res); err != nil {
		return Zone{}, err
	}
	return res, nil
}

// CreateZone create a zone
func (client *Client) CreateZone(zone interface{}) (Zone, error) {
	res := Zone{}
	status, err := client.post(zonePath(nil), zone, &res)
	if err != nil {
		return Zone{}, err
	}
	if status == 201 {
		return res, nil
	}
	return res, errors.New("Failed to create zone")
}

// Delete a Zone
func (zone *Zone) Delete(client *Client) error {

	_, status, err := client.sendRequest("DELETE", zonePath(zone.ID), nil)
	if err != nil {
		return err
	}
	if status == 204 {
		return nil
	}
	return errors.New("Failed to delete zone")
}
