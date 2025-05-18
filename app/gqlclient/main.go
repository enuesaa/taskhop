package gqlclient

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/app/gqlclient/adapter"
)

func New(config *conf.Config) Connector {
	url := fmt.Sprintf("http://%s/graphql", config.Address)
	client := adapter.NewClient(http.DefaultClient, url, nil)

	connector := Connector{
		gql:    client,
		config: config,
		LogWriter: LogWriter{
			gql: client,
		},
	}
	return connector
}

type Connector struct {
	gql    adapter.GQLClient
	config *conf.Config

	LogWriter LogWriter
}
