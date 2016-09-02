package gopikacloud

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestZone_zonePath(t *testing.T) {
	var pathTests = []struct {
		input    interface{}
		expected string
	}{
		{nil, "zones/"},
		{42, "zones/42/"},
		{Zone{ID: 64}, "zones/64/"},
	}

	for _, pt := range pathTests {
		actual := zonePath(pt.input)
		if actual != pt.expected {
			t.Errorf("zonePath(%+v): expected %s, actual %s", pt.input, pt.expected, actual)
		}
	}
}

func TestZone_Zones(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/zones/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": 1, "domain_name": "foo.com", "serial": 10, "created_at": "2016-08-23T21:59:14.000251Z"}]`)
	})

	zones, err := client.Zones()

	if err != nil {
		t.Errorf("Zones returned error: %v", err)
	}

	want := []Zone{{ID: 1, DomainName: "foo.com", Serial: 10, CreatedAt: "2016-08-23T21:59:14.000251Z"}}
	if !reflect.DeepEqual(zones, want) {
		t.Errorf("Zones returned %+v, want %+v", zones, want)
	}
}

func TestZone_Zone(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/zones/42/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":42, "domain_name":"example.com"}]`)
	})

	zone, err := client.Zone(42)

	if err != nil {
		t.Errorf("Zone returned error: %v", err)
	}

	want := []Zone{{ID: 42, DomainName: "example.com"}}
	if !reflect.DeepEqual(zone, want) {
		t.Errorf("Zone returned %+v, want %+v", zone, want)
	}
}

func TestZone_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/zones/42/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	zone := Zone{ID: 42}
	err := zone.Delete(client)

	if err != nil {
		t.Errorf("Delete returned error: %v", err)
	}
}
