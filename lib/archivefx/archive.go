package archivefx

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func (i *ArvSrv) Archive() (io.Reader, error) {
	zipb := bytes.NewBuffer([]byte{})
	zipw := zip.NewWriter(zipb)
	defer zipw.Close()

	err := filepath.Walk(i.cli.GetWorkdir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		fpath, err := filepath.Rel(i.cli.GetWorkdir(), path)
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
