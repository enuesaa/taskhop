package connector

import (
	"fmt"
	"net/http"
)

func (c *Connector) UnArchive(commanderAddress, workdir string) error {
	url := fmt.Sprintf("http://%s/storage/archive", commanderAddress)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("failed to download archive file. status code: %s", res.Status)
	}

	return c.Arvi.UnArchive(res.Body, workdir)
}
