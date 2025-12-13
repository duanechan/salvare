package db

import (
	"errors"

	"github.com/duanechan/salvare/internal/config"
)

type BackupFunc func() []string
type RestoreFunc func() error
type CompressFunc func() error

type Driver interface {
	Backup() ([]byte, error)
	Restore() error
	Compress() error
}

type driverConstructor func(conn *config.Conn) Driver

type drivers map[string]driverConstructor

func loadDriverRegistry() drivers {
	return drivers{
		"postgres":   func(conn *config.Conn) Driver { return &PostgresDriver{Conn: conn} },
		"postgresql": func(conn *config.Conn) Driver { return &PostgresDriver{Conn: conn} },
		"mysql":      nil,
	}
}

type baseDriver struct {
	Conn *config.Conn
}

func GetDriver(conn *config.Conn) (Driver, error) {
	constructor, exists := loadDriverRegistry()[conn.Scheme]
	if !exists {
		return nil, errors.New("unsupported driver")
	}

	return constructor(conn), nil
}
