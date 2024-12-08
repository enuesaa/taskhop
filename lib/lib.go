package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func New() ILib {
	return Lib{}
}

type ILib interface{}

type Lib struct{
	Arv archivefx.I
	Cmd cmdfx.I
	Task taskfx.I
	Log logfx.I
}
