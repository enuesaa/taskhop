package adapter

import "io"

func (c *Adapter) DownloadAssets(dest io.Writer) error {
	res, err := c.get("/assets")
	if err != nil {
		return err
	}
	defer res.Body.Close() //nolint:errcheck

	if _, err := io.Copy(dest, res.Body); err != nil {
		return err
	}
	return nil
}

func (c *Adapter) UploadAssets(body io.Reader) error {
	res, err := c.post("/upload", "application/zip", body)
	if err != nil {
		return err
	}
	defer res.Body.Close() //nolint:errcheck

	return nil
}
