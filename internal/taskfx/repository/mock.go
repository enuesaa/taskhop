package repository

import (
	"bytes"
	"io"
)

type Mock struct{}

func (i *Mock) Read(path string) (io.Reader, error) {
	var buf bytes.Buffer

	return &buf, nil
}
