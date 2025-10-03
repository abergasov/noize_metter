package entities

import (
	"fmt"
	"time"
)

type NoiseWeather struct {
	Timestamp      time.Time `json:"timestamp"`
	TimestampNum   int64     `json:"timestamp_num"`
	WindSpeedAvg   string    `json:"wind_speed_avg"`
	WindDirection  string    `json:"wind_direction"`
	AirTemperature string    `json:"air_temperature"`
	RelHumidity    string    `json:"rel_humidity"`
	AirPressure    string    `json:"air_pressure"`
	RainIntensity  string    `json:"rain_intensity"`
	HailIntensity  string    `json:"hail_intensity"`
}

func (n *NoiseWeather) String() string {
	return fmt.Sprintf(
		"WindSpeedAvg: %s, WindDirection: %s, AirTemperature: %s, RelHumidity: %s, AirPressure: %s, RainIntensity: %s, HailIntensity: %s",
		n.WindSpeedAvg,
		n.WindDirection,
		n.AirTemperature,
		n.RelHumidity,
		n.AirPressure,
		n.RainIntensity,
		n.HailIntensity,
	)
}
