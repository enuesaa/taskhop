package connector

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetStorageArchive(dest io.Writer) error {
	address := c.Address()
	url := fmt.Sprintf("http://%s/storage/archive", address)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("failed to download archive file. status code: %s", res.Status)
	}
	if _, err := io.Copy(dest, res.Body); err != nil {
		return err
	}

	return nil
}
