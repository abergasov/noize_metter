package noise_metter_test

import (
	"noize_metter/internal/entities"
	testhelpers "noize_metter/internal/test_helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScrapeWeatherSensorData(t *testing.T) {
	container := testhelpers.GetClean(t)
	require.NoError(t, container.ServiceNoise.Auth())
	require.NoError(t, container.ServiceNoise.ScrapeWeatherSensorData())
}

func TestParseWeatherSensorData(t *testing.T) {
	cases := map[string]entities.NoiseWeather{
		"B6VlAHicY2BgYGAE4otMEBqEuRggICwxszgxJ5GB3gDkhggjS1MzQ0tjetsNDAaG0GAnelsLBsxAbKRnSnd7ZYCYBWS3qRHd7ZYDYlYgNjTWM6C33dxAzAbExpZ6FvS2mx2KDQ0MLPToHOzyQMwBxAb0D3IGBSDmHCC7lehu4+ABAHTIDDg=": {
			WindSpeedAvg:   "2.5",
			WindDirection:  "252",
			AirTemperature: "13.0",
			RelHumidity:    "39.8",
			AirPressure:    "1008.2",
			RainIntensity:  "0.0",
			HailIntensity:  "0.0",
		},
		"B6U/AHicY2BgYGAE4otMEBqEuRggICwxszgxJ5GB3gDkhggjS1MzQ0tjetsNDAaG0GAnelsLBsxAbKBnSnd7ZYCYBYiNTAzpbrccELMCsa6BngW97eYGYjYgNjQw0DOgs93sUAy021LPkr52ywMxBwMordHb1wwMCkDMOUB2K9HdxsEDAM3/DFo=": {
			WindSpeedAvg:   "0.5",
			WindDirection:  "241",
			AirTemperature: "-0.8",
			RelHumidity:    "100.0",
			AirPressure:    "1009.9",
			RainIntensity:  "0.0",
			HailIntensity:  "0.0",
		},
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
