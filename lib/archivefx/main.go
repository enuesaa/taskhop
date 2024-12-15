package archivefx

import (
	"io"

	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/archivefx/repository"
)

func New(cl cli.ICli, repo repository.I) I {
	return &Impl{
		cli:  cl,
		repo: repo,
	}
}

type I interface {
	Archive() (io.Reader, error)
	UnArchive(r io.Reader, dest string) error
}
type Impl struct {
	cli  cli.ICli
	repo repository.I
}
