package noise_metter

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"noize_metter/internal/entities"
	"regexp"
	"strings"

	"github.com/gorilla/websocket"
)

var (
	reg = regexp.MustCompile("[^a-zA-Z0-9.-]+")
)

func (s *Service) ScrapeWeatherSensorData() error {
	sessionID, _ := s.session.Load().(string)
	if sessionID == "" {
		return fmt.Errorf("session is empty")
	}
	u, err := url.Parse(fmt.Sprintf("ws://%s/websocket/data/", s.conf.RemoteHost))
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

	authMessage := make([]byte, 0, 54)
	authMessage = append(authMessage, 0, 150, 0, 0)
	authData, _ := json.Marshal(map[string]string{"session_id": sessionID})
	authMessage = append(authMessage, authData...)

	if err = conn.WriteMessage(websocket.TextMessage, []byte(base64.StdEncoding.EncodeToString(authMessage))); err != nil {
		return fmt.Errorf("send session ID failed: %w", err)
	}

	for {
		_, msg, errR := conn.ReadMessage()
		if errR != nil {
			return fmt.Errorf("read: %w", errR)
		}

		result := make([]byte, 0, 5+3)
		result = append(result, msg[:5]...)
		result = append(result, 'A', '=', '=')
		if err = conn.WriteMessage(websocket.TextMessage, result); err != nil {
			return fmt.Errorf("send command failed: %w", err)
		}
		if len(msg) > 40 {
			strMsg := string(msg)
			res, errP := s.ParseWeatherSensorData(strMsg)
			if errP != nil {
				continue
			}
			println(res.String())
		}
	}
}

func (s *Service) ParseWeatherSensorData(base64Data string) (*entities.NoiseWeather, error) {
	for len(base64Data)%4 != 0 {
		base64Data += "="
	}
	raw, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding failed: %w", err)
	}
	z := findZlib(raw)
	if z == nil {
		return nil, fmt.Errorf("zlib decoding failed")
	}

	plain := inflate(z)
	data := strings.Split(reg.ReplaceAllString(string(plain), `%`), "%")
	if data[1] != "Vaisala" {
		// something went wrong
		return nil, fmt.Errorf("not Vaisala data")
	}
	return &entities.NoiseWeather{
		AirPressure:    getData(data, 8),
		AirTemperature: getData(data, 6),
		WindSpeedAvg:   getData(data, 4),
		WindDirection:  getData(data, 5),
		RelHumidity:    getData(data, 7),
		RainIntensity:  getData(data, 9),
		HailIntensity:  getData(data, 10),
	}, nil
}

func getData(b []string, index int) string {
	if index >= len(b) {
		return "0.0"
	}
	return b[index]
}

func findZlib(b []byte) []byte {
	i := bytes.Index(b, []byte{0x78, 0x9c})
	if i < 0 {
		i = bytes.Index(b, []byte{0x78, 0xDA})
	}
	if i < 0 {
		return nil
	}
	return b[i:]
}

func inflate(b []byte) []byte {
	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil
	}
	defer r.Close()
	out, _ := io.ReadAll(r)
	return out
}
