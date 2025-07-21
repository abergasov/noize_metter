package repository

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"noize_metter/internal/entities"
	"noize_metter/internal/utils"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func (r *Repo) DumpNoiseRawData(items []entities.NoiseMeasures) error {
	return saveItems(r.conf.StorageFolder, "noise_raw", items)
}

func saveItems[T any](storageFolder, postfix string, items []T) error {
	chunks := utils.ChunkSlice(items, 500)
	for i := range chunks {
		filePath := getFileName(storageFolder, postfix)
		if err := saveChunk(filePath, chunks[i]); err != nil {
			return fmt.Errorf("failed to save chunk %d to file %s: %w", i, filePath, err)
		}
	}
	return nil
}

func saveChunk[T any](filePath string, items []T) error {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if err := json.NewEncoder(gz).Encode(items); err != nil {
		return fmt.Errorf("failed to encode items: %w", err)
	}
	if err := gz.Close(); err != nil {
		return fmt.Errorf("failed to close gzip writer: %w", err)
	}

	if err := utils.AtomicallySaveToFile(filePath, buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write gzip file %s: %w", filePath, err)
	}
	return nil
}

func getFileName(storageFolder, postfix string) string {
	id := uuid.NewString()[:8]
	fileName := fmt.Sprintf("%s_%s_%s.json.gz", time.Now().UTC().Format("20060102T150405"), id, postfix)
	return filepath.Join(storageFolder, fileName)
}
