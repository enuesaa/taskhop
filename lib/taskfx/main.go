package taskfx

import (
	"time"

	"github.com/enuesaa/taskhop/conf"
)

func New(config *conf.Config, repo IRepository) ITaskSrv {
	return &TaskSrv{
		config: config,
		repo:   repo,
		current: Task{
			Status: StatusRegistration,
			Cmd: "",
		},
		assetsDownloaded: false,
		errch: make(chan error, 1),
		completedch: make(chan bool, 1),
		lastHealthy: time.Now(),
	}
}

type ITaskSrv interface {
	Prompt()
	Get() Task
	Register() (string, error)
	MakeCompleted() error
	Subscribe() <-chan error
}

type TaskSrv struct {
	config *conf.Config
	repo   IRepository

	current Task
	assetsDownloaded bool
	errch chan error
	completedch chan bool
	lastHealthy time.Time
}
