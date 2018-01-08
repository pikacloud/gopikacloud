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

func TestSuperNetwork_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/run/agents/foobar/supernetwork/", func(w http.ResponseWriter, r *http.Request) {
		want := map[string]interface{}{"cidr": "10.42.0.0/16", "dns_domain": "foo.bar."}

		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		testRequestJSON(t, r, want)

		fmt.Fprintf(w, `{"cidr": "10.42.0.0/16", "dns_domain": "foo.bar.", "key": "a"}`)
	})

	superNetwork := SuperNetwork{Cidr: "10.42.0.0/16", DNSDomain: "foo.bar."}
	res, err := client.CreateSuperNetwork(superNetwork, "foobar")

	if err != nil {
		t.Errorf("Create superNetwork returned error: %v", err)
	}

	want := SuperNetwork{Cidr: "10.42.0.0/16", DNSDomain: "foo.bar.", Key: "a"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Create superNetwork returned %+v, want %+v", res, want)
	}
}
