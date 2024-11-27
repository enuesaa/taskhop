package connector

import (
	"fmt"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

func New(address string) (*Client, error) {
	url := fmt.Sprintf("http://%s/graphql", address)
	c := clientv2.NewClient(http.DefaultClient, url, nil)
	client := &Client{Client: c}

	if err := client.validate(); err != nil {
		return nil, err
	}

	return client, nil
}
