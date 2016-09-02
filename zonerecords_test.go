package gopikacloud

import "testing"

func TestZoneRecord_zonerecordPath(t *testing.T) {
	var pathTests = []struct {
		zoneInput       interface{}
		zoneRecordInput *ZoneRecord
		expected        string
	}{
		{13, nil, "zones/13/records/"},
		{Zone{ID: 13}, nil, "zones/13/records/"},
		{13, &ZoneRecord{ID: 42}, "zones/13/records/42/"},
		{Zone{ID: 13}, &ZoneRecord{ID: 42}, "zones/13/records/42/"},
	}

	for _, pt := range pathTests {
		actual := zoneRecordPath(pt.zoneInput, pt.zoneRecordInput)
		if actual != pt.expected {
			t.Errorf("zoneRecordPath(%+v, %+v): expected %s, actual %s", pt.zoneInput, pt.zoneRecordInput, pt.expected, actual)
		}
	}
}

// func TestZone_Zones(t *testing.T) {
// 	setup()
// 	defer teardown()
//
// 	mux.HandleFunc("/v1/zones/", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `[{"id": 1, "domain_name": "foo.com", "serial": 10, "created_at": "2016-08-23T21:59:14.000251Z"}]`)
// 	})
//
// 	zones, err := client.Zones()
//
// 	if err != nil {
// 		t.Errorf("Zones returned error: %v", err)
// 	}
//
// 	want := []Zone{{ID: 1, DomainName: "foo.com", Serial: 10, CreatedAt: "2016-08-23T21:59:14.000251Z"}}
// 	if !reflect.DeepEqual(zones, want) {
// 		t.Errorf("Zones returned %+v, want %+v", zones, want)
// 	}
// }
//
// func TestZoneRecord_ZoneRecord(t *testing.T) {
// 	setup()
// 	defer teardown()
//
// 	mux.HandleFunc("/v1/zones/13/records/42/", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `[{"id":13}]`)
// 	})
//
// 	zone, err := client.Zone(42)
//
// 	if err != nil {
// 		t.Errorf("Zone returned error: %v", err)
// 	}
//
// 	want := []Zone{{ID: 42, DomainName: "example.com"}}
// 	if !reflect.DeepEqual(zone, want) {
// 		t.Errorf("Zone returned %+v, want %+v", zone, want)
// 	}
// }
//
// func TestZoneRecord_Delete(t *testing.T) {
// 	setup()
// 	defer teardown()
//
// 	mux.HandleFunc("/v1/zones/13/records/42/", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "DELETE")
// 		w.WriteHeader(http.StatusNoContent)
// 	})
//
// 	zoneRecord := ZoneRecord{ID: 42, ZoneId: 13}
// 	err := zoneRecord.Delete(client)
//
// 	if err != nil {
// 		t.Errorf("Delete returned error: %v", err)
// 	}
// }
