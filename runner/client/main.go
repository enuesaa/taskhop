package client

import (
	"fmt"
	"net/http"
)

func New(address Address) GraphQLClient {
	url := fmt.Sprintf("http://%s/graphql", address)

	return NewClient(http.DefaultClient, url, nil)
}
