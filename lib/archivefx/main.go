package archivefx

import (
	"io"

	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib/archivefx/repository"
)

func New(config *conf.Config, repo repository.IRepository) IArvSrv {
	return &ArvSrv{
		config: config,
		repo:   repo,
	}
}

type IArvSrv interface {
	Archive() (io.Reader, error)
	UnArchive(r io.Reader, dest string) error
}
type ArvSrv struct {
	config *conf.Config
	repo   repository.IRepository
}
