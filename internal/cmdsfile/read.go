package cmdsfile

import (
	"os"

	"gopkg.in/yaml.v3"
)

func (i *Impl) Read(filename string) (CmdsFile, error) {
	var cmdsfile CmdsFile

	f, err := os.Open(filename)
	if err != nil {
		return cmdsfile, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cmdsfile); err != nil {
		return cmdsfile, err
	}

	return cmdsfile, nil
}
