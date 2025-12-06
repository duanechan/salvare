package main

import (
	"fmt"
	"os"

	"github.com/duanechan/salvare/internal/salvare"
)

func main() {
	app, err := salvare.LoadState()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if err := app.ParseRun(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
