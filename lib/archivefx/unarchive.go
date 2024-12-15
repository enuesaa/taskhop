package archivefx

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func (i *ArvSrv) UnArchive(r io.Reader, dest string) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	zipr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return err
	}

	for _, file := range zipr.File {
		fpath := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		f, err := os.Create(fpath)
		if err != nil {
			return err
		}
		defer f.Close()

		fr, err := file.Open()
		if err != nil {
			return err
		}
		defer fr.Close()

		if _, err = io.Copy(f, fr); err != nil {
			return err
		}
	}

	return nil
}
