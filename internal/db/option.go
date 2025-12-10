package db

type BackupOption func([]string)

func FullBackup(dbURL string) BackupOption {
	return func(s []string) {
		s = append(s, dbURL)
	}
}
