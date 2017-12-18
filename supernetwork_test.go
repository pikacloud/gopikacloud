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

	mux.HandleFunc("/v1/run/supernetwork/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"user":42, "key": "foobar"}`)
	})

	sn, err := client.SuperNetwork("tata")

	if err != nil {
		t.Errorf("SuperNetwork returned error: %v", err)
	}

	want := SuperNetwork{User: 42, Key: "foobar"}
	if !reflect.DeepEqual(sn, want) {
		t.Errorf("SuperNetwork returned %+v, want %+v", sn, want)
	}
}
