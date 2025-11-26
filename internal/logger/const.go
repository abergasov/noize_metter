package logger

import (
	"log/slog"
	"time"
)

// this file is used to define helper function for logger
// main purpose is to have standardization for log output
// this will allow us to use indexer to search for specific field

const (
	SlogFlagRequestType  = "request_type"
	SlogFlagRemoteTarget = "remote_target"
	SlogFlagLatency      = "latency"
)

func WithString(key, val string) slog.Attr {
	return slog.String(key, val)
}

func WithUnt64(key string, val uint64) slog.Attr {
	return slog.Uint64(key, val)
}

func WithInt64(key string, val int64) slog.Attr {
	return slog.Int64(key, val)
}

func WithFloat64(key string, val float64) slog.Attr {
	return slog.Float64(key, val)
}

func WithInt(key string, val int) slog.Attr {
	return slog.Int(key, val)
}

func WithService(serviceName string) slog.Attr {
	return WithString("service_name", serviceName)
}

func WithDuration(startTime time.Time) slog.Attr {
	return slog.Int64("duration_ms", time.Since(startTime).Milliseconds())
}

func WithExternalRPC(name string) slog.Attr {
	return slog.String("external_rpc", name)
}

func WithHTTPRequest() slog.Attr {
	return slog.String("transport_type", "http_request")
}

func WithGRPCRequest() slog.Attr {
	return slog.String("transport_type", "grpc_request")
}

func WithLatencyFlag() slog.Attr {
	return slog.String(SlogFlagLatency, "true")
}

func WithRemoteTarget(target string) slog.Attr {
	return slog.String(SlogFlagRemoteTarget, target)
}

func WithFunctionName(funcName string) slog.Attr {
	return slog.String("function_name", funcName)
}

func WithPackage(packageName string) slog.Attr {
	return slog.String("package", packageName)
}

func WithJobID(jobID string) slog.Attr {
	return slog.String("job_id", jobID)
}

func WithFileName(fileName string) slog.Attr {
	return slog.String("file_name", fileName)
}

func WithFilePath(fileName string) slog.Attr {
	return slog.String("file_path", fileName)
}

func WithWorkerGroup(workerGroup string) slog.Attr {
	return slog.String("worker_group", workerGroup)
}

func WithPoolCode(poolCode string) slog.Attr {
	return slog.String("pool_code", poolCode)
}

func WithCoinS(coinName string) slog.Attr {
	return slog.String("coin", coinName)
}

func WithNodeIndex(nodeIndex int64) slog.Attr {
	return slog.Int64("nodeindex", nodeIndex)
}

func WithHeight(height uint64) slog.Attr {
	return slog.Uint64("height", height)
}

func WithHash(hash string) slog.Attr {
	return slog.String("hash", hash)
}

func WithPreviousBlockHash(prevHash string) slog.Attr {
	return slog.String("previous_block_hash", prevHash)
}
