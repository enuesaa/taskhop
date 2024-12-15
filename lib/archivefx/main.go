package archivefx

import (
	"io"

	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/archivefx/repository"
)

func New(cl cli.ICli, repo repository.IRepository) IArvSrv {
	return &ArvSrv{
		cli:  cl,
		repo: repo,
	}
}

type IArvSrv interface {
	Archive() (io.Reader, error)
	UnArchive(r io.Reader, dest string) error
}
type ArvSrv struct {
	cli  cli.ICli
	repo repository.IRepository
}
