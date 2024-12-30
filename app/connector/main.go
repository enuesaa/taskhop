package connector

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/conf"
)

func New(config *conf.Config) Connector {
	url := fmt.Sprintf("http://%s/graphql", config.Address)
	client := NewClient(http.DefaultClient, url, nil)
	connector := Connector{
		GraphQLClient: client,
		config:        config,
	}

	return connector
}

type Connector struct {
	GraphQLClient
	config *conf.Config
}
