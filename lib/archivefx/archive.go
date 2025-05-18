package archivefx

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func (i *ArvSrv) Archive(dir string) (io.Reader, error) {
	zipb := bytes.NewBuffer([]byte{})
	zipw := zip.NewWriter(zipb)
	defer zipw.Close()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fbytes, err := i.repo.Read(path)
		if err != nil {
			return err
		}

		fpath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		w, err := zipw.Create(fpath)
		if err != nil {
			return err
		}

		_, err = io.Copy(w, bytes.NewReader(fbytes))
		return err
	})

	if err != nil {
		return nil, err
	}
	return zipb, nil
}
