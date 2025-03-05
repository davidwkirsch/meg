package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var transport = &http.Transport{
	TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	DisableKeepAlives: true,
	DialContext: (&net.Dialer{
		Timeout:    error
}

func goRequest(r request) response {
	httpClient.Timeout = r.timeout

	if !r.followLocation {
		httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	var req *http.Request
	var err error
	if r.body != "" {
		req, err = http.NewRequest(r.method, r.URL(), bytes.NewBuffer([]byte(r.body)))
	} else {
		req, err = http.NewRequest(r.method, r.URL(), nil)
	}

	if err != nil {
		return response{request: r, err: err}
	}
	req.Close = true

	if !r.HasHeader("Host") {
		r.headers = append(r.headers, fmt.Sprintf("Host: %s", r.Hostname()))
	}

	if !r.HasHeader("User-Agent") {
		r.headers = append(r.headers, fmt.Sprintf("User-Agent: %s", "Go-http-client/1.1"))
	}

	for _, h := range r.headers {
		parts := strings(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return response{request: r, err: err}
	}
	body, _ := ioutil.ReadAll(resp.Body)

	hs := make([]string, 0)
	for k, vs := range resp.Header {
		for _, v := range vs {
			hs = append(hs, fmt.Sprintf("%s: %s", k, v))
		}
	}

	return response{
		request:    r,
		status:     resp.Status// Example usage
	req := request{
		method:         "GET",
		host:           "https://example.com",
		path:           "/",
		headers:        []string{"Custom-Header: value"},
		followLocation: true,
		timeout:        10 * time.Second,
	}
	resp := goRequest(req)
	if resp.err != nil {
		fmt.Println("Request failed:", resp.err)
	} else {
		fmt.Println("Response status:", resp.status)
		fmt.Println("Response headers:", resp.headers)
		fmt.Println("Response body:", string(resp.body))
	}
}
