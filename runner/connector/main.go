package connector

import (
	"fmt"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

func New(address string) (*Client, error) {
	url := fmt.Sprintf("http://%s/graphql", address)
	c := clientv2.NewClient(http.DefaultClient, url, nil)

	return &Client{Client: c}, nil
}
