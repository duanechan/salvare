package command

import "errors"

func DriverMiddleware(handler handler) handler {
	return func(state *State, args []string) error {
		if state.driver == nil {
			return errors.New("driver not set")
		}
		return handler(state, args)
	}
}
