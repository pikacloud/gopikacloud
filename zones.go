package gopikacloud

import (
	"errors"
	"fmt"
)

// Zone definition
type Zone struct {
	ID         int    `json:"id,omitempty"`
	DomainName string `json:"domain_name"`
	CreatedAt  string `json:"created_at,omitempty"`
	Serial     int    `json:"serial,omitempty"`
}

// ZoneWrapper wraps Zone
type ZoneWrapper struct {
	Zone []Zone
}

func zoneIdentifier(value interface{}) string {
	switch value := value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
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
func (client *Client) Zone(zone interface{}) ([]Zone, error) {
	zones := []Zone{}
	if err := client.get(zonePath(zone), &zones); err != nil {
		return []Zone{}, err
	}
	return zones, nil
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
