package gopikacloud

import (
	"errors"
	"fmt"
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
}

// ZoneRecordsWrapper wraps ZoneRecord
type ZoneRecordsWrapper struct {
	ZoneRecord []ZoneRecord
}

func zoneRecordIdentifier(value interface{}) string {
	switch value := value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case ZoneRecord:
		return fmt.Sprintf("%d", value.ID)
	}
	return ""
}

func zoneRecordPath(zone interface{}, zonerecord *ZoneRecord) string {
	str := fmt.Sprintf("zones/%s/records/", zoneIdentifier(zone))
	if zonerecord != nil {
		str += fmt.Sprintf("%s/", zoneRecordIdentifier(*zonerecord))
	}
	return str
}

// ZoneRecords lists DNS zone Records of a specific zone
func (client *Client) ZoneRecords(zone interface{}) ([]ZoneRecord, error) {
	zoneRecords := []ZoneRecord{}
	if err := client.get(zoneRecordPath(zone, nil), &zoneRecords); err != nil {
		return []ZoneRecord{}, err
	}
	return zoneRecords, nil
}

// ZoneRecord retrieve a specific zone record
func (client *Client) ZoneRecord(zone interface{}) ([]Zone, error) {
	zones := []Zone{}
	if err := client.get(zonePath(zone), &zones); err != nil {
		return []Zone{}, err
	}
	return zones, nil
}

// Delete a zone record
func (zoneRecord *ZoneRecord) Delete(client *Client) error {
	_, status, err := client.sendRequest("DELETE", zoneRecordPath(zoneRecord.ZoneID, zoneRecord), nil)
	if err != nil {
		return err
	}
	if status == 204 {
		return nil
	}
	return errors.New("Failed to delete zone record")
}
