package notificator

import (
	"time"
)

type Notificator interface {
	SendTaskErrMessage(service string, startedAt, finishedAt time.Time, errs ...error) error
	SendInfoMessage(message string, args ...string) error
}
