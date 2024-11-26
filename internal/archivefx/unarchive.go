package archivefx

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func (i *Impl) UnArchive() error {
	dest := "./aaa"

	zipr, err := zip.OpenReader("archive.zip")
	if err != nil {
		return err
	}
	defer zipr.Close()

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
