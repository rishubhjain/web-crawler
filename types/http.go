package types

import (
	"net/http"
	"time"
)

type HTTPGet interface {
	Get(url string) (*http.Response, error)
}

type RetryingClient struct {
	Client      *http.Client
	MaxAttempts int
}

func (rc *RetryingClient) Get(url string) (*http.Response, error) {
	resp, err := rc.Client.Get(url)
	if err != nil {
		if rc.MaxAttempts <= 0 {
			return resp, err
		}
		rc.MaxAttempts--
		// ToDo : Make this configurable
		<-time.Tick(2 * time.Second)
		rc.Get(url)
	}
	return resp, err
}
