package taskfx

import "time"

type Status int

const (
	StatusRegistration Status = iota
	StatusPrompt
	StatusProceeding
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
	i.completedch <- true

	return nil
}

func (i *TaskSrv) Subscribe() <-chan error {
	return i.errch
}
