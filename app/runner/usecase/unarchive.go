package usecase

import (
	"bytes"
	"context"
)

func (u *UseCase) UnArchive() error {
	u.Logi.Info(context.Background(), "donwload archive files...")

	var buf bytes.Buffer
	u.conn.GetStorageArchive(&buf)

	return u.Arvi.UnArchive(&buf, u.config.Workdir)
}
