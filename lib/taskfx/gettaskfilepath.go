package taskfx

import "path/filepath"

func (i *TaskSrv) getTaskFilePath() string {
	return filepath.Join(i.cli.GetWorkdir(), "cmds.yml")
}
