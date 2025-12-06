package db

import (
	"errors"
	"fmt"
	"net/url"
)

type Driver interface {
	Backup() error
	Restore() error
	Compress() error
}

type baseDriver struct {
	conn *url.URL
}

func GetDriver(dbURL *url.URL) (Driver, error) {
	switch dbURL.Scheme {
	case "postgres", "postgresql":
		fmt.Println("Using PostgreSQL driver")
		return &PostgresDriver{conn: dbURL}, nil
	case "mysql":
		return nil, nil
	default:
		return nil, errors.New("unsupported scheme")
	}
}
