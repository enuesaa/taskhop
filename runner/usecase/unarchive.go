package usecase

import "bytes"

func (u *UseCase) UnArchive() error {
	var buf bytes.Buffer
	u.client.GetStorageArchive(&buf)

	return u.Arvi.UnArchive(&buf, u.config.Workdir)
}
