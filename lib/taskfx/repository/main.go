package repository

import (
	"bytes"
	"io"
	"os"
)

func New() I {
	return &Impl{}
}

type I interface {
	Read(path string) (io.Reader, error)
}
type Impl struct{}

func (i *Impl) Read(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		return nil, err
	}

	return &buf, nil
}
