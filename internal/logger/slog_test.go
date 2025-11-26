package logger_test

import (
	"fmt"
	"log/slog"
	"noize_metter/internal/logger"

	"os"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_SLogger_purelog_with_stdout(t *testing.T) {
	// given
	require.NoError(t, os.Setenv("debug", "true"))
	appLog := logger.NewAppSLogger()
	appLog.Info("test")

	// when, then
	concurrentlyLogIt(appLog.With(logger.WithDuration(time.Now()), logger.WithExternalRPC("test")))
}

func Test_SLogger_multiply_writers(t *testing.T) {
	// given
	appLog := logger.NewAppSLogger(slog.String("additional", "value"), slog.String("additional2", "value"))

	// when, then
	concurrentlyLogIt(appLog)
}

func Test_SLoggerCheckConstants(t *testing.T) {
	// given
	appLog := logger.NewAppSLogger()

	// when
	l := appLog.With(
		logger.WithDuration(time.Now()),
		logger.WithPreviousBlockHash(uuid.NewString()),
		logger.WithHash(uuid.NewString()),
		logger.WithRemoteTarget(uuid.NewString()),
		logger.WithPoolCode(uuid.NewString()),
		logger.WithFunctionName(uuid.NewString()),
		logger.WithPackage(uuid.NewString()),
		logger.WithJobID(uuid.NewString()),
		logger.WithFileName(uuid.NewString()),
		logger.WithWorkerGroup(uuid.NewString()),
		logger.WithCoinS("BSV"),
		logger.WithNodeIndex(1),
		logger.WithHeight(2),
	)

	// then
	l.Info("test")
}

func concurrentlyLogIt(appLog logger.AppLogger) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				appLog.Info("message")
				appLog.Error("message", fmt.Errorf("error"))
				appLog.Error("message", fmt.Errorf("error"), slog.String("key", uuid.NewString()))
			}
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			customLogger1 := appLog.With(slog.String("keyA", uuid.NewString()))
			customLogger1.Info("customLogger message")
			customLogger1.Error("customLogger message", fmt.Errorf("customLogger error"))
			customLogger1.Error("customLogger message", fmt.Errorf("customLogger error"), slog.String("keyB", uuid.NewString()))
		}()
	}
	wg.Wait()
}
