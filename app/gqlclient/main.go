package gqlclient

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/conf"
)

func New(config *conf.Config) Connector {
	url := fmt.Sprintf("http://%s/graphql", config.Address)
	client := NewClient(http.DefaultClient, url, nil)

	connector := Connector{
		Gql: client,
		config: config,
	}
	return connector
}

type Connector struct {
	Gql GraphQLClient
	config *conf.Config
}
