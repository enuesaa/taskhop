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
		assetsDownloaded: false,
	}
}

type ITaskSrv interface {
	Prompt() error
	Get() Task
	Register() (string, error)
	MakeCompleted() error
}

type TaskSrv struct {
	config *conf.Config
	repo   IRepository

	current Task
	assetsDownloaded bool
}
