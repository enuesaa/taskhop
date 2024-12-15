package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

type Lib struct {
	Arv  archivefx.I
	Cmd  cmdfx.I
	Log  logfx.I
	Task taskfx.I
}
