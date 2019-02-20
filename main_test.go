package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJokeHandler(t *testing.T) {

	t.Run("Joke Test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "localhost:8080/", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		GetJokeHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", res.Status)
		}
	})
}

func TestRouting(t *testing.T) {

	r := http.NewServeMux()
	r.HandleFunc("/", GetJokeHandler)
	srv := httptest.NewServer(r)
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d := string(b[:])
	if len(d) < 10 {
		t.Fatalf("expected a string got %v", d)
	}
}
