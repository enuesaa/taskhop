package archivefx

import (
	"io"

	"github.com/enuesaa/taskhop/internal/archivefx/repository"
)

func New(repo repository.I) I {
	return &Impl{
		repo: repo,
	}
}

type I interface {
	Archive() (io.Reader, error)
	UnArchive() error
}
type Impl struct {
	repo repository.I
}
