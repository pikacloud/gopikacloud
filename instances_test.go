package gopikacloud

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestInstance_instancePath(t *testing.T) {
	var pathTests = []struct {
		input    interface{}
		expected string
	}{
		{nil, "loadbalancers/instances/"},
		{42, "loadbalancers/instances/42/"},
		{"42", "loadbalancers/instances/42/"},
		{Instance{ID: 64}, "loadbalancers/instances/64/"},
	}

	for _, pt := range pathTests {
		actual := instancePath(pt.input)
		if actual != pt.expected {
			t.Errorf("instancePath(%+v): expected %s, actual %s", pt.input, pt.expected, actual)
		}
	}
}

func TestInstance_Instances(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancers/instances/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":42, "sticky":false, "region":"lon1" }]`)
	})

	instances, err := client.Instances()

	if err != nil {
		t.Errorf("Instances returned error: %v", err)
	}

	want := []Instance{{ID: 42, Sticky: false, RegionName: "lon1"}}
	if !reflect.DeepEqual(instances, want) {
		t.Errorf("Instances returned %+v, want %+v", instances, want)
	}
}

func TestInstance_Instance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancers/instances/42/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":42, "sticky":false, "region":"lon1" }`)
	})

	instance, err := client.Instance(42)

	if err != nil {
		t.Errorf("Instance returned error: %v", err)
	}

	want := Instance{ID: 42, Sticky: false, RegionName: "lon1"}
	if !reflect.DeepEqual(instance, want) {
		t.Errorf("Instance returned %+v, want %+v", instance, want)
	}
}

func TestInstance_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancers/instances/", func(w http.ResponseWriter, r *http.Request) {
		want := make(map[string]interface{})
		want = map[string]interface{}{"region": "lon1", "sticky": false}

		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		testRequestJSON(t, r, want)

		fmt.Fprintf(w, `{"id":42, "sticky":false, "region": "lon1" }`)
	})

	instance := Instance{RegionName: "lon1"}
	res, err := client.CreateInstance(instance)

	if err != nil {
		t.Errorf("Create instance returned error: %v", err)
	}

	want := Instance{ID: 42, Sticky: false, RegionName: "lon1"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Create instance returned %+v, want %+v", res, want)
	}
}

func TestInstance_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/loadbalancers/instances/42/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	instance := Instance{ID: 42}
	err := instance.Delete(client)

	if err != nil {
		t.Errorf("Delete returned error: %v", err)
	}
}
