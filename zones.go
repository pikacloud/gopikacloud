package gopikacloud

import "fmt"

type Zone struct {
	ID         int    `json:"id,omitempty"`
	DomainName string `json:"domain_name"`
	CreatedAt  string `json:"created_at,omitempty"`
	Serial     int    `json:"serial,omitempty"`
}

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

func (client *Client) Zones() ([]Zone, error) {
	zones := []Zone{}
	if err := client.get(zonePath(nil), &zones); err != nil {
		return []Zone{}, err
	}
	return zones, nil
}

func (client *Client) Zone(zone interface{}) ([]Zone, error) {
	zones := []Zone{}
	if err := client.get(zonePath(zone), &zones); err != nil {
		return []Zone{}, err
	}
	return zones, nil
}
