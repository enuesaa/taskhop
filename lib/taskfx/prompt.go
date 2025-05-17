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

	for {
		fmt.Printf("> ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		if text == "exit" {
			i.errch <- ErrPromptExit
			return
		}
		i.current.Cmd = text
		i.current.Status = StatusProceeding

		<-i.completedch
	}
}
