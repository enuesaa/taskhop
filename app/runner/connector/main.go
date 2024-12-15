package connector

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/cli"
)

func New(cl cli.ICli) Connector {
	url := fmt.Sprintf("http://%s/graphql", cl.GetAddress())
	client := NewClient(http.DefaultClient, url, nil)
	connector := Connector{
		GraphQLClient: client,
		cli:           cl,
	}

	return connector
}

type Connector struct {
	GraphQLClient
	cli cli.ICli
}
