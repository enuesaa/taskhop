package cmdfx

import "io"

func New() ICmdSrv {
	return &CmdSrv{}
}

type ICmdSrv interface {
	Exec(writer io.Writer, command string, workdir string) error
}
type CmdSrv struct{}
