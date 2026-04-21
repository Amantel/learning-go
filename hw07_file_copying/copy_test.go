package main

import (
	"os"
	"testing"
)

func readFile(t *testing.T, path string) []byte {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	return data
}

func runCopy(t *testing.T, offset, limit int64) {
	t.Helper()

	src := "testdata/input.txt"
	dst := "out.txt"

	err := Copy(src, dst, offset, limit)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}
}

func checkEqual(t *testing.T, gotPath, wantPath string) {
	t.Helper()

	got := readFile(t, gotPath)
	want := readFile(t, wantPath)

	if string(got) != string(want) {
		t.Fatalf("files differ:\nGOT:\n%s\nWANT:\n%s", got, want)
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name   string
		offset int64
		limit  int64
		want   string
	}{
		{"no offset no limit", 0, 0, "testdata/out_offset0_limit0.txt"},
		{"limit 10", 0, 10, "testdata/out_offset0_limit10.txt"},
		{"limit 1000", 0, 1000, "testdata/out_offset0_limit1000.txt"},
		{"limit 10000", 0, 10000, "testdata/out_offset0_limit10000.txt"},
		{"offset 100 limit 1000", 100, 1000, "testdata/out_offset100_limit1000.txt"},
		{"offset 6000 limit 1000", 6000, 1000, "testdata/out_offset6000_limit1000.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runCopy(t, tt.offset, tt.limit)
			checkEqual(t, "out.txt", tt.want)

			os.Remove("out.txt")
		})
	}
}
