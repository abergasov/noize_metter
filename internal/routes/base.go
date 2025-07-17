package routes

import (
	"noize_metter/internal/logger"
	"noize_metter/internal/service/noise_metter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	appAddr      string
	log          logger.AppLogger
	serviceNoise *noise_metter.Service
	httpEngine   *fiber.App
}

// InitAppRouter initializes the HTTP Server.
func InitAppRouter(log logger.AppLogger, serviceNoise *noise_metter.Service, address string) *Server {
	app := &Server{
		appAddr:      address,
		httpEngine:   fiber.New(fiber.Config{}),
		serviceNoise: serviceNoise,
		log:          log.With(logger.WithService("http")),
	}
	app.httpEngine.Use(recover.New())
	app.initRoutes()
	return app
}

func (s *Server) initRoutes() {
	s.httpEngine.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
}

// Run starts the HTTP Server.
func (s *Server) Run() error {
	s.log.Info("Starting HTTP server", logger.WithString("port", s.appAddr))
	return s.httpEngine.Listen(s.appAddr)
}

func (s *Server) Stop() error {
	return s.httpEngine.Shutdown()
}
