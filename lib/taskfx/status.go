package taskfx

type Status int

const (
	StatusWaiting Status = iota
	StatusProceeding
	StatusCompleted
)
