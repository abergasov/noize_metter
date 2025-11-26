package requests

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"noize_metter/internal/logger"
	"time"
)

var (
	sharedTransport = &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}
	insecureTransport = &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	httpClient = &http.Client{
		Transport: sharedTransport,
		Timeout:   30 * time.Second,
	}
	insecureClient = &http.Client{
		Transport: insecureTransport,
	}
)

func curlWithBody[T any](
	ctx context.Context,
	conf *Config,
	method,
	targetURL string,
	payloadJSON []byte,
	headers map[string]string,
) (res T, statusCode int, err error) {
	req, err := http.NewRequestWithContext(ctx, method, targetURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return res, 0, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return WrappedRun[T](req, conf)
}

func WrappedRun[T any](req *http.Request, config *Config) (res T, statusCode int, err error) {
	if config.basicAuth != nil {
		req.SetBasicAuth(config.basicAuth.username, config.basicAuth.password)
	}
	startTime := time.Now()
	res, statusCode, err = execute[T](req, config)
	if config.logger != nil {
		requestName := config.requestMark
		if requestName == "" {
			requestName = req.RequestURI
		}
		config.logger.With(
			logger.WithExternalRPC(requestName),
			logger.WithRemoteTarget(req.URL.Host),
			logger.WithLatencyFlag(),
			logger.WithHTTPRequest(),
			logger.WithDuration(startTime),
		).Info("rpc call started")
	}
	return res, statusCode, err
}

func execute[T any](req *http.Request, conf *Config) (res T, statusCode int, err error) {
	var client *http.Client
	if conf.skipTLSVerification {
		client = insecureClient
	} else {
		client = httpClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return res, 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return res, 0, fmt.Errorf("failed to create gzip reader, code %d: %w", resp.StatusCode, err)
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()
	if conf.Decoder != nil {
		return res, resp.StatusCode, conf.Decoder(reader)
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		return res, resp.StatusCode, fmt.Errorf("failed to read response body, code %d: %w", resp.StatusCode, err)
	}

	var result T
	if err = json.Unmarshal(b, &result); err != nil {
		return res, resp.StatusCode, fmt.Errorf("failed to unmarshal response, code %d: %w", resp.StatusCode, err)
	}
	return result, resp.StatusCode, nil
}
