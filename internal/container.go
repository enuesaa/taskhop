package internal

import (
	"github.com/enuesaa/taskhop/internal/cmdfx"
	"github.com/enuesaa/taskhop/internal/taskfx"
	"github.com/enuesaa/taskhop/internal/logfx"
	"github.com/enuesaa/taskhop/internal/runfx"
)

func NewContainer(
	cmdi cmdfx.I,
	taski taskfx.I,
	logi logfx.I,
	runi runfx.I,
) Container {
	return Container{
		Cmdi: cmdi,
		Taski: taski,
		Logi: logi,
		Runi: runi,
	}
}

type Container struct {
	Cmdi cmdfx.I
	Taski taskfx.I
	Logi logfx.I
	Runi runfx.I
}
