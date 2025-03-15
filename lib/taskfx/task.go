package taskfx

type Status int

const (
	StatusRegistration Status = iota
	StatusDownloadAssets
	StatusPrompt
	StatusProceeding
	StatusCompleted
	StatusEnd // TODO
)

type Task struct {
	Status Status
	Cmds  []string
}
