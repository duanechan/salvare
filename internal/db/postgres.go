package db

import (
	"fmt"
	"os/exec"
)

type PostgresDriver baseDriver

func (d PostgresDriver) Backup() error {
	fmt.Println(d.conn.String())
	cmd := exec.Command("pg_dump", d.conn.String())

	dump, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Println(string(dump))

	return nil
}

func (d PostgresDriver) Restore() error {

	return nil
}

func (d PostgresDriver) Compress() error {

	return nil
}
