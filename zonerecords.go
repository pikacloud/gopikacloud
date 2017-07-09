package gopikacloud

import (
	"errors"
	"fmt"
	"strconv"
)

// ZoneRecord definition
type ZoneRecord struct {
	ID       int    `json:"id,omitempty"`
	ZoneID   int    `json:"zone,omitempty"`
	Rtype    string `json:"rtype"`
	Ipv4     string `json:"ipv4,omitempty"`
	Ipv6     string `json:"ipv6,omitempty"`
	Name     string `json:"name,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Content  string `json:"content,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

func zoneRecordIdentifier(value interface{}) string {
	switch value := value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case ZoneRecord:
		return fmt.Sprintf("%d", value.ID)
	case string:
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%d", valueInt)
	}

	return ""
}

func zoneRecordPath(zone interface{}, zonerecord interface{}) string {
	str := fmt.Sprintf("dns/zones/%s/records/", zoneIdentifier(zone))
	if zonerecord != nil {
		str += fmt.Sprintf("%s/", zoneRecordIdentifier(zonerecord))
	}
	return str
}

// ZoneRecords lists DNS zone Records of a specific zone
func (client *Client) ZoneRecords(zone interface{}) ([]ZoneRecord, error) {
	zoneRecords := []ZoneRecord{}
	if err := client.Get(zoneRecordPath(zone, nil), &zoneRecords); err != nil {
		return []ZoneRecord{}, err
	}
	return zoneRecords, nil
}

// ZoneRecord fetch a single DNS zone record
func (client *Client) ZoneRecord(zone interface{}, zonerecord interface{}) (ZoneRecord, error) {
	zoneRecord := ZoneRecord{}
	if err := client.Get(zoneRecordPath(zone, zonerecord), &zoneRecord); err != nil {
		return ZoneRecord{}, err
	}
	return zoneRecord, nil
}

// CreateZoneRecord create a DNS zone record
func (client *Client) CreateZoneRecord(zone interface{}, zonerecord interface{}) (ZoneRecord, error) {
	res := ZoneRecord{}
	status, err := client.Post(zoneRecordPath(zone, nil), zonerecord, &res)
	if err != nil {
		return ZoneRecord{}, err
	}
	if status == 201 {
		return res, nil
	}
	return res, errors.New("Failed to create zone record")
}

// Delete a zone record
func (zoneRecord *ZoneRecord) Delete(client *Client) error {
	_, status, err := client.sendRequest("DELETE", zoneRecordPath(zoneRecord.ZoneID, zoneRecord.ID), nil)
	if err != nil {
		return err
	}
	if status == 204 {
		return nil
	}
	return errors.New("Failed to delete zone record")
}
