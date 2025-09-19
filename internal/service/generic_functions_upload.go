package service

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	uploadDataDuration = 1 * time.Minute // Duration for uploading data
)

type HasTimestamp interface {
	GetTimestampNum() int64
}

func BGUploadData[T HasTimestamp](ctx context.Context, log logger.AppLogger, conf *config.AppConfig, hostURL, storageFolder string) {
	ticker := time.NewTicker(uploadDataDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			UploadData[T](log, conf, hostURL, storageFolder)
		}
	}
}

func UploadData[T HasTimestamp](log logger.AppLogger, conf *config.AppConfig, hostURL, storageFolder string) {
	entries, err := os.ReadDir(storageFolder)
	if err != nil {
		log.Fatal("error reading storage folder", err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		l := log.With(logger.WithString("file", name))
		if !strings.HasSuffix(name, ".gz") {
			continue
		}
		filePath := filepath.Join(storageFolder, name)
		data, errL := LoadChunk[T](filePath)
		if errL != nil {
			l.Error("failed to load noise data file", errL)
			continue
		}
		if errL = UploadChunk[T](l, hostURL, conf.APIKey, conf.BoxIP, filePath, data); errL != nil {
			l.Error("failed to upload chunk", errL)
			continue
		}
		if err = os.Remove(filePath); err != nil {
			l.Error("failed to remove file after upload", err)
		}
		l.Info("successfully uploaded and removed file", logger.WithInt("items", len(data)))
	}
}

func UploadChunk[T HasTimestamp](l logger.AppLogger, hostURL, apiKey, boxIP, fileName string, data []T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	filtered := make([]T, 0, len(data))
	zeroRows := 0
	for i := range data {
		if data[i].GetTimestampNum() == 0 {
			zeroRows++
			continue
		}
		filtered = append(filtered, data[i])
	}
	if zeroRows > 0 {
		l.Info("skipping data with zero timestamp",
			logger.WithInt("zero_data", zeroRows),
			logger.WithInt("valid_data", len(data)-zeroRows),
			logger.WithInt("total", len(data)),
		)
	}

	if len(filtered) == 0 {
		l.Info("no valid data to upload")
		return nil
	}

	_, code, err := utils.PostCurl[any](ctx, hostURL, map[string]any{
		"file_name": fileName,
		"source":    boxIP,
		"measures":  filtered,
	}, map[string]string{
		"Content-Type": "application/json",
		"auth-mapi":    apiKey,
	})
	if code == http.StatusOK {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to upload %d data %d: %w", len(filtered), code, err)
	}
	return fmt.Errorf("unexpected status code: %d", code)
}

func LoadChunk[T any](filePath string) ([]T, error) {
	data, err := utils.LoadFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load file %s: %w", filePath, err)
	}
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer r.Close()

	var items []T
	if err = json.NewDecoder(r).Decode(&items); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return items, nil
}
