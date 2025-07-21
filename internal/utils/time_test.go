package utils_test

import (
	"github.com/stretchr/testify/require"
	"noize_metter/internal/utils"
	"strconv"
	"testing"
	"time"
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
