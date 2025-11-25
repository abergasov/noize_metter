package repository

import (
	"context"
	"fmt"
	"net/http"
	"noize_metter/internal/utils/requests"
)

func (r *Repo) FetchLatestVersion(ctx context.Context) (string, error) {
	type v struct {
		Version string `json:"version"`
	}
	hostURL := fmt.Sprintf("%s/api-mapi/v1/private/noiser/version", r.conf.DataHost)
	res, code, err := requests.GetCurl[v](ctx, hostURL, map[string]string{
		"Content-Type": "application/json",
		"auth-mapi":    r.conf.APIKey,
	})
	if err != nil {
		return "", fmt.Errorf("get noiser version, code %d: %w", code, err)
	}
	if code != http.StatusOK {
		return "", fmt.Errorf("get noiser version: %d", code)
	}
	return res.Version, nil
}
