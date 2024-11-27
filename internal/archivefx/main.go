package archivefx

import (
	"io"

	"github.com/enuesaa/taskhop/internal/archivefx/repository"
)

func New(repo repository.I) I {
	return &Impl{
		repo: repo,
		Workdir: ".",
	}
}

type I interface {
	Use(workdir string)
	Archive() (io.Reader, error)
	UnArchive(r io.Reader, dest string) error
}
type Impl struct {
	repo repository.I
	Workdir string
}
