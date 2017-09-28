package gopikacloud

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTerminal_terminalPath(t *testing.T) {
	var pathTests = []struct {
		aid      string
		tid      string
		expected string
	}{
		{"foo", "bar", "run/agents/foo/docker/terminals/bar/"},
		{"foo", "", "run/agents/foo/docker/terminals/"},
	}

	for _, pt := range pathTests {
		actual := terminalPath(pt.aid, pt.tid)
		if actual != pt.expected {
			t.Errorf("terminalPath(aid:%s tid:%s): expected %s, actual %s", pt.aid, pt.tid, pt.expected, actual)
		}
	}
}

func TestTerminal_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/run/agents/foo/docker/terminals/bar/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	terminal := Terminal{Aid: "foo", Tid: "bar"}
	err := terminal.Delete(client)

	if err != nil {
		t.Errorf("Delete returned error: %v", err)
	}
}

func TestTerminal_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/run/agents/foo/docker/terminals/", func(w http.ResponseWriter, r *http.Request) {
		want := make(map[string]interface{})
		want = map[string]interface{}{"cid": "xyz"}

		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		testRequestJSON(t, r, want)

		fmt.Fprintf(w, `{"aid": "foo", "tid": "bar", "cid":"xyz"}`)
	})

	terminal := Terminal{Cid: "xyz"}
	res, err := client.CreateTerminal("foo", terminal)

	if err != nil {
		t.Errorf("Create terminal returned error: %v", err)
	}

	want := Terminal{Aid: "foo", Tid: "bar", Cid: "xyz"}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Create terminal returned %+v, want %+v", res, want)
	}
}

func TestTerminal_Terminals(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/run/agents/foo/docker/terminals/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"aid": "foo", "tid": "bar", "cid": "xyz"},{"aid": "foo", "tid": "bar1", "cid": "xyz1"}]`)
	})

	terminals, err := client.Terminals("foo")

	if err != nil {
		t.Errorf("Terminals returned error: %v", err)
	}

	want := []Terminal{{Aid: "foo", Tid: "bar", Cid: "xyz"}, {Aid: "foo", Tid: "bar1", Cid: "xyz1"}}
	if !reflect.DeepEqual(terminals, want) {
		t.Errorf("Terminals returned %+v, want %+v", terminals, want)
	}
}

func TestTerminal_Terminal(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/run/agents/foo/docker/terminals/bar/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"aid": "foo", "tid": "bar", "cid": "xyz"}`)
	})

	terminal, err := client.Terminal("foo", "bar")

	if err != nil {
		t.Errorf("Terminal returned error: %v", err)
	}

	want := Terminal{Aid: "foo", Tid: "bar", Cid: "xyz"}
	if !reflect.DeepEqual(terminal, want) {
		t.Errorf("Terminal returned %+v, want %+v", terminal, want)
	}
}
