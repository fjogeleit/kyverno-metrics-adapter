package client

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type Client interface {
	FetchMetrics() (io.Reader, error)
}

type client struct {
	endpoint   string
	httpClient *http.Client
}

func (c *client) FetchMetrics() (io.Reader, error) {
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		log.Printf("[ERROR] : %v\n", err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("User-Agent", "Kyverno Metrics Adapter")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return new(bytes.Buffer), err
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf, err
}

func NewClient(endpoint string) Client {
	return &client{endpoint, &http.Client{}}
}
