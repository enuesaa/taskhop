package usecase

import "bytes"

func (c *UseCase) UnArchive(workdir string) error {
	var buf bytes.Buffer
	c.client.GetStorageArchive(&buf)

	return c.Arvi.UnArchive(&buf, workdir)
}
