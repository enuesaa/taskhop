package usecase

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/connector"
)

func New(client connector.Client, c internal.Container) *UseCase {
	return &UseCase{
		Container: c,
		client: client,
		Workdir: ".",
	}
}

type UseCase struct {
	internal.Container
	client connector.Client
	Workdir string
}
