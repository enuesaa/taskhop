package internal

import (
	"github.com/enuesaa/taskhop/internal/archivefx"
	"github.com/enuesaa/taskhop/internal/cmdfx"
	"github.com/enuesaa/taskhop/internal/logfx"
	"github.com/enuesaa/taskhop/internal/runfx"
	"github.com/enuesaa/taskhop/internal/taskfx"
)

func NewContainer(
	arvi archivefx.I,
	cmdi cmdfx.I,
	taski taskfx.I,
	logi logfx.I,
	runi runfx.I,
) Container {
	return Container{
		Arvi:  arvi,
		Cmdi:  cmdi,
		Taski: taski,
		Logi:  logi,
		Runi:  runi,
	}
}

type Container struct {
	Arvi  archivefx.I
	Cmdi  cmdfx.I
	Taski taskfx.I
	Logi  logfx.I
	Runi  runfx.I
}
