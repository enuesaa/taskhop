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
		endCh: make(chan bool, 1),
	}
}

type ITaskSrv interface {
	Prompt() error
	Get() Task
	Register() (string, error)
	MakeCompleted() error
	SubscribeEnd() <-chan bool
}

type TaskSrv struct {
	config *conf.Config
	repo   IRepository

	current Task
	assetsDownloaded bool
	endCh chan bool
}
