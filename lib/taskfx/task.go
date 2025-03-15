package taskfx

type Status int

const (
	StatusWaiting Status = iota
	StatusProceeding
	StatusCompleted
)

type Task struct {
	Status Status
	Cmds  []string
}
