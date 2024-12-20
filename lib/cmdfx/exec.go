package cmdfx

import (
	"io"
	"os/exec"
)

func (i *CmdSrv) Exec(writer io.Writer, command string) (*exec.Cmd, error) {
	cmd := exec.Command("bash", "-c", command)

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}
