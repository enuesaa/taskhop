package taskfx

type Completed bool

func (i *TaskSrv) MakeCompleted() error {
	i.current.Status = StatusPrompt

	go i.Prompt()

	return nil
}
