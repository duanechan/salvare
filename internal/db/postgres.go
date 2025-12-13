package db

import (
	"fmt"
	"os"
	"os/exec"
)

type PostgresDriver baseDriver

func (d PostgresDriver) Backup() ([]byte, error) {
	tmp, _ := os.CreateTemp("", "pgpass-*")
	defer os.Remove(tmp.Name())

	line := fmt.Sprintf("%s:%s:%s:%s:%s\n",
		d.Conn.Hostname,
		d.Conn.Port,
		d.Conn.Database,
		d.Conn.Username,
		d.Conn.Password,
	)

	tmp.WriteString(line)
	tmp.Chmod(0600)
	tmp.Close()

	cmd := exec.Command("pg_dump",
		"-h", d.Conn.Hostname,
		"-p", d.Conn.Port,
		"-U", d.Conn.Username,
		d.Conn.Database,
	)

	cmd.Env = append(os.Environ(),
		"PGPASSFILE="+tmp.Name(),
	)

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
