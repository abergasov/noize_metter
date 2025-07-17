package noise_metter

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	reSession = regexp.MustCompile(`const _webSession\s*=\s*'([^']+)'`)
)

func (s *Service) bgSetSession() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			if err := s.Auth(); err != nil {
				s.log.Error("failed to set session: %v", err)
			}
		}
	}
}

func (s *Service) Auth() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	form := url.Values{}
	form.Set("p1", s.conf.RemotePass)
	form.Set("cmd", "login")
	form.Set("login_cont_path", "")
	form.Set("disable-pwd-mgr-1", "disable-pwd-mgr-1")
	form.Set("disable-pwd-mgr-2", "disable-pwd-mgr-2")
	form.Set("disable-pwd-mgr-3", "disable-pwd-mgr-3")
	form.Set("p2", s.conf.RemotePass)

	targetURL := fmt.Sprintf("http://%s/do", s.conf.RemoteHost)
	encoded := form.Encode()
	req, err := http.NewRequestWithContext(s.ctx, http.MethodPost, targetURL, strings.NewReader(encoded))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(encoded)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", fmt.Sprintf("http://%s", s.conf.RemoteHost))
	req.Header.Set("Referer", fmt.Sprintf("http://%s/login?", s.conf.RemoteHost))
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: got %d", resp.StatusCode)
	}
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader: %w", err)
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()

	body, _ := io.ReadAll(reader)

	sessionMatch := reSession.FindStringSubmatch(string(body))
	if len(sessionMatch) < 2 {
		return fmt.Errorf("invalid session header")
	}

	s.session.Store(sessionMatch[1])
	return nil
}
