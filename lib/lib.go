package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

type Lib struct {
	Arv  archivefx.IArvSrv
	Cmd  cmdfx.ICmdSrv
	Log  logfx.ILogSrv
	Task taskfx.ITaskSrv
}
