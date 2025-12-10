package command

import (
	"fmt"
	"time"

	"github.com/duanechan/salvare/internal/db"
	"github.com/duanechan/salvare/internal/dbbackup"
	"github.com/duanechan/salvare/internal/models"
)

func CommandBackup(state *State, args []string) (*models.Metrics, error) {
	start := time.Now()
	options := db.ParseBackupOptions(state.Config, args)
	data, err := state.driver.Backup(options...)

	if err != nil {
		return nil, err
	}

	if err := dbbackup.NewDumper(state.Config, false).WriteBackup(data); err != nil {
		return nil, err
	}
	fmt.Println("Backup created!")

	end := time.Since(start)
	return &models.Metrics{Took: end.Milliseconds()}, nil
}
