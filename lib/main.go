package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
	"github.com/enuesaa/taskhop/lib/procfx"
)

func New(
	arv archivefx.IArvSrv,
	cmd cmdfx.ICmdSrv,
	task taskfx.ITaskSrv,
	proc procfx.IProcSrv,
	log logfx.ILogSrv,
) Lib {
	return Lib{
		Arv:  arv,
		Cmd:  cmd,
		Task: task,
		Proc: proc,
		Log:  log,
	}
}
