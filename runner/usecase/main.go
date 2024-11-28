package usecase

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/runner/connector"
)

func New(config cli.Config, conn connector.Connector, c internal.Container) *UseCase {
	return &UseCase{
		Container: c,
		conn: conn,
		config: config,
	}
}

type UseCase struct {
	internal.Container
	conn connector.Connector
	config cli.Config
}
