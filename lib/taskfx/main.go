package taskfx

import "github.com/enuesaa/taskhop/conf"

func New(config *conf.Config, repo IRepository) ITaskSrv {
	return &TaskSrv{
		config: config,
		repo:   repo,
	}
}

type ITaskSrv interface {
	Read() (Task, error)
	Validate(task Task) error
}

type TaskSrv struct {
	config *conf.Config
	repo   IRepository
}
