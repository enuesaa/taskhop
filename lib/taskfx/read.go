package taskfx

import "gopkg.in/yaml.v3"

func (i *TaskSrv) Read() (Task, error) {
	filename := i.getTaskFilePath()
	var task Task

	f, err := i.repo.Read(filename)
	if err != nil {
		return task, err
	}

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&task); err != nil {
		return task, err
	}

	return task, nil
}
