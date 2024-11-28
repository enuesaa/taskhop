package archivefx

import (
	"io"

	"github.com/enuesaa/taskhop/internal/archivefx/repository"
	"github.com/enuesaa/taskhop/internal/cli"
)

func New(config cli.Config, repo repository.I) I {
	return &Impl{
		config: config,
		repo: repo,
		Workdir: ".",
	}
}

type I interface {
	Archive() (io.Reader, error)
	UnArchive(r io.Reader, dest string) error
}
type Impl struct {
	config cli.Config
	repo repository.I
	Workdir string
}
