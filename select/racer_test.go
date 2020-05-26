package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("response takes less than 10 second", func(t *testing.T) {
		slowServer := makeServer(12 * time.Millisecond)
		fastServer := makeServer(11 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil || got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("response takes less than 10 second", func(t *testing.T) {
		slowServer := makeServer(1 * time.Second)
		fastServer := makeServer(2 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeServer(durationInMS time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(durationInMS)
		w.WriteHeader(http.StatusOK)
	}))
}
