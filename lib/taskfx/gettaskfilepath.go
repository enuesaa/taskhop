package taskfx

import "path/filepath"

func (i *Impl) getTaskFilePath() string {
	return filepath.Join(i.cli.GetWorkdir(), "cmds.yml")
}
