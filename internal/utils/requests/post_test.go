package requests_test

import (
	"context"
	"fmt"
	"net/http"
	testhelpers "noize_metter/internal/test_helpers"
	"noize_metter/internal/utils/requests"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPostCurl(t *testing.T) {
	// given
	header := uuid.NewString()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	sampleResult := TestResponse{
		ID1: uuid.New(),
		ID2: uuid.New(),
		ID3: uuid.New(),
	}
	sampleRequest := TestRequest{
		ID1: uuid.New(),
		ID2: uuid.New(),
		ID3: uuid.New(),
	}
	headers := map[string]string{
		"Custom": header,
	}
	srv := testhelpers.NewTestServer(t)
	srv.RegisterHandler(http.MethodPost, "/test", func(ctx *fiber.Ctx) error {
		require.Equal(t, header, ctx.GetReqHeaders()["Custom"][0])
		var req TestRequest
		require.NoError(t, ctx.BodyParser(&req))
		require.Equal(t, sampleRequest, req)
		return ctx.JSON(sampleResult)
	})
	srv.Start(t)
	t.Run("should serve error in case of non existing address", func(t *testing.T) {
		// when
		res, code, err := requests.PostCurl[TestResponse](ctx, "http://127.0.0.1:1/test", sampleRequest, headers)

		// then
		require.ErrorContains(t, err, "connect: connection refused")
		require.Zerof(t, code, "should return 0 code")
		require.Zero(t, res, "should return empty response")
	})
	t.Run("should serve correct request", func(t *testing.T) {
		// when
		res, code, err := requests.PostCurl[TestResponse](ctx, fmt.Sprintf("%s/test", srv.Address()), sampleRequest, headers)

		// then
		require.NoError(t, err, "should not return error")
		require.Equal(t, http.StatusOK, code, "should return 200 code")
		require.Equal(t, sampleResult, res, "should return correct response")
	})
}
