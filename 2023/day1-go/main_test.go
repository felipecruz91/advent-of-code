package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	got := run(f)
	require.Equal(t, 56108, got)
}
