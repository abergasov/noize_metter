package ces

import (
	"context"
	"fmt"
	"net/http"
	"noize_metter/internal/utils/requests"
	"time"
)

func (s *Service) FillToken() error {
	ctx, cancel := context.WithTimeout(s.ctx, 1*time.Minute)
	defer cancel()
	urlList := append(URLListIMDC1, URLListIMDC2...)
	for _, host := range urlList {
		token, err := s.CreateToken(ctx, host)
		if err != nil {
			return fmt.Errorf("failed to get token for host %s: %w", host, err)
		}
		s.tokenMU.Lock()
		s.token[host] = token
		s.tokenMU.Unlock()
	}
	return nil
}

func (s *Service) CreateToken(ctx context.Context, baseURL string) (string, error) {
	type login struct {
		Token string `json:"token"`
	}
	res, code, err := requests.PostCurl[login](
		ctx,
		fmt.Sprintf("%s/api/user/login", baseURL), map[string]string{
			"email":    s.conf.CESUser,
			"password": s.conf.CESPass,
		},
		nil,
		requests.WithSkipTLSVerification(true),
	)
	if err != nil {
		return "", fmt.Errorf("failed autorize user in system, url: %s: %w", baseURL, err)
	}
	if code == http.StatusBadRequest {
		// try register user
		_, code, err = requests.PostCurl[login](ctx, fmt.Sprintf("%s/api/user/register", baseURL), map[string]string{
			"email":    s.conf.CESUser,
			"password": s.conf.CESPass,
			"name":     "mapicron-sync",
		}, nil, requests.WithSkipTLSVerification(true))
		if err != nil {
			return "", fmt.Errorf("failed to register user in system: %w", err)
		}
		if code != http.StatusOK {
			return "", fmt.Errorf("failed to register user in system, wrong code: %d", code)
		}
		res, code, err = requests.PostCurl[login](ctx, fmt.Sprintf("%s/api/user/login", baseURL), map[string]string{
			"email":    s.conf.CESUser,
			"password": s.conf.CESPass,
		}, nil, requests.WithSkipTLSVerification(true))
		if err != nil {
			return "", fmt.Errorf("failed to autorize user after registration: %w", err)
		}
		if code != http.StatusOK {
			return "", fmt.Errorf("failed to autorize user after registration, wrong code: %d", code)
		}
		return res.Token, nil
	}
	if code != http.StatusOK {
		return "", fmt.Errorf("failed autorize user in system, wrong code: %d", code)
	}
	return res.Token, nil
}
