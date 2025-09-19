package noise_metter

import (
	"noize_metter/internal/logger"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Service) bgPruneOldFiles() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.pruneOldFiles()
		}
	}
}

func (s *Service) pruneOldFiles() {
	cutoff := time.Now().Add(-6 * 24 * time.Hour)
	entries, err := os.ReadDir(s.conf.StorageNoiseFolder)
	if err != nil {
		s.log.Fatal("failed to read directory", err, logger.WithString("folder", s.conf.StorageNoiseFolder))
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		l := s.log.With(logger.WithString("file", name))
		if !strings.HasSuffix(name, ".gz") {
			continue
		}
		parts := strings.Split(name, "_")
		if len(parts) < 2 {
			l.Info("skipping file with unexpected name format")
			continue
		}

		t, errP := time.Parse("20060102T150405", parts[0]) // e.g. 20231001T150405
		if errP != nil {
			l.Error("skipping file with unexpected date format", err)
			continue
		}
		if t.Before(cutoff) {
			fullPath := filepath.Join(s.conf.StorageNoiseFolder, name)
			if err = os.Remove(fullPath); err != nil {
				l.Error("failed to remove old file", err)
				continue
			}
			l.Info("removed old file")
		}
	}
}
