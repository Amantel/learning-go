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

func Copy(fromPath, toPath string, offset, limit int64) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	size := info.Size()

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	remaining := size - offset

	var copyLimit int64
	if limit > 0 {
		copyLimit = min(limit, remaining)
	} else {
		copyLimit = remaining
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	reader := io.LimitReader(file, copyLimit)

	outputFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	bar := pb.Full.Start64(copyLimit)
	defer bar.Finish()

	barReader := bar.NewProxyReader(reader)

	written, err := io.Copy(outputFile, barReader)
	if err != nil {
		return err
	}

	fmt.Println("copied:", written, "bytes")

	return nil
}
