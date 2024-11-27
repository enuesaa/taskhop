package gqlclient

import "net/url"

func (c *Client) Address() string {
	parsed, err := url.Parse(c.Client.BaseURL)
	if err != nil {
		return ""
	}
	return parsed.Host // like `localhost:3000`
}
