package app

import (
	"bufio"
	"fmt"
	"os"
)

func Prompt() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("> ")
	if !scanner.Scan() {
		return "", fmt.Errorf("end")
	}

	text := scanner.Text()
	if text == "exit" {
		return "", fmt.Errorf("end")
	}
	return text, nil
}
