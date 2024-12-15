package cmdfx

import "os/exec"

func (i *CmdSrv) Kill(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}
