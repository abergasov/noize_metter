package ces_test

import (
	"context"
	"noize_metter/internal/service/ces"
	testhelpers "noize_metter/internal/test_helpers"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestService_CreateToken(t *testing.T) {
	container := testhelpers.GetClean(t)
	urlList := append(ces.URLListIMDC1, ces.URLListIMDC2...)
	for _, host := range urlList {
		println("--- testing", host)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		token, err := container.ServiceCes.CreateToken(ctx, host)
		if err != nil {
			t.Logf("failed to get token for host %s: %s", host, err.Error())
		}
		println("token:", token)
	}
	require.NoError(t, container.ServiceCes.FillToken())
	tanks, err := container.ServiceCes.GetAllTanks(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, tanks)
}
