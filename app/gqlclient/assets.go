package gqlclient

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Connector) DownloadAssets(dest io.Writer) error {
	url := fmt.Sprintf("http://%s/assets", c.config.Address)

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
