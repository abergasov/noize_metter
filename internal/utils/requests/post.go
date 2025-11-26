package requests

import (
	"context"
	"encoding/json"
	"fmt"
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
