package cmdexec

import (
	"io"
	"os/exec"
)

func New() I {
	return &Impl{}
}

type I interface {
	Exec(writer io.Writer, command string) (*exec.Cmd, error)
	MultiExec(writer io.Writer, commands []string) error
	Kill(cmd *exec.Cmd) error
}
type Impl struct{}

// TODO:
// おそらく cmd store みたいなのが必要
