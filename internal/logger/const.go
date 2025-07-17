package logger

import (
	_ "embed"
)

func WithMethod(method string) StringWith {
	return StringWith{Key: "_method", Val: method}
}

func WithService(serviceName string) StringWith {
	return StringWith{Key: "_service", Val: serviceName}
}
