package noise_metter

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"noize_metter/internal/entities"
	"regexp"
	"strings"
)

var (
	reg = regexp.MustCompile("[^a-zA-Z0-9.-]+")
)

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
