package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns faster address", func(t *testing.T) {
		slowServer := NewDelayedServer(20 * time.Millisecond)
		fastServer := NewDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		have := fastUrl
		want, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("received an error when not expecting one: %s", err)
		}

		if have != want {
			t.Errorf("have %s want %s", have, want)
		}
	})

	t.Run("times out", func(t *testing.T) {
		serverA := NewDelayedServer(11 * time.Millisecond)
		serverB := NewDelayedServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Fatalf("did not receive expected error: %s", err)
		}
	})
}

func NewDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}),
	)
}
