package entities

type NoiseWeather struct {
	WindSpeedAvg   string `json:"wind_speed_avg"`
	WindDirection  string `json:"wind_direction"`
	AirTemperature string `json:"air_temperature"`
	RelHumidity    string `json:"rel_humidity"`
	AirPressure    string `json:"air_pressure"`
	RainIntensity  string `json:"rain_intensity"`
	HailIntensity  string `json:"hail_intensity"`
}
