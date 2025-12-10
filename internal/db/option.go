package db

import (
	"github.com/duanechan/salvare/internal/config"
)

type BackupOption func([]string) []string

func ParseBackupOptions(config *config.Config, args []string) []BackupOption {
	if len(args) < 1 {
		return []BackupOption{FullBackup(config.ConnectionString())}
	}

	return []BackupOption{}
}

func FullBackup(dbURL string) BackupOption {
	return func(s []string) []string {
		return append(s, dbURL)
	}
}
