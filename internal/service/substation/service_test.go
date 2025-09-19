package substation_test

import (
	"noize_metter/internal/service/substation"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFloat32(t *testing.T) {
	table := map[float32][]uint16{
		40:     {16928, 0},
		34:     {16904, 0},
		16:     {16768, 0},
		29:     {16872, 0},
		140748: {18441, 29440},
	}
	for expected, payload := range table {
		require.Equal(t, expected, substation.ParseFloat32(payload[0], payload[1]))
	}

	nTable := map[float32]int{
		40:     0,
		34:     2,
		16:     4,
		29:     6,
		140748: 8,
	}
	values := []uint16{16928, 0, 16904, 0, 16768, 0, 16872, 0, 18441, 29440}
	for expected, offset := range nTable {
		require.Equal(t, expected, substation.ParseFloat32V(values, offset))
	}
}
