package gqlclient

import (
	"fmt"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

func New(commanderAddres string) *Client {
	url := fmt.Sprintf("http://%s/graphql", commanderAddres)
	c := clientv2.NewClient(http.DefaultClient, url, nil)

	return &Client{Client: c}
}
