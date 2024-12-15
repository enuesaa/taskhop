package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
	"github.com/enuesaa/taskhop/lib/procfx"
)

type Lib struct {
	Arv  archivefx.IArvSrv
	Cmd  cmdfx.ICmdSrv
	Log  logfx.ILogSrv
	Task taskfx.ITaskSrv
	Proc procfx.IProcSrv
}
