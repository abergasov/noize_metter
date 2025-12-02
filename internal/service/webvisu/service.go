package webvisu

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/hex"
	"net/http"
	"net/http/cookiejar"

	"fmt"
	"noize_metter/internal/config"
	"noize_metter/internal/entities"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/utils"

	"sync/atomic"
	"time"
)

const (
	baseURL = "https://192.168.1.12:8081/webvisu/WebVisuV3.bin"
	cookie  = "OriginalDevicePixelRatio=1.25"
	ua      = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"
)

type Service struct {
	ctx  context.Context
	log  logger.AppLogger
	conf *config.AppConfig
	repo *repository.Repo

	session atomic.Value
	cookie  atomic.Value
	items   *utils.RWSlice[entities.NoiseMeasures]
}

func NewService(ctx context.Context, log logger.AppLogger, conf *config.AppConfig, repo *repository.Repo) *Service {
	srv := &Service{
		ctx:     ctx,
		log:     log,
		conf:    conf,
		repo:    repo,
		session: atomic.Value{},
		cookie:  atomic.Value{},
		items:   utils.NewRWSlice[entities.NoiseMeasures](),
	}
	return srv
}

func (s *Service) Run() error {
	ctx, cancel := context.WithTimeout(s.ctx, 10*time.Second)
	defer cancel()

	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("create cookie jar: %w", err)
	}
	requests := []*http.Request{
		s.generateRequest(ctx, "AgAAAFXNEAABAAIAzdlWAhAAAAAAAAAAIoSAAAIAAAAlhIAAAQAAAA", "40", "0200000055c38d100001000200c3acc3bc1d47100000000000000022c284c280000200000025c284c2800001000000"),
		s.generateRequest(ctx, "", "314", "0200000055c38d100001000200c3acc3bc1d47220100000000000022c284c280000200000025c284c2800002000000c28101c28e0210c286c2800061646d696e0011c280c282000bc383775b31c2b8c2bcc2aac2bdc2b1c3b0c3be70c3bc083fc38617c390c28ac2b7c2af1ac3844c32c38ec298c281c386c3a1c392187dc38143c38f6d475823c39f02c39a62c39e36c3b86411c29e5d1bc3847b38c2bd5dc3aac38629355c1ec2a4c286c399c384c2ab23c2bf1536c2a4c3b334c2b2265258720d194ac3a2c2ad723cc2b20f551dc2a0413a4469c2b3c3b81a0626c38b2521c29064c2bf2250c39118c2a06d71c392c3b0c39dc38f7ec3871d491c6fc3b4c3b6c2a4240e4d7010296f4a78c39515c3ab3031c38b56101131493dc286c38c61c3bfc296c3afc283424911c3b2c395c296c393c2b44b78c289c3bc337408c2bd1b233fc29a7435c38364165dc2806a38c38406c39ac38fc2890c28c385c3abc2bdc2800e62c38ac390c29ec2b7c282094416c29237655e26620225242f1168c2b05b45c3a34d06c3a3c2bec2a8c3adc28349c2a501c2afc2a30ac39343c3b13dc3a43c214dc3a32dc2bc000000c3a7c29ac29a69c3a7c388c393"),
		s.generateRequest(ctx, "", "136", "0200000055c38d100004000100c3acc3bc1d47700000000000000001c3acc280004170706c69636174696f6e00000008005c0000000100000002003132372e302e302e3100c38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39dc38cc39d"),
		s.generateRequest(ctx, "AgAAAFXNEAAEAAMAzdlWAggAAAAAAAAAA4SAAKUEDCY", "32", "0200000055c38d100004000300c3acc3bc1d47080000000000000003c284c2800015c3981f51"),
		s.generateRequest(ctx, "AgAAAFXNEAAEAAQAzdlWAigAAAAAAAAAhAGkAAGQgAAEAgAAFQAAAAAAAAClBAwmAoyAAAAAAAD7AsECAACgPw", "64", "0200000055c38d100004000400c3acc3bc1d472800000000000000c28401c2a40001c290c2800004020000150000000000000015c3981f5102c28cc2800000000000c3bb02c381020000c2a03f"),
		s.generateRequest(ctx, "AgAAAFXNEAAEAAQAzdlWAiwAAAAAAAAAhAGoAAGQgAAAABAAAAAAAAAAAAClBAwmApCAAAAADwAHAAAAAAEAAGVuAAA", "68", "0200000055c38d100004000400c3acc3bc1d472c00000000000000c28401c2a80001c290c2800000001000000000000000000015c3981f5102c290c2800000000f000700000000010000656e0000"),
		s.generateRequest(ctx, "", "72", "0200000055c38d100004000400c3acc3bc1d473000000000000000c28401c2ac0001c290c2800000001000000000000000000015c3981f5102c294c280000100000056697375616c697a6174696f6e000000"),
		s.generateRequest(ctx, "AgAAAFXNEAAEAAQAzdlWAhgAAAAAAAAAhAGUAAGQgAABAAAAAAAAAAAAAAClBAwm", "48", "0200000055c38d100004000400c3acc3bc1d471800000000000000c28401c2940001c290c2800001000000000000000000000015c3981f51"),
	}

	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	for _, r := range requests {
		res, errR := client.Do(r)
		if errR != nil {
			return errR
		}
		defer res.Body.Close()
		println(res.StatusCode, res.Body == http.NoBody)
	}
	// now we authorized

	rList := []*http.Request{
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
		s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c28857c2a6c3b11800000000000000c28401c2940001c290c280000100000000000000000000005cc2b97806"),
	}
	println("-----------------------")
	for _, req := range rList {
		//req := s.generateRequest(ctx, "", "48", "0200000055c38d100004000400c3acc3bc1d471800000000000000c28401c2940001c290c2800001000000000000000000000015c3981f51")
		res, errR := client.Do(req)
		if errR != nil {
			return errR
		}
		defer res.Body.Close()

		println(res.StatusCode, res.Body == http.NoBody)
	}

	return nil
}

func (s *Service) generateRequest(ctx context.Context, customHeader, contentLength, hexPayload string) *http.Request {
	var body *bytes.Reader
	if hexPayload != "" {
		// hexPayload must be pure hex, no spaces/newlines
		data, err := hex.DecodeString(hexPayload)
		if err != nil {
			s.log.Fatal("failed decode hex payload", err)
		}
		body = bytes.NewReader(data)
	} else {
		body = bytes.NewReader(nil)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL, body)
	if err != nil {
		s.log.Fatal("failed create request", err)
	}

	if customHeader != "" {
		req.Header.Set("3S-Repl-Content", customHeader)
	}

	req.Header.Set("Content-Length", contentLength)
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Origin", "https://192.168.1.12:8081")
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="99", "Chromium";v="142"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("Cookie", cookie)

	return req
}
