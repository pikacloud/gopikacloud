package gopikacloud

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSuperNetwork_SuperNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/run/agents/foobar/supernetwork/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"key": "foobar", "cidr": "10.42.0.0/24", "dns_domain": "pikacloud.local"}`)
	})

	sn, err := client.SuperNetwork("foobar")

	if err != nil {
		t.Errorf("SuperNetwork returned error: %v", err)
	}

	want := SuperNetwork{Key: "foobar", Cidr: "10.42.0.0/24", DNSDomain: "pikacloud.local"}
	if !reflect.DeepEqual(sn, want) {
		t.Errorf("SuperNetwork returned %+v, want %+v", sn, want)
	}
}
