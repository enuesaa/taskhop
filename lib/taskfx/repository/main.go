package repository

import (
	"bytes"
	"io"
	"os"
)

func New() IRepository {
	return &Repository{}
}

type IRepository interface {
	Read(path string) (io.Reader, error)
}
type Repository struct{}

func (i *Repository) Read(path string) (io.Reader, error) {
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
