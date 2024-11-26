package archivefx

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (i *Impl) Archive() (io.Reader, error) {
	zipb := bytes.NewBuffer([]byte{})
	zipw := zip.NewWriter(zipb)
	defer zipw.Close()

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		fpath, err := filepath.Rel(".", path)
		if err != nil {
			return err
		}
		w, err := zipw.Create(fpath)
		if err != nil {
			return err
		}

		_, err = io.Copy(w, f)
		return err
	})

	if err != nil {
		return nil, err
	}
	return zipb, nil
}
