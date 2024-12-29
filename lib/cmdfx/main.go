package cmdfx

import (
	"io"
	"os/exec"
)

func New() ICmdSrv {
	return &CmdSrv{}
}

type ICmdSrv interface {
	Exec(writer io.Writer, command string, workdir string) error
	Kill(cmd *exec.Cmd) error
}
type CmdSrv struct{}
