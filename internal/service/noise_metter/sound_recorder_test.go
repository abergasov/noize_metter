package noise_metter_test

import (
	testhelpers "noize_metter/internal/test_helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecordAudio(t *testing.T) {
	container := testhelpers.GetClean(t)
	require.NoError(t, container.ServiceNoise.RecordSound())
}
