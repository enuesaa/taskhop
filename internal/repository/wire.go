//go:build wireinject
// +build wireinject

package repository

import "github.com/google/wire"

func newCmdRepository() CmdRepository {
	return CmdRepository{}
}
func newFsRepository() FsRepository {
	return FsRepository{}
}
func newRepos(fs FsRepository, cmd CmdRepository) Repos {
	return Repos{
		Fs: &fs,
		Cmd: &cmd,
	}
}

func New() Repos {
	wire.Build(
		newCmdRepository,
		newFsRepository,
		newRepos,
	)

	return Repos{}
}
