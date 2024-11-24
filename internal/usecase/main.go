package usecase

import "github.com/enuesaa/taskhop/internal/fs"

func New(fsrepo fs.FsRepositoryInterface) Usecase {
	return Usecase{
		fs: fsrepo,
	}
}

type Usecase struct {
	fs fs.FsRepositoryInterface
}
