package repository

type Repos struct {
	Fs  FsRepositoryInterface
	Cmd CmdRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:  &FsRepository{},
		Cmd: &CmdRepository{},
		Log: &LogRepository{},
	}
}
