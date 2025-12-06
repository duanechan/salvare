package salvare

import "fmt"

type salvareCmd struct {
	name        string
	description string
	usage       string
	callback    func(string, []string) error
}

type commands map[string]salvareCmd

func loadCmdRegistry() commands {
	return commands{
		"backup": {
			callback: func(s1 string, s2 []string) error {
				fmt.Println("Hello!")
				return nil
			},
		},
	}
}
