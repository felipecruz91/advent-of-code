package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistance(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		f, err := os.Open("data/test.txt")
		require.NoError(t, err)
		defer f.Close()

		got, err := distance(f)

		require.NoError(t, err)
		require.Equal(t, 11, got)
	})

	t.Run("input", func(t *testing.T) {
		f, err := os.Open("data/input.txt")
		require.NoError(t, err)
		defer f.Close()

		got, err := distance(f)

		require.NoError(t, err)
		require.Equal(t, 1580061, got)
	})
}

func TestScore(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		f, err := os.Open("data/test.txt")
		require.NoError(t, err)
		defer f.Close()

		got, err := score(f)

		require.NoError(t, err)
		require.Equal(t, 31, got)
	})

	t.Run("input", func(t *testing.T) {
		f, err := os.Open("data/input.txt")
		require.NoError(t, err)
		defer f.Close()

		got, err := score(f)

		require.NoError(t, err)
		require.Equal(t, 23046913, got)
	})
}
