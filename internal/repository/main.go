package repository

import (
	"github.com/enuesaa/taskhop/internal/repository/cmd"
	"github.com/enuesaa/taskhop/internal/repository/fs"
	"github.com/enuesaa/taskhop/internal/repository/logging"
)

type Repos struct {
	Fs  fs.FsRepositoryInterface
	Cmd cmd.CmdRepositoryInterface
	Log logging.LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:  &fs.FsRepository{},
		Cmd: &cmd.CmdRepository{},
		Log: &logging.LogRepository{},
	}
}
