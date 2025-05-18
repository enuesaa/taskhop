package adapter

import (
	"fmt"
	"io"
	"net/http"
)

func New(address string) *Adapter {
	url := fmt.Sprintf("http://%s/graphql", address)
	gql := NewClient(http.DefaultClient, url, nil)

	return &Adapter{
		address: address,
		gql:     gql,
	}
}

type Adapter struct {
	address string
	gql     GQLClient
}

func (c *Adapter) get(endpoint string) (*http.Response, error) {
	url := fmt.Sprintf("http://%s%s", c.address, endpoint)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("err status code: %s", res.Status)
	}
	return res, nil
}

func (c *Adapter) post(endpoint string, contentType string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("http://%s%s", c.address, endpoint)

	res, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("err status code: %s", res.Status)
	}
	return res, nil
}
