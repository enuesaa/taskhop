package runfx

type Status int

const (
	StatusWaiting Status = iota
	StatusProceeding
	StatusCompleted
)
