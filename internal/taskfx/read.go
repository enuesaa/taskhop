package taskfx

import (
	"os"

	"gopkg.in/yaml.v3"
)

func (i *Impl) Read() (Task, error) {
	filename := i.getTaskFilePath()
	var task Task

	f, err := os.Open(filename)
	if err != nil {
		return task, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&task); err != nil {
		return task, err
	}

	return task, nil
}
