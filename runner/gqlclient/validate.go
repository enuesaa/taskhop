package gqlclient

import "errors"

var ErrInvalidAddress = errors.New("invalid address")

func (c *Client) Validate() error {
	if c.Address() == "" {
		return ErrInvalidAddress
	}
	return nil
}
