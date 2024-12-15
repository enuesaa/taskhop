package repository

import (
	"io"
	"os"
	"path/filepath"
)

func New() IRepository {
	return &Repository{}
}

type IRepository interface {
	IsExist(path string) bool
	IsDir(path string) (bool, error)
	CreateDir(path string) error
	Create(path string, body []byte) error
	HomeDir() (string, error)
	WorkDir() (string, error)
	Remove(path string) error
	Read(path string) ([]byte, error)
	ListDirs(path string) ([]string, error)
	ListFiles(path string) ([]string, error)
}
type Repository struct{}

func (i *Repository) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (i *Repository) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (i *Repository) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (i *Repository) Create(path string, body []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		return err
	}
	return nil
}

func (i *Repository) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (i *Repository) WorkDir() (string, error) {
	return os.Getwd()
}

func (i *Repository) Remove(path string) error {
	return os.RemoveAll(path)
}

func (i *Repository) Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (i *Repository) ListDirs(path string) ([]string, error) {
	list := make([]string, 0)
	err := filepath.Walk(path, func(fpath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			list = append(list, fpath)
		}
		return nil
	})
	return list, err
}

func (i *Repository) ListFiles(path string) ([]string, error) {
	list := make([]string, 0)
	err := filepath.Walk(path, func(fpath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		list = append(list, fpath)
		return nil
	})
	return list, err
}
