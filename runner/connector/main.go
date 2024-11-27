package connector

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/client"
)

func New(address client.Address, cli client.GraphQLClient, c internal.Container) *Connector {
	return &Connector{
		Container: c,
		address: address,
		client: cli,
	}
}

type Connector struct {
	internal.Container
	address client.Address
	client client.GraphQLClient
}
