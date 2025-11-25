package testhelpers

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/require"
)

type TestServer struct {
	address    string
	HTTPEngine *fiber.App
}

// NewTestServer creates a new test server for testing http requests.
func NewTestServer(t *testing.T) *TestServer {
	appPort := GetFreePort(t)
	fakeServer := &TestServer{
		address: fmt.Sprintf(":%d", appPort),
		HTTPEngine: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
	}
	fakeServer.HTTPEngine.Use(recover.New())
	return fakeServer
}

func (ts *TestServer) Address() string {
	return fmt.Sprintf("http://127.0.0.1%s", ts.address)
}

func (ts *TestServer) RegisterHandler(method, path string, handler func(ctx *fiber.Ctx) error) {
	ts.HTTPEngine.Add(method, path, handler)
}

func (ts *TestServer) Start(t *testing.T) {
	go func() {
		require.NoError(t, ts.HTTPEngine.Listen(ts.address))
	}()
	t.Cleanup(func() {
		require.NoError(t, ts.HTTPEngine.Shutdown())
	})
}
