package taskfx

import "github.com/enuesaa/taskhop/conf"

func New(config *conf.Config, repo IRepository) ITaskSrv {
	return &TaskSrv{
		config: config,
		repo:   repo,
		current: Task{
			Status: StatusRegistration,
			Cmds: []string{},
		},
	}
}

type ITaskSrv interface {
	Prompt() error
	Get() Task
	Register() (string, error)
	NotifyCompleted() error
}

type TaskSrv struct {
	config *conf.Config
	repo   IRepository

	current Task
}
