package taskfx

import "fmt"

func (i *TaskSrv) Get() (Task, error) {
	if len(i.current.Cmds) == 0 {
		return Task{}, fmt.Errorf("waiting")
	}
	return i.current, nil
}
