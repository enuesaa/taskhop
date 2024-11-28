package usecase

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/runner/connector"
)

func New(config cli.Config, client connector.Client, c internal.Container) *UseCase {
	return &UseCase{
		Container: c,
		client: client,
		config: config,
	}
}

type UseCase struct {
	internal.Container
	client connector.Client
	config cli.Config
}
