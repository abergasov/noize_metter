package noise_metter_test

import (
	"fmt"
	testhelpers "noize_metter/internal/test_helpers"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServiceTime(t *testing.T) {
	tt := time.Now().Format(time.DateOnly)
	tmsS := "12:22:32"
	tms, errP := time.Parse(time.DateTime, fmt.Sprintf("%s %s", tt, tmsS))
	require.NoError(t, errP)
	t.Logf("tms: %v", tms.Format(time.DateTime))
}

func TestService(t *testing.T) {
	container := testhelpers.GetClean(t)
	container.ServiceNoise.Run()
}
