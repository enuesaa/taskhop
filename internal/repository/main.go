package repository

type Repos struct {
	Fs  FsRepositoryInterface
	Cmd CmdRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs: &FsRepository{},
		Cmd: &CmdRepository{},
	}
}
