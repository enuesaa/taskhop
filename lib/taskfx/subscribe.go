package taskfx

func (i *TaskSrv) Subscribe() <-chan Status {
	return i.ch
}
