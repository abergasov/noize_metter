package utils_test

import (
	"noize_metter/internal/utils"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAtomicallySaveLoadFile(t *testing.T) {
	t.Run("should save and load data atomically", func(t *testing.T) {
		// given
		tmpDir := prepareDir(t)

		// when
		testFile := filepath.Join(tmpDir, "testfile")
		testData := []byte("Hello, World!")
		require.NoError(t, utils.AtomicallySaveToFile(testFile, testData))

		// then
		_, err := os.Stat(testFile)
		require.NoError(t, err)

		loadedData, err := utils.LoadFromFile(testFile)
		require.NoError(t, err)
		require.Equal(t, testData, loadedData)
	})
	t.Run("should serve error in case of bad checksum", func(t *testing.T) {
		// given
		tmpDir := prepareDir(t)
		testFile := filepath.Join(tmpDir, "testfile")
		testData := []byte("Correct Data")
		require.NoError(t, utils.AtomicallySaveToFile(testFile, testData))

		// when
		f, err := os.OpenFile(testFile, os.O_WRONLY, 0o644)
		require.NoError(t, err)
		_, err = f.WriteAt([]byte("Bad Data"), 8) // overwrite part of the data
		require.NoError(t, err)
		require.NoError(t, f.Close())

		// then
		data, err := utils.LoadFromFile(testFile)
		require.Error(t, err)
		require.Contains(t, err.Error(), "checksum mismatch")
		require.Nil(t, data)
	})
	t.Run("should serve error in case of file too short", func(t *testing.T) {
		// given
		nonExistentFile := "/path/to/nonexistent/file"

		// when
		data, err := utils.LoadFromFile(nonExistentFile)

		// then
		require.ErrorIs(t, err, os.ErrNotExist)
		require.Nil(t, data)
	})
	t.Run("should serve error in case of content too short", func(t *testing.T) {
		// given
		tmpDir := prepareDir(t)

		// when
		testFile := filepath.Join(tmpDir, "shortfile")
		err := os.WriteFile(testFile, []byte{0x01, 0x02, 0x03}, 0o644) // create a file shorter than the checksum size
		require.NoError(t, err)

		// then
		_, err = utils.LoadFromFile(testFile)
		require.Error(t, err)
		require.Contains(t, err.Error(), "file is too short")
	})
	t.Run("should not able save in case of permissions issue", func(t *testing.T) {
		// given
		tmpDir := prepareDir(t)
		testFile := filepath.Join(tmpDir, "permissionsfile")
		require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "permissionsfile"), []byte{}, 0o644))
		require.NoError(t, os.Chmod(testFile, 0o600))

		// when
		require.NoError(t, utils.AtomicallySaveToFile(testFile, []byte("permission test data")))

		// then
		info, err := os.Stat(testFile)
		require.NoError(t, err)
		require.Equal(t, os.FileMode(0o600), info.Mode().Perm())
	})
}

func prepareDir(t *testing.T) string {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", uuid.NewString())
	require.NoError(t, err, "failed to create temp directory")
	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(tmpDir))
	})
	return tmpDir
}
