package noise_metter

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	cloudflarebp "github.com/DaRealFreak/cloudflare-bp-go"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (s *Service) Auth() error {
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
	req.Header.Set("Content-Length", "152")
	req.Header.Set("Cookie", "A3A-01562-F0=4g09w9kz3wz87li2kere0t8am89ne33u")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", fmt.Sprintf("http://%s", s.conf.RemoteHost))
	req.Header.Set("Referer", fmt.Sprintf("http://%s/", s.conf.RemoteHost))
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("Redirect to:", req.URL)
			// Optionally add headers on redirected request:
			req.Header.Set("User-Agent", "Mozilla/5.0 ...")
			return nil // return http.ErrUseLastResponse to stop redirect
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client.Transport = cloudflarebp.AddCloudFlareByPass(client.Transport)
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
	log.Println("Status:", resp.Status)
	log.Println("Body:", string(body))
	return nil
}
