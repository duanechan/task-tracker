package main

import (
	"fmt"
	"os"

	task "github.com/duanechan/task-tracker-cli/internal"
)

func main() {
	state, err := task.LoadState()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		state.DisplayCommands()
		os.Exit(1)
	}

	if err = state.Run(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
