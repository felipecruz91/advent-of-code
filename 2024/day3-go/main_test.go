package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getRealMultiInstructions(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		input := strings.NewReader("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
		require.Equal(t, 161, getRealMultiInstructions(input))
	})
	t.Run("input", func(t *testing.T) {
		f, err := os.Open("data/input.txt")

		require.NoError(t, err)
		require.Equal(t, 164730528, getRealMultiInstructions(f))
	})
}
