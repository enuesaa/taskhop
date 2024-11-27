package connector

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/client"
)

func New(cli client.GraphQLClient, c internal.Container) *Connector {
	return &Connector{
		Container: c,
		client: cli,
	}
}

type Connector struct {
	internal.Container
	client client.GraphQLClient
}
