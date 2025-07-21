package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc64"
	"io"
	"os"
	"path/filepath"
)

// AtomicallySaveToFile saves the given data to the given file atomically
// append a checksum to the data before writing it
// file is either fully updated or not updated at all
// achieved by writing to a temporary file and renaming it to the target file
// renaming is atomic operation at least on linux.
func AtomicallySaveToFile(fileName string, data []byte) error {
	checkSum := crc64.New(crc64.MakeTable(crc64.ISO))
	if _, err := checkSum.Write(data); err != nil {
		return fmt.Errorf("cannot calculate checksum: %w", err)
	}

	checkSumBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(checkSumBytes, checkSum.Sum64())
	resultData := make([]byte, 0, len(data)+len(checkSumBytes))
	resultData = append(resultData, checkSumBytes...)
	resultData = append(resultData, data...)

	dir, file := filepath.Split(fileName)
	if dir == "" {
		dir = "."
	}

	f, err := os.CreateTemp(dir, file)
	if err != nil {
		return fmt.Errorf("cannot create temp file: %w", err)
	}
	defer func() {
		if err != nil {
			// Don't leave the temp file lying around on error.
			_ = os.Remove(f.Name()) // yes, ignore the error, not much we can do about it.
		}
	}()
	defer f.Close() //nolint

	name := f.Name()
	if _, err = io.Copy(f, bytes.NewReader(resultData)); err != nil {
		return fmt.Errorf("cannot write data to tempfile %q: %w", name, err)
	}
	// fsync is important, otherwise os.Rename could rename a zero-length file
	if err = f.Sync(); err != nil {
		return fmt.Errorf("can't flush tempfile %q: %w", name, err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("can't close tempfile %q: %w", name, err)
	}

	// get the file mode from the original file and use that for the replacement file, too.
	destInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		// no original file
	} else if err != nil {
		return err
	} else {
		sourceInfo, errS := os.Stat(name)
		if errS != nil {
			return errS
		}

		if sourceInfo.Mode() != destInfo.Mode() {
			if err = os.Chmod(name, destInfo.Mode()); err != nil {
				return fmt.Errorf("can't set filemode on tempfile %q: %w", name, err)
			}
		}
	}
	if err = os.Rename(name, fileName); err != nil {
		return fmt.Errorf("cannot replace %q with tempfile %q: %w", fileName, name, err)
	}
	return nil
}

// LoadFromFile loads the data from the given file and verifies the checksum. It returns the data without the checksum
func LoadFromFile(path string) (data []byte, err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return nil, os.ErrNotExist
	}
	r, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer r.Close() //nolint

	data, err = io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(data) < 8 {
		return nil, errors.New("file is too short")
	}
	fileCrc := binary.LittleEndian.Uint64(data[:8])
	dataCrc := crc64.Checksum(data[8:], crc64.MakeTable(crc64.ISO))
	if fileCrc != dataCrc {
		return nil, fmt.Errorf("checksum mismatch: %x != %x", fileCrc, dataCrc)
	}
	return data[8:], nil
}
