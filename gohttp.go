package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var transport = &http.Transport{
	TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	DisableKeepAlives: true,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
}

var httpClient = &http.Client{
	Transport: transport,
}

type request struct {
	method         string
	path           string
	host           string
	headers        []string
	body           string
	followLocation bool
	timeout        time.Duration
}

type response struct {
	request    request
	status     string
	statusCode int
	headers    []string
	body       []byte
	err        error
}

func (r request) Hostname() string {
	u, err := url.Parse(r.host)
	if err != nil {
		return "unknown"
	}
	return u.Hostname()
}

func (r request) URL() string {
	return r.host + r.path
}

func (r request) HasHeader(h string) bool {
	norm := func(s string) string {
		return strings.ToLower(strings.TrimSpace(s))
	}
	for _, candidate := range r.headers {
		p := strings.SplitN(candidate, ":", 2)
		if norm(p[0]) == norm(h) {
			return true
		}
	}
	return false
}

func goRequest(r request) response {
	httpClient.Timeout = r.timeout

	if !r.followLocation {
		httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
	http.Request
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
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			continue
		}
		req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	resp, err := httpClient.Do(req)
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
		status:     resp.Status,
		statusCode: resp.StatusCode,
		headers:    hs,
		body:       body,
	}
}

func main() {
	// Example usage
	req := request{
		method:         "GET",
		host:           "https://example.com",
		path:           "/",
		headers:        []string{"Custom-Header: value"},
		followLocation: true,
		timeout:        10 * time.Second,
	}
	resp := goRequest(req)
	if.err)
	} else {
		fmt.Println("Response status:", resp.status)
		fmt.Println("Response headers:", resp.headers)
		fmt.Println("Response body:", string(resp.body))
	}
}
