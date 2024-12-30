package runner

import "bytes"

func (a *Runner) UnArchive() error {
	var buf bytes.Buffer
	a.conn.GetStorageArchive(&buf)

	return a.lib.Arv.UnArchive(&buf, a.config.Workdir)
}
