package transport

import (
	"net/http"
	"time"
)

func HttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
	return client
}
