package taskfx

import "path/filepath"

func (i *Impl) getTaskFilePath() string {
	return filepath.Join(i.config.Workdir, "cmds.yml")
}
