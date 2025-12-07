package main

import (
	"fmt"
	"os"

	"github.com/duanechan/salvare/internal/command"
)

func main() {
	app, err := command.LoadState()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		// app.displayHelp() or usage()
		fmt.Println("Error: not enough args")
		os.Exit(1)
	}

	if err := app.ParseRun(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
