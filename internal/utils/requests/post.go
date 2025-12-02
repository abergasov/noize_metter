package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PostCurl is a generic function to send POST request with headers and return response
// additional tracing system can be added on demand
// expect json response
// T is a type of response, automatically unmarshalled from json
func PostCurl[T any](
	ctx context.Context,
	targetURL string,
	payload any,
	headers map[string]string,
	opts ...Option,
) (res T, statusCode int, err error) {
	config := NewDefaultConfig()
	for _, opt := range opts {
		opt(config)
	}

	c, cancel := context.WithTimeout(ctx, config.requestTimeout)
	defer cancel()

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return res, 0, fmt.Errorf("unable to marshal payload: %w", err)
	}
	return curlWithBody[T](c, config, http.MethodPost, targetURL, payloadJSON, headers)
}

func PostCurlOctetStream(ctx context.Context, targetURL string, headers map[string]string, opts ...Option) (res string, statusCode int, err error) {
	config := NewDefaultConfig()
	opts = append(opts, WithDecoder(func(reader io.Reader) error {
		b, eRead := io.ReadAll(reader)
		if eRead != nil {
			return fmt.Errorf("failed to read response body, err: %w", eRead)
		}
		res = string(b)
		return nil
	}))
	for _, opt := range opts {
		opt(config)
	}

	c, cancel := context.WithTimeout(ctx, config.requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(c, http.MethodPost, targetURL, http.NoBody)
	if err != nil {
		return res, 0, fmt.Errorf("unable to create request: %w", err)
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	_, code, err := WrappedRun[string](req, config)
	return res, code, err
}
