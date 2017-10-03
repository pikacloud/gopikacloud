package gopikacloud

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUser_Me(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/me/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":42, "email": "foo@bar.com"}`)
	})

	user, err := client.Me()

	if err != nil {
		t.Errorf("User returned error: %v", err)
	}

	want := User{ID: 42, Email: "foo@bar.com"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("User returned %+v, want %+v", user, want)
	}
}
