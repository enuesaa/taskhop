package usecase

import "bytes"

func (u *UseCase) UnArchive() error {
	var buf bytes.Buffer
	u.conn.GetStorageArchive(&buf)

	return u.Arvi.UnArchive(&buf, u.config.Workdir)
}
