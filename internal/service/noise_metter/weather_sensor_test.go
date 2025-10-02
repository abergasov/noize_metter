package noise_metter_test

import (
	"noize_metter/internal/entities"
	testhelpers "noize_metter/internal/test_helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseWeatherSensorData(t *testing.T) {
	cases := map[string]entities.NoiseWeather{
		"B6VSAHicY2BgYGAE4otMEBqEuRggICwxszgxJ5GB3gDkhggjS1MzQ0tjetsNDAaG0GAnelsLBsxAbKRnQnd7ZYCYBYiNjczobrccELMCsaWeAd3t5gZiNiA2N9Kju8fZodjQwMBYz5C": {
			WindSpeedAvg:   "2.4",
			WindDirection:  "326",
			AirTemperature: "9.0",
			RelHumidity:    "72.6",
			AirPressure:    "1003.1",
			RainIntensity:  "0.0",
			HailIntensity:  "0.0",
		},
		"B6VnAHicY2BgYGAE4otMEBqEuRggICwxszgxJ5GB3gDkhggjS1MzQ0tjetsNDAaG0GAnelsLBsxAbKBnQnd7ZYCYBYiNjOhvtxwQswKxroGeIb3t5gZiNiA2NDDQM6Cz3exQDLTbUo/OqVweiDkYQGmN3r5mYFAAYs4BsluJ7jYOHgAAp6QMTQ==": {
			WindSpeedAvg:   "0.4",
			WindDirection:  "224",
			AirTemperature: "-0.1",
			RelHumidity:    "100.0",
			AirPressure:    "1009.3",
			RainIntensity:  "0.0",
			HailIntensity:  "0.0",
		},
		"B6VpAHiczdBLDoIwAIThARFEr2DAuKdpeZh06xUMhG2XJuw4qTdiEM4wOsmXLv+2ABDRJ97O1RnbhvCewxSg3nqHsfbdw/lG3eY3oH891dnvDuSMlXevlFDdOHm7oCNV1sjjF0rJWSv/9WzHtjettl3SifSvBm6U/6h9lxf/ZwuVVAxJ": {
			WindSpeedAvg:   "1.0",
			WindDirection:  "231",
			AirTemperature: "-0.1",
			RelHumidity:    "100.0",
			AirPressure:    "1009.4",
			RainIntensity:  "0.0",
			HailIntensity:  "0.0",
		},
	}
	container := testhelpers.GetClean(t)
	for s, expected := range cases {
		result, err := container.ServiceNoise.ParseWeatherSensorData(s)
		require.NoError(t, err)

		require.Equal(t, expected.WindSpeedAvg, result.WindSpeedAvg)
		require.Equal(t, expected.WindDirection, result.WindDirection)
		require.Equal(t, expected.AirTemperature, result.AirTemperature)
		require.Equal(t, expected.RelHumidity, result.RelHumidity)
		require.Equal(t, expected.AirPressure, result.AirPressure)
		require.Equal(t, expected.RainIntensity, result.RainIntensity)
		require.Equal(t, expected.HailIntensity, result.HailIntensity)
	}
}
