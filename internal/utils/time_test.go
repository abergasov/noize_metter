package utils_test

import (
	"noize_metter/internal/utils"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeToDayInt(t *testing.T) {
	table := map[string]string{
		"2024-05-22 14:30:00": "20240522",
		"2024-05-21 14:30:00": "20240521",
		"2024-12-31 23:59:59": "20241231",
	}
	for timeStr, expectedTime := range table {
		parsedTime, err := time.Parse(time.DateTime, timeStr)
		require.NoError(t, err)
		require.Equal(t, expectedTime, utils.TimeToDayInt(parsedTime))
		expectedTimeInt, err := strconv.ParseInt(expectedTime, 10, 64)
		require.NoError(t, err)
		require.Equal(t, expectedTimeInt, utils.TimeToDayIntNum(parsedTime))
	}
}

func TestRoundToNearest5Minutes(t *testing.T) {
	testCases := map[string]string{
		"2024-01-12 14:14:23": "2024-01-12 14:10:00",
		"2024-01-12 14:26:47": "2024-01-12 14:25:00",
		"2024-01-12 14:30:00": "2024-01-12 14:30:00",
		"2024-01-12 14:34:59": "2024-01-12 14:30:00",
		"2024-01-12 14:00:00": "2024-01-12 14:00:00",
		"2024-01-12 14:07:15": "2024-01-12 14:05:00",
		"2024-01-12 14:12:45": "2024-01-12 14:10:00",
		"2024-01-12 14:55:10": "2024-01-12 14:55:00",
		"2024-01-12 14:59:59": "2024-01-12 14:55:00",
		"2024-01-12 15:00:00": "2024-01-12 15:00:00",
		"2024-01-12 23:59:59": "2024-01-12 23:55:00",
		"2024-01-12 00:00:00": "2024-01-12 00:00:00",
		"2024-01-12 00:02:30": "2024-01-12 00:00:00",
		"2024-01-12 00:03:30": "2024-01-12 00:00:00",
		"2024-01-12 00:04:30": "2024-01-12 00:00:00",
		"2024-01-12 00:05:30": "2024-01-12 00:05:00",
		"2024-01-12 12:45:00": "2024-01-12 12:45:00",
		"2024-01-12 12:46:00": "2024-01-12 12:45:00",
		"2024-01-12 12:47:00": "2024-01-12 12:45:00",
		"2024-01-12 12:48:00": "2024-01-12 12:45:00",
		"2024-01-12 12:49:00": "2024-01-12 12:45:00",
	}

	for inputStr, expectedStr := range testCases {
		input, err := time.Parse("2006-01-02 15:04:05", inputStr)
		require.NoError(t, err)
		expected, err := time.Parse("2006-01-02 15:04:05", expectedStr)
		require.NoError(t, err)

		require.Equalf(t, expected, utils.RoundToNearest5Minutes(input), "failed for %s", expectedStr)
	}
}
