package db

import (
	"errors"
	"net/url"
)

type Driver interface {
	Backup() ([]byte, error)
	Restore() error
	Compress() error
}

type driverConstructor func(*url.URL) Driver

type drivers map[string]driverConstructor

func loadDriverRegistry() drivers {
	return drivers{
		"postgres":   func(u *url.URL) Driver { return &PostgresDriver{Conn: u} },
		"postgresql": func(u *url.URL) Driver { return &PostgresDriver{Conn: u} },
		"mysql":      nil,
	}
}

type baseDriver struct {
	Conn *url.URL
}

func GetDriver(dbURL *url.URL) (Driver, error) {
	constructor, exists := loadDriverRegistry()[dbURL.Scheme]
	if !exists {
		return nil, errors.New("unsupported driver")
	}

	return constructor(dbURL), nil
}
