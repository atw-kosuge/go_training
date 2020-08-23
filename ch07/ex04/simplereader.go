package main

import (
	"io"
)

// SimpleReader reader
type SimpleReader struct {
	text  string
	count int64
}

// NewReader create reader
func NewReader(s string) *SimpleReader {
	return &SimpleReader{s, 0}
}

func (reader *SimpleReader) Read(buffer []byte) (count int, err error) {
	if reader.count >= int64(len(reader.text)) {
		return 0, io.EOF
	}
	count = copy(buffer, reader.text[reader.count:])
	reader.count += int64(count)
	return
}
