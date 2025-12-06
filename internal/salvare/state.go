package salvare

import (
	"errors"
	"fmt"

	"github.com/duanechan/salvare/internal/config"
)

type State struct {
	cmdRegistry commands
	config      *config.Config
}

func LoadState() (*State, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	return &State{
		cmdRegistry: loadCmdRegistry(),
		config:      config,
	}, nil
}

func (s State) ParseRun(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	name, rest := args[1], args[1:]

	cmd, exists := s.cmdRegistry[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	return cmd.callback(name, rest)
}
