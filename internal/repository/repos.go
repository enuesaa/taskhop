package repository

func newRepos(fs FsRepository, cmd CmdRepository) Repos {
	return Repos{
		Fs: &fs,
		Cmd: &cmd,
	}
}

type Repos struct {
	Fs  FsRepositoryInterface
	Cmd CmdRepositoryInterface
}
