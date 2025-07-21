package noise_metter_test

import (
	testhelpers "noize_metter/internal/test_helpers"
	"testing"
)

func TestService(t *testing.T) {
	container := testhelpers.GetClean(t)
	container.ServiceNoise.Run()
}
