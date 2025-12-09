package command

import (
	"errors"
	"fmt"

	"github.com/duanechan/salvare/internal/config"
	"github.com/duanechan/salvare/internal/models"
)

func CommandConfig(state *State, args []string) (*models.Metrics, error) {
	if state.Config.IsIncomplete() {
		return nil, errors.New("incomplete configuration file")
	}

	if err := config.WriteConfig(&config.Config{}); err != nil {
		return nil, err
	}

	fmt.Println("Created empty Salvare configuration file.")

	return nil, nil
}
