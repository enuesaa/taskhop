package runner

import (
	"bytes"
	"context"
)

func (a *App) UnArchive() error {
	a.lib.Log.Info(context.Background(), "donwload archive files...")

	var buf bytes.Buffer
	a.conn.GetStorageArchive(&buf)

	return a.lib.Arv.UnArchive(&buf, a.cli.GetWorkdir())
}
