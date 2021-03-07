package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns the fastest url", func(t *testing.T) {
		slowServer := makeWebServer(20 * time.Millisecond)
		fastServer := makeWebServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("got %q, want %q", got, want)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond in time", func(t *testing.T) {
		server := makeWebServer(25 * time.Millisecond)

		_, err := ConfigurableRacer(server.URL, server.URL, 15*time.Millisecond)

		if err == nil {
			t.Error("expected to get an error but didn't get one")
		}
	})
}

func makeWebServer(responseDelay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(responseDelay)
				w.WriteHeader(http.StatusOK)
			}))
}
