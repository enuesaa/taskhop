package taskfx

type Status int

const (
	StatusRegistration Status = iota
	StatusPrompt
	StatusProceeding
	StatusCompleted
)

type Task struct {
	Status Status
	Cmds  []string
}
