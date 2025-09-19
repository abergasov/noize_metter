package utils_test

import (
	"noize_metter/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsInSlice(t *testing.T) {
	t.Parallel()
	t.Run("should return true if value is in slice", func(t *testing.T) {
		// given
		slice := []string{"a", "b", "c"}
		value := "b"

		// when
		result := utils.ContainsInSlice(slice, value)

		// then
		require.True(t, result)
	})
	t.Run("should return false if value is not in slice", func(t *testing.T) {
		// given
		slice := []string{"a", "b", "c"}
		value := "d"

		// when
		result := utils.ContainsInSlice(slice, value)

		// then
		require.False(t, result)
	})
}

func TestChunkSlice(t *testing.T) {
	require.Equal(t, [][]uint64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, utils.ChunkSlice([]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
	require.Equal(t, [][]uint64{{1, 2, 3}, {4, 5, 6}, {7, 8}}, utils.ChunkSlice([]uint64{1, 2, 3, 4, 5, 6, 7, 8}, 3))
}

func TestExcludeFromSlice(t *testing.T) {
	src := []string{"a", "b", "c", "d", "e"}
	exclude := []string{"b", "d"}
	require.Equal(t, []string{"a", "c", "e"}, utils.ExcludeFromSlice(src, exclude))

	src = []string{"a", "", "", "c"}
	require.Equal(t, []string{"a", "c"}, utils.ExcludeFromSlice(src, []string{""}))
}
