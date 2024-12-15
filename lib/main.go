package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func New(
	arv archivefx.I,
	cmd cmdfx.I,
	task taskfx.I,
	log logfx.I,
) Lib {
	return Lib{
		Arv:  arv,
		Cmd:  cmd,
		Task: task,
		Log:  log,
	}
}
