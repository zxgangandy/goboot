package utils

import (
	"context"
	"net/http"
	"time"
)

const (
	maxIdleConns        = 100
	maxConnsPerHost     = 100
	maxIdleConnsPerHost = 100
	timeout             = 10
)

func GetDefaultHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = maxIdleConns
	t.MaxConnsPerHost = maxConnsPerHost
	t.MaxIdleConnsPerHost = maxIdleConnsPerHost

	httpClient := &http.Client{
		Timeout:   timeout * time.Second,
		Transport: t,
	}

	return httpClient
}

func GetHttpClient(maxIdleConns, maxConnsPerHost, maxIdleConnsPerHost, timeout int) *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = maxIdleConns
	t.MaxConnsPerHost = maxConnsPerHost
	t.MaxIdleConnsPerHost = maxIdleConnsPerHost

	httpClient := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: t,
	}

	return httpClient
}

func Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	var traceId string
	if v := ctx.Value(TraceKey); v != nil {
		traceId = v.(string)
	} else {
		traceId = RandomString(TraceLen)
	}
	req.Header.Add(TraceKey, traceId)
	req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DoDefault(ctx context.Context, req *http.Request) (*http.Response, error) {
	var traceId string
	if v := ctx.Value(TraceKey); v != nil {
		traceId = v.(string)
	} else {
		traceId = RandomString(TraceLen)
	}
	req.Header.Add(TraceKey, traceId)
	req.WithContext(ctx)

	res, err := GetDefaultHttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
