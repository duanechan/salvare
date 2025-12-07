package command

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/duanechan/salvare/internal/config"
	"github.com/duanechan/salvare/internal/db"
)

type State struct {
	cmdRegistry commands
	driver      db.Driver
	Config      *config.Config
}

func LoadState() (*State, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	var driver db.Driver
	if cfg.ConnectionString() != config.EmptyConnString {
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

	cmd, exists := s.cmdRegistry[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	return cmd.callback(s, rest)
}

const (
	defaultLayout string = "Backup_2006_01_02_15_04_05_%s.sql"
)

func (s *State) dumpFile(data []byte) error {
	filename := time.Now().Format(fmt.Sprintf(defaultLayout, s.Config.Conn.Database))
	path := filepath.Join(s.Config.BackupDirectory, filename)

	if err := os.MkdirAll(s.Config.BackupDirectory, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
