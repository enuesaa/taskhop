package usecase

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/gqlclient"
)

func New(cli gqlclient.Client, c internal.Container) *UseCase {
	return &UseCase{
		Container: c,
		client: cli,
	}
}

type UseCase struct {
	internal.Container
	client gqlclient.Client
}
