package dbbackup

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/duanechan/salvare/internal/config"
)

type Dumper struct {
	Compressed bool
	Dir        string
	DbName     string
}

func NewDumper(config *config.Config, compressed bool) *Dumper {
	return &Dumper{
		Compressed: compressed,
		Dir:        config.BackupDirectory,
		DbName:     config.Conn.Database,
	}
}

const (
	defaultLayout string = "Backup_2006_01_02_15_04_05_%s.sql"
)

func (d *Dumper) WriteBackup(data []byte) error {
	filename := time.Now().Format(fmt.Sprintf(defaultLayout, d.DbName))
	path := filepath.Join(d.Dir, filename)

	if err := os.MkdirAll(d.Dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
