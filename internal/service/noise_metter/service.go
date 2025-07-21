package noise_metter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"noize_metter/internal/config"
	"noize_metter/internal/entities"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/utils"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
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
		log:     log.With(logger.WithService("noise_metter")),
		conf:    conf,
		repo:    repo,
		session: atomic.Value{},
		cookie:  atomic.Value{},
		items:   utils.NewRWSlice[entities.NoiseMeasures](),
	}
	go srv.bgDumpData()
	go srv.bgPruneOldFiles()
	go srv.bgUploadData()
	return srv
}

func (s *Service) Run() {
	s.log.Info("starting Noise Metter service...")
	s.session.Store("")
	if err := s.Auth(); err != nil {
		s.log.Fatal("auth err: ", err)
	}
	go s.bgSetSession()
	// Placeholder for actual service logic
	select {
	case <-s.ctx.Done():
		s.log.Info("Noise Metter service stopped.")
		return
	default:
		if err := s.connectForSession(); err != nil {
			s.log.Error("failed to connect for session", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func (s *Service) Stop() {
	s.log.Info("stopping service")
	s.dumpData()
}

func (s *Service) connectForSession() error {
	sessionID := s.session.Load().(string)
	if sessionID == "" {
		return fmt.Errorf("session is empty")
	}
	u, err := url.Parse(fmt.Sprintf("ws://%s/live", s.conf.RemoteHost))
	if err != nil {
		return fmt.Errorf("invalid server URL: %w", err)
	}

	header := http.Header{}
	header.Set("Origin", s.conf.RemoteHost)

	cookies := s.cookie.Load().(*cookiejar.Jar).Cookies(u)
	cookieHeader := ""
	for i, c := range cookies {
		if i > 0 {
			cookieHeader += "; "
		}
		cookieHeader += c.Name + "=" + c.Value
	}
	header.Set("Cookie", cookieHeader)

	conn, _, err := websocket.DefaultDialer.DialContext(s.ctx, u.String(), header)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	// Expect initial "Password:\n"
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	// Send authentication
	if err = conn.WriteMessage(websocket.TextMessage, []byte("session_id="+sessionID)); err != nil {
		return fmt.Errorf("send session ID failed: %w", err)
	}
	if err = conn.WriteMessage(websocket.TextMessage, []byte("START 1234")); err != nil {
		return fmt.Errorf("send START command failed: %w", err)
	}

	// Stream data
	type streamData struct {
		Data struct {
			Timer  string    `json:"timer"`
			Field2 []float64 `json:"123"`
		} `json:"data"`
	}
	loc := time.Now().Location()
	for {
		_, msg, err = conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		if strings.Contains(string(msg), "settings") {
			continue
		}
		var data streamData
		if err = json.Unmarshal(msg, &data); err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}
		tms, errP := time.ParseInLocation("15:04:05", data.Data.Timer, loc)
		if errP != nil {
			return fmt.Errorf("parse time: %w", errP)
		}
		if len(data.Data.Field2) != 5 {
			return fmt.Errorf("expected 5 noise measures, got %d", len(data.Data.Field2))
		}
		s.items.Add(entities.NoiseMeasures{
			Timestamp:    tms,
			TimestampNum: utils.TimeToDayIntNum(tms),

			LAeqDT:  data.Data.Field2[0],
			LAf:     data.Data.Field2[1],
			LCPK:    data.Data.Field2[2],
			LAeqG10: data.Data.Field2[3],
			LAeqG5:  data.Data.Field2[4],
		})
	}
}
