package notificator_test

import (
	testhelpers "noize_metter/internal/test_helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlackNotificator_SendErrMessage(t *testing.T) {
	container := testhelpers.GetClean(t)
	require.NoError(t, container.ServiceSlackNotificator.SendInfoMessage("test message", "a", "b", "c"))
}
