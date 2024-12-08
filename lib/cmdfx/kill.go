package cmdfx

import (
	"os/exec"
)

func (i *Impl) Kill(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}
