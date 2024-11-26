package archivefx

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Archive() error {
	zipb := bytes.NewBuffer([]byte{})
	zipw := zip.NewWriter(zipb)

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
		return err
	}
	if err := zipw.Close(); err != nil {
		return err
	}

	zipf, err := os.Create("archive.zip")
	if err != nil {
		return err
	}
	defer zipf.Close()
	
	if _, err := io.Copy(zipf, zipb); err != nil {
		return err
	}
	return nil
}
