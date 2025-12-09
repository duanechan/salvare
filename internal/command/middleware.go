package command

import (
	"errors"

	"github.com/duanechan/salvare/internal/models"
)

func DriverMiddleware(handler handler) handler {
	return func(state *State, args []string) (*models.Metrics, error) {
		if state.driver == nil {
			return nil, errors.New("driver not set")
		}
		return handler(state, args)
	}
}
