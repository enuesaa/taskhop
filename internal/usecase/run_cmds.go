package usecase

import "bytes"

func (u *Usecase) RunCmds() (string, error) {
	var out bytes.Buffer
	commands := []string{
		"echo a",
		"echo b",
		"cd /var",
		"pwd",
	}
	if err := u.cmdexec.MultiExec(&out, commands); err != nil {
		return out.String(), err
	}
	return out.String(), nil
}
