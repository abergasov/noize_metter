package utils

// ContainsInSlice checks if slice contains value
func ContainsInSlice[T comparable](slice []T, value T) bool {
	for _, entry := range slice {
		if entry == value {
			return true
		}
	}
	return false
}

func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	chunks := make([][]T, 0, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ExcludeFromSlice returns a new slice with all values from excludeValues removed from slice
// not optimized for performance
func ExcludeFromSlice[T comparable](slice, excludeValues []T) []T {
	result := make([]T, 0, len(slice))
	for _, entry := range slice {
		if !ContainsInSlice(excludeValues, entry) {
			result = append(result, entry)
		}
	}
	return result
}
