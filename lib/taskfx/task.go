package taskfx

import "time"

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
	Cmd    string
}

func (i *TaskSrv) Get() Task {
	i.lastHealthy = time.Now()

	return i.current
}

func (i *TaskSrv) MakeCompleted() error {
	i.current.Status = StatusPrompt
	go i.Prompt()

	return nil
}

func (i *TaskSrv) Subscribe() <-chan error {
	return i.errch
}
