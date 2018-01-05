package gopikacloud

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestZoneRecord_zonerecordPath(t *testing.T) {
	var pathTests = []struct {
		zoneInput       interface{}
		zoneRecordInput interface{}
		expected        string
	}{
		{13, nil, "dns/zones/13/records/"},
		{Zone{ID: 13}, nil, "dns/zones/13/records/"},
		{13, 42, "dns/zones/13/records/42/"},
		{13, ZoneRecord{ID: 42}, "dns/zones/13/records/42/"},
		{Zone{ID: 13}, ZoneRecord{ID: 42}, "dns/zones/13/records/42/"},
	}

	for _, pt := range pathTests {
		actual := zoneRecordPath(pt.zoneInput, pt.zoneRecordInput)
		if actual != pt.expected {
			t.Errorf("zoneRecordPath(%+v, %+v): expected %s, actual %s", pt.zoneInput, pt.zoneRecordInput, pt.expected, actual)
		}
	}
}

func TestZoneRecord_ZoneRecords(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/zones/13/records/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": 42, "zone": 13, "rtype": "A", "ipv4": "127.0.0.1", "ttl": 1800}]`)
	})

	zoneRecords, err := client.ZoneRecords(13)

	if err != nil {
		t.Errorf("Zone records returned error: %v", err)
	}

	want := []ZoneRecord{{ID: 42, Rtype: "A", Ipv4: "127.0.0.1", TTL: 1800, ZoneID: 13}}
	if !reflect.DeepEqual(zoneRecords, want) {
		t.Errorf("Zone records returned %+v, want %+v", zoneRecords, want)
	}
}

func TestZone_ZoneRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/zones/13/records/42/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":42, "zone": 13, "rtype": "A", "ipv4": "127.0.0.1", "ttl": 1800}`)
	})

	zoneRecord, err := client.ZoneRecord(13, 42)

	if err != nil {
		t.Errorf("Zone record returned error: %v", err)
	}

	want := ZoneRecord{ID: 42, Rtype: "A", Ipv4: "127.0.0.1", TTL: 1800, ZoneID: 13}
	if !reflect.DeepEqual(zoneRecord, want) {
		t.Errorf("Zone record returned %+v, want %+v", zoneRecord, want)
	}
}

func TestZoneRecord_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/zones/13/records/", func(w http.ResponseWriter, r *http.Request) {
		want := map[string]interface{}{"rtype": "A", "ipv4": "127.0.0.1"}

		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		testRequestJSON(t, r, want)

		fmt.Fprintf(w, `{"id": 42, "zone": 13, "rtype": "A", "ipv4": "127.0.0.1"}`)
	})

	zoneRecord := ZoneRecord{Rtype: "A", Ipv4: "127.0.0.1"}
	res, err := client.CreateZoneRecord(13, zoneRecord)

	if err != nil {
		t.Errorf("Create zone returned error: %v", err)
	}

	want := ZoneRecord{ID: 42, Rtype: "A", Ipv4: "127.0.0.1", ZoneID: 13}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Create zone returned %+v, want %+v", res, want)
	}
}

func TestZoneRecord_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/dns/zones/13/records/42/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	zoneRecord := ZoneRecord{ID: 42, ZoneID: 13}
	err := zoneRecord.Delete(client)

	if err != nil {
		t.Errorf("Delete returned error: %v", err)
	}
}
