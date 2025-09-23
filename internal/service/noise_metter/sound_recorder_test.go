package noise_metter_test

import (
	"noize_metter/internal/service/noise_metter"
	testhelpers "noize_metter/internal/test_helpers"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRecordAudio(t *testing.T) {
	container := testhelpers.GetClean(t)
	require.NoError(t, container.ServiceNoise.RecordSound(&noise_metter.RecordTask{
		StartTime: time.Now().Add(-1 * time.Minute),
		Duration:  30 * time.Second,
	}))
}
