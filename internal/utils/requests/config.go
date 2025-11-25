package requests

import (
	"io"
	"noize_metter/internal/logger"
	"time"
)

type requestBasicAuth struct {
	username string
	password string
}

type Config struct {
	requestMark         string
	logger              logger.AppLogger
	requestTimeout      time.Duration
	skipTLSVerification bool
	basicAuth           *requestBasicAuth
	Decoder             func(reader io.Reader) error
}
type Option func(*Config)

func NewDefaultConfig() *Config {
	return &Config{
		requestTimeout:      30 * time.Second,
		skipTLSVerification: false,
	}
}

func WithDecoder(decoder func(io.Reader) error) func(*Config) {
	return func(c *Config) {
		c.Decoder = decoder
	}
}

func WithRequestMark(mark string) Option {
	return func(c *Config) {
		c.requestMark = mark
	}
}

func WithLogger(l logger.AppLogger) Option {
	return func(c *Config) {
		c.logger = l
	}
}

func WithSkipTLSVerification(skip bool) Option {
	return func(c *Config) {
		c.skipTLSVerification = skip
	}
}

func WithRequestTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.requestTimeout = timeout
	}
}

func WithBasicAuth(user, pass string) Option {
	return func(c *Config) {
		c.basicAuth = &requestBasicAuth{
			username: user,
			password: pass,
		}
	}
}
