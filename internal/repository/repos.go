package repository

type Repos struct {
	Fs  FsRepositoryInterface
	Cmd CmdRepositoryInterface
}
