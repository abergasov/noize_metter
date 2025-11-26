package logger

import (
	"io"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"
	"sync"
)

type SLogger struct {
	logWriters []*slog.Logger // since we not modify this slice, we able avoid mutex usage
}

var _ AppLogger = (*SLogger)(nil)

func replacer(_ []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case "msg":
		return slog.String("message", a.Value.String())
	case "err":
		return slog.String("error", a.Value.String())
	}
	return a
}

func getLastCommitHash() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	data := strings.Split(strings.ReplaceAll(info.Main.Version, "+dirty", ""), "-")
	res := data[len(data)-1]
	if len(res) > 7 {
		return res[:7]
	}
	return res
}

// NewAppSLogger creates a new logger instance which able to write to multiple writers.
// By default, if Config.Writers are empty, it will write to stdout.
func NewAppSLogger(args ...slog.Attr) *SLogger {
	writers := []io.Writer{os.Stdout}
	logs := make([]*slog.Logger, 0, len(writers))
	for _, w := range writers {
		var handler slog.Handler
		handler = slog.NewJSONHandler(w, &slog.HandlerOptions{
			ReplaceAttr: replacer,
		})
		attrs := make([]any, 0, len(args)+1)
		for _, arg := range args {
			attrs = append(attrs, arg)
		}

		if commitHash := getLastCommitHash(); commitHash != "" {
			attrs = append(attrs, slog.String("commit", commitHash))
		}

		lw := slog.New(handler).With(attrs...)
		logs = append(logs, lw)
	}
	return &SLogger{logWriters: logs}
}

func (l *SLogger) Info(message string, args ...slog.Attr) {
	params := prepareSlogParams(nil, args)
	l.processWriters(func(lg *slog.Logger) {
		lg.Info(message, params...)
	})
}

func (l *SLogger) Error(message string, err error, args ...slog.Attr) {
	params := prepareSlogParams(err, args)
	l.processWriters(func(lg *slog.Logger) {
		lg.Error(message, params...)
	})
}

func (l *SLogger) Fatal(message string, err error, args ...slog.Attr) {
	l.Error(message, err, args...)
	os.Exit(1)
}

func (l *SLogger) Warn(message string, args ...slog.Attr) {
	params := prepareSlogParams(nil, args)
	l.processWriters(func(lg *slog.Logger) {
		lg.Warn(message, params...)
	})
}

func (l *SLogger) With(args ...slog.Attr) AppLogger {
	logs := make([]*slog.Logger, 0, len(l.logWriters))
	for _, lg := range l.logWriters {
		logs = append(logs, lg.With(prepareSlogParams(nil, args)...))
	}
	return &SLogger{logWriters: logs}
}

func (l *SLogger) processWriters(processor func(*slog.Logger)) {
	var wg sync.WaitGroup
	wg.Add(len(l.logWriters))
	for i := range l.logWriters {
		go func(j int) {
			processor(l.logWriters[j])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func prepareSlogParams(err error, args []slog.Attr) []any {
	params := make([]any, 0, len(args)+1)
	if err != nil {
		params = append(params, err)
	}
	argsMap := make(map[string]slog.Attr, len(args))
	for _, arg := range args {
		argsMap[arg.Key] = arg
	}
	for _, arg := range argsMap {
		params = append(params, arg)
	}
	return params
}
