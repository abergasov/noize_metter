package ces_test

import (
	testhelpers "noize_metter/internal/test_helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_RunIteration(t *testing.T) {
	container := testhelpers.GetClean(t)
	require.Empty(t, container.ServiceCes.RunIteration())
}
