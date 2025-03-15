package taskfx

func (i *TaskSrv) Get() Task {
	return i.current
}
