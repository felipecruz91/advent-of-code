package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSafe(t *testing.T) {
	t.Run("gradually increasing", func(t *testing.T) {
		require.True(t, isSafe([]int{1, 2, 3, 4, 5}))
	})
	t.Run("gradually decreasing", func(t *testing.T) {
		require.True(t, isSafe([]int{5, 4, 3, 2, 1}))
	})
	t.Run("not gradually increasing or decreasing", func(t *testing.T) {
		require.False(t, isSafe([]int{1, 3, 2, 4, 5}))
	})
}

func Test_safeReports(t *testing.T) {
	t.Run("safe because the levels are all decreasing by 1 or 2.", func(t *testing.T) {
		require.Equal(t, 1, safeReports(strings.NewReader("7 6 4 2 1")))
	})
	t.Run("Unsafe because 2 7 is an increase of 5.", func(t *testing.T) {
		require.Equal(t, 0, safeReports(strings.NewReader("1 2 7 8 9")))
	})
	t.Run("Unsafe because 6 2 is a decrease of 4.", func(t *testing.T) {
		require.Equal(t, 0, safeReports(strings.NewReader("9 7 6 2 1")))
	})
	t.Run("Unsafe because 1 3 is increasing but 3 2 is decreasing.", func(t *testing.T) {
		require.Equal(t, 0, safeReports(strings.NewReader("1 3 2 4 5")))
	})
	t.Run("Unsafe because 4 4 is neither an increase or a decrease.", func(t *testing.T) {
		require.Equal(t, 0, safeReports(strings.NewReader("8 6 4 4 1")))
	})
	t.Run("Safe because the levels are all increasing by 1, 2, or 3.", func(t *testing.T) {
		require.Equal(t, 1, safeReports(strings.NewReader("1 3 6 7 9")))
	})
	t.Run("Safe reports from input", func(t *testing.T) {
		f, err := os.Open("data/input.txt")
		require.NoError(t, err)
		defer f.Close()

		require.Equal(t, 257, safeReports(f))
	})
}

func Test_dampenerSafeReports(t *testing.T) {
	t.Run("safe because the levels are all decreasing by 1 or 2.", func(t *testing.T) {
		require.Equal(t, 1, dampenerSafeReports(strings.NewReader("7 6 4 2 1")))
	})
	t.Run("Unsafe regardless of which level is removed.", func(t *testing.T) {
		require.Equal(t, 0, dampenerSafeReports(strings.NewReader("1 2 7 8 9")))
	})
	t.Run("Unsafe regardless of which level is removed.", func(t *testing.T) {
		require.Equal(t, 0, dampenerSafeReports(strings.NewReader("9 7 6 2 1")))
	})
	t.Run("Safe by removing the second level, 3.", func(t *testing.T) {
		require.Equal(t, 1, dampenerSafeReports(strings.NewReader("1 3 2 4 5")))
	})
	t.Run("Safe by removing the third level, 4.", func(t *testing.T) {
		require.Equal(t, 1, dampenerSafeReports(strings.NewReader("8 6 4 4 1")))
	})
	t.Run("Safe without removing any level.", func(t *testing.T) {
		require.Equal(t, 1, dampenerSafeReports(strings.NewReader("1 3 6 7 9")))
	})
	t.Run("Safe reports from input", func(t *testing.T) {
		f, err := os.Open("data/input.txt")
		require.NoError(t, err)
		defer f.Close()

		require.Equal(t, 328, dampenerSafeReports(f))
	})
}
