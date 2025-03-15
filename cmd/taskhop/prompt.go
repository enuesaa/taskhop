package main

import (
	"bufio"
	"fmt"
	"os"
)

func Prompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		if text == "exit" {
			break
		}
	}
	os.Exit(0)
}
