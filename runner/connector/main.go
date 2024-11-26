package connector

import "github.com/enuesaa/taskhop/runner/client"

func New(address client.Address, cli client.GraphQLClient) *Connector {
	return &Connector{
		address: address,
		client: cli,
	}
}

type Connector struct {
	address client.Address
	client client.GraphQLClient
}
