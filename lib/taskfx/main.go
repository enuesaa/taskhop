package taskfx

import (
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
)

func New(config *conf.Config, repo repository.IRepository) ITaskSrv {
	return &TaskSrv{
		config: config,
		repo: repo,
	}
}

type ITaskSrv interface {
	Read() (Task, error)
	Validate(task Task) error
}

type TaskSrv struct {
	config *conf.Config
	repo repository.IRepository
}
