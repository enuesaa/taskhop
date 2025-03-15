package taskfx

type Status int

const (
	StatusWaiting Status = iota
	StatusProceeding
	StatusCompleted
)

func (i *TaskSrv) GetStatus() Status {
	return i.status
}
