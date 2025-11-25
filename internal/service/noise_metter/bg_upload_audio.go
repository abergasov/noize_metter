package noise_metter

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"noize_metter/internal/logger"
	"noize_metter/internal/utils"
	"noize_metter/internal/utils/requests"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	uploadAudioDuration = 1 * time.Minute // Duration for saving audio files
)

func (s *Service) bgUploadAudio() {
	ticker := time.NewTicker(uploadAudioDuration)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			if err := s.uploadAudioWrapper(); err != nil {
				s.log.Error("error uploading audio files", err)
			}
		}
	}
}

func (s *Service) uploadAudioWrapper() error {
	entries, err := os.ReadDir(s.conf.StorageAudioFolder)
	if err != nil {
		return fmt.Errorf("error reading audio storage folder: %w", err)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if !strings.HasSuffix(e.Name(), ".wav") {
			continue
		}
		l := s.log.With(logger.WithString("file", e.Name()))
		filePath := filepath.Join(s.conf.StorageAudioFolder, e.Name())
		if err = s.uploadAudio(filePath, e.Name()); err != nil {
			l.Error("error uploading audio file", err)
			continue
		}
		if err = os.Remove(filePath); err != nil {
			l.Error("failed to remove file after upload", err)
		}
		l.Info("successfully uploaded and removed file")
	}
	return nil
}

func (s *Service) uploadAudio(filePath, fileName string) error {
	ctx, cancel := context.WithTimeout(s.ctx, 1*time.Minute)
	defer cancel()

	hostURL := fmt.Sprintf("%s/api-mapi/v1/private/noiser/upload_wav", s.conf.DataHost)
	data, errL := utils.LoadFromFile(filePath)
	if errL != nil {
		return fmt.Errorf("error loading audio file: %w", errL)
	}
	_, code, err := requests.PostCurl[any](ctx, hostURL, map[string]any{
		"file_name": fileName,
		"source":    s.conf.BoxIP,
		"measures":  base64.StdEncoding.EncodeToString(data),
	}, map[string]string{
		"Content-Type": "application/json",
		"auth-mapi":    s.conf.APIKey,
	})
	if code == http.StatusOK {
		return nil
	}
	if err != nil {
		err = fmt.Errorf("error uploading audio file, code %d: %w", code, err)
	}
	return nil
}
