//go:build wireinject
// +build wireinject

package repository

import "github.com/google/wire"

func New() Repos {
	wire.Build(newFsRepository, newCmdRepository, newRepos)

	return Repos{}
}
