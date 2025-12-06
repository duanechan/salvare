package db

import "errors"

type Driver interface {
	Backup() error
	Restore() error
	Compress() error
}

func ReadDriver(scheme string) (Driver, error) {
	switch scheme {
	case "postgres", "postgresql":
		return nil, nil
	default:
		return nil, errors.New("unsupported scheme")
	}
}
