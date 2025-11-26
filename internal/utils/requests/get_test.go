package requests_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"noize_metter/internal/logger"
	testhelpers "noize_metter/internal/test_helpers"
	"noize_metter/internal/utils/requests"

	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetCurl(t *testing.T) {
	// given
	header := uuid.NewString()
	appLog := logger.NewAppSLogger()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	sampleResult := TestResponse{
		ID1: uuid.New(),
		ID2: uuid.New(),
		ID3: uuid.New(),
	}
	headers := map[string]string{
		"Custom": header,
	}
	srv := testhelpers.NewTestServer(t)
	srv.HTTPEngine.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{"john": "doe"},
	}))
	srv.RegisterHandler(http.MethodGet, "/test", func(ctx *fiber.Ctx) error {
		customHeader, ok := ctx.GetReqHeaders()["Custom"]
		require.True(t, ok)
		require.Equal(t, header, customHeader[0])
		time.Sleep(100 * time.Millisecond)
		return ctx.JSON(sampleResult)
	})
	srv.Start(t)

	t.Run("should return 401 without basic auth", func(t *testing.T) {
		// when
		res, code, err := requests.GetCurl[TestResponse](
			ctx,
			fmt.Sprintf("%s/test", srv.Address()),
			headers,
			requests.WithRequestMark("test_request"),
		)

		// then
		require.Error(t, err)
		require.Equal(t, http.StatusUnauthorized, code, "should return 401 code")
		require.Zero(t, res)
	})
	t.Run("should serve error in case of non existing address", func(t *testing.T) {
		// when
		res, code, err := requests.GetCurl[TestResponse](
			ctx,
			"http://127.0.0.1:1/test",
			headers,
			requests.WithLogger(appLog),
			requests.WithRequestMark("test_request"),
			requests.WithBasicAuth("john", "doe"),
		)

		// then
		require.ErrorContains(t, err, "connect: connection refused")
		require.Zerof(t, code, "should return 0 code")
		require.Zero(t, res, "should return empty response")
	})
	t.Run("should serve correct request", func(t *testing.T) {
		// when
		res, code, err := requests.GetCurl[TestResponse](
			ctx,
			fmt.Sprintf("%s/test", srv.Address()),
			headers,
			requests.WithRequestMark("test_request"),
			requests.WithBasicAuth("john", "doe"),
		)

		// then
		require.NoError(t, err, "should not return error")
		require.Equal(t, http.StatusOK, code, "should return 200 code")
		require.Equal(t, sampleResult, res, "should return correct response")
	})
	t.Run("should serve correct request with custom decoder", func(t *testing.T) {
		// when
		var dRes TestResponse
		_, code, err := requests.GetCurl[TestResponse](
			ctx,
			fmt.Sprintf("%s/test", srv.Address()),
			headers,
			requests.WithDecoder(func(reader io.Reader) error {
				b, e := io.ReadAll(reader)
				if e != nil {
					return e
				}
				return json.Unmarshal(b, &dRes)
			}),
			requests.WithRequestMark("test_request"),
			requests.WithBasicAuth("john", "doe"),
		)

		// then
		require.NoError(t, err, "should not return error")
		require.Equal(t, http.StatusOK, code, "should return 200 code")
		require.Equal(t, sampleResult, dRes, "should return correct response")
	})
	t.Run("should cancel request in case small timeout", func(t *testing.T) {
		// when
		res, code, err := requests.GetCurl[TestResponse](
			ctx,
			fmt.Sprintf("%s/test", srv.Address()),
			headers,
			requests.WithLogger(appLog),
			requests.WithRequestMark("test_request"),
			requests.WithSkipTLSVerification(true),
			requests.WithRequestTimeout(10*time.Millisecond),
			requests.WithBasicAuth("john", "doe"),
		)

		// then
		require.Error(t, err, "should return error")
		require.Zero(t, code, "should return 0 code")
		require.Zero(t, res, "should return empty response")
	})
}
