package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	println("Starting vars", fromPath, toPath, offset, limit)

	info, err := os.Stat(fromPath)
	if err != nil {
		panic(err)
	}

	if !info.Mode().IsRegular() {
		fmt.Println("Skipping non-regular file:", fromPath)
		panic(ErrUnsupportedFile)
	}

	size := info.Size()
	fmt.Println("File size:", size, "bytes")

	if offset > size {
		panic(ErrOffsetExceedsFileSize)
	}

	var copyLimit = size

	if limit > 0 {
		copyLimit = min(limit, size-offset)
	} else {
		copyLimit = size - offset
	}

	if limit > size {
		copyLimit = size
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return err
	}

	reader := io.LimitReader(file, copyLimit)

	outputFile, err := os.Create(toPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := outputFile
	// writer := io.Discard // copy to void

	bar := pb.Full.Start64(copyLimit)
	barReader := bar.NewProxyReader(reader)

	byteRead, err := io.Copy(writer, barReader)
	if err != nil {
		panic(err)
	}

	defer func() {
		bar.Finish()
	}()

	fmt.Println("*** I wanted to copy ", copyLimit, "bytes and copied", byteRead, "bytes")

	return nil
}

// CHECK BYTE DIFF BETWEEN
// go run . --from=/Users/mikhailmacherkevich/Downloads/mikhail-fnranrwpnytawhfevbbq_2026-04-15T09_16_31.615Z.webm --to=/tmp/ccc --offset=0
// go run . --from=/Users/mikhailmacherkevich/Downloads/mikhail-fnranrwpnytawhfevbbq_2026-04-15T09_16_31.615Z.webm --to=/tmp/ccc --offset=100000000
