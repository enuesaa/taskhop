package connector

import "errors"

var ErrInvalidAddress = errors.New("invalid address")

func (c *Client) validate() error {
	if c.Address() == "" {
		return ErrInvalidAddress
	}
	return nil
}
