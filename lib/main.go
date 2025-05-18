package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func New(
	arv archivefx.IArvSrv,
	cmd cmdfx.ICmdSrv,
	task taskfx.ITaskSrv,
	log logfx.ILogSrv,
) *Lib {
	return &Lib{
		Arv:  arv,
		Cmd:  cmd,
		Task: task,
		Log:  log,
	}
}
