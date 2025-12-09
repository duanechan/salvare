package command

import (
	"fmt"
	"time"

	"github.com/duanechan/salvare/internal/models"
)

func CommandBackup(state *State, args []string) (*models.Metrics, error) {
	start := time.Now()
	data, err := state.driver.Backup()
	end := time.Since(start)

	if err != nil {
		return nil, err
	}

	fmt.Println("Backup created!")
	if err := state.dumpFile(data); err != nil {
		return nil, err
	}

	return &models.Metrics{Took: end.Milliseconds()}, nil
}
