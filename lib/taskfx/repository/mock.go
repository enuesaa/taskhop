package repository

import (
	"io"
	"strings"
)

type Mock struct{
	Cmds string
}

func (i *Mock) Read(path string) (io.Reader, error) {
	return strings.NewReader(i.Cmds), nil
}
