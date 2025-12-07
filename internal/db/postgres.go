package db

import (
	"os/exec"
)

type PostgresDriver baseDriver

func (d PostgresDriver) Backup() ([]byte, error) {
	cmd := exec.Command("pg_dump", d.Conn.String())

	dump, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return dump, nil
}

func (d PostgresDriver) Restore() error {

	return nil
}

func (d PostgresDriver) Compress() error {

	return nil
}
