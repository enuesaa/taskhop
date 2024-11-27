package client

import (
	"fmt"
	"net/http"
)

func New(commanderAddres string) GraphQLClient {
	url := fmt.Sprintf("http://%s/graphql", commanderAddres)

	return NewClient(http.DefaultClient, url, nil)
}
