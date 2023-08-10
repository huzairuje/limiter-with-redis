package health

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestAuth(t *testing.T) {
	t.Run("TestHealthSuccess", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			Get("http://localhost:1234/api/v1/health/check")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 200 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
		}
	})

	t.Run("TestHealthPingSuccess", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			Get("http://localhost:1234/api/v1/health/ping")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if resp.StatusCode() != 200 {
			t.Errorf("\"Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
		}
	})

}
