package taskfx

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrPromptExit = errors.New("exit")

func (i *TaskSrv) Prompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("> ")
	if !scanner.Scan() {
		return
	}

	text := scanner.Text()
	if text == "exit" {
		i.errch <- ErrPromptExit
		return
	}
	i.current.Cmd = text
	i.current.Status = StatusProceeding
}
