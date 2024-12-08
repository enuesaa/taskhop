package lib

import (
	"github.com/enuesaa/taskhop/internal/archivefx"
	"github.com/enuesaa/taskhop/internal/cmdfx"
	"github.com/enuesaa/taskhop/internal/logfx"
	"github.com/enuesaa/taskhop/internal/taskfx"
)

func NewContainer(
	arvi archivefx.I,
	cmdi cmdfx.I,
	taski taskfx.I,
	logi logfx.I,
) Container {
	return Container{
		Arvi:  arvi,
		Cmdi:  cmdi,
		Taski: taski,
		Logi:  logi,
	}
}

type Container struct {
	Arvi  archivefx.I
	Cmdi  cmdfx.I
	Taski taskfx.I
	Logi  logfx.I
}
