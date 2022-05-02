package main

import (
	"io"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: add a Read([]byte) (int, error) method to MyReader.

// ????
func (r MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, io.EOF
	}

	return len(b), nil
}

func exerciseReader() {
	reader.Validate(MyReader{})
}
