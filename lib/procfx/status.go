package procfx

type Status int

const (
	StatusWaiting Status = iota
	StatusProceeding
	StatusCompleted
)

func (i *ProcSrv) GetStatus() Status {
	return i.status
}

func (i *ProcSrv) Subscribe() <-chan Status {
	return i.ch
}
