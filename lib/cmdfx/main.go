package cmdfx

import (
	"io"
	"os/exec"
)

func New() ICmdSrv {
	return &CmdSrv{}
}

type ICmdSrv interface {
	Exec(writer io.Writer, command string) (*exec.Cmd, error)
	MultiExec(writer io.Writer, commands []string, workdir string) error
	Kill(cmd *exec.Cmd) error
}
type CmdSrv struct {}
