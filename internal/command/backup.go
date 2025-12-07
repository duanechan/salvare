package command

import "fmt"

func CommandBackup(state *State, args []string) error {
	data, err := state.driver.Backup()
	if err != nil {
		return err
	}

	fmt.Println("Backup created!")
	state.dumpFile(data)

	return nil
}
