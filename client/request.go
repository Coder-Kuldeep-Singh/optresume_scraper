package client

import (
	"crypto/tls"
	"net/http"
	"time"
)

func Request(URL string) (*http.Response, error) {
	// set client to send request
	tr := &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     30 * time.Second,
		DisableCompression:  true,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	// setup client
	client := &http.Client{Transport: tr}

	// Setup to send request DOMAIN/URL
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	// set header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	// send Direct Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
