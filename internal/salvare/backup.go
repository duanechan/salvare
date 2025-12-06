package salvare

func CommandBackup(s *State, args []string) error {
	return s.driver.Backup()
}
