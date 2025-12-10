package command

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/duanechan/salvare/internal/config"
	"github.com/duanechan/salvare/internal/db"
)

type State struct {
	cmdRegistry registry
	driver      db.Driver
	Config      *config.Config
}

func LoadState() (*State, error) {
	cfg, err := config.LoadConfig()
	if err != nil && (!errors.Is(err, config.ConfigFileNotExists) &&
		!errors.Is(err, config.ConfigFileEOF)) {
		return nil, err
	}

	var driver db.Driver
	if cfg != nil && cfg.ConnectionString() != config.EmptyConnString {
		dbURL, err := url.Parse(cfg.ConnectionString())
		if err != nil {
			return nil, err
		}

		driver, err = db.GetDriver(dbURL)
		if err != nil {
			return nil, err
		}
	}

	return &State{
		cmdRegistry: loadCmdRegistry(),
		driver:      driver,
		Config:      cfg,
	}, nil
}

func (s *State) ParseRun(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	name, rest := args[0], args[1:]

	subCommands, exists := s.cmdRegistry[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	sub := ""
	if len(rest) >= 1 {
		sub = rest[0]
	}
	command, exists := subCommands[sub]
	if !exists {
		// command.displayUsage()
		return fmt.Errorf("command '%s %s' does not exist", name, sub)
	}

	metrics, err := command.callback(s, rest)
	if err != nil {
		return err
	}

	if metrics != nil {
		fmt.Printf("Operation took %dms\n", metrics.Took)
	}

	return nil
}
