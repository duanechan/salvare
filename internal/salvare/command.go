package salvare

type handler func(*State, []string) error

type salvareCmd struct {
	name        string
	description string
	usage       string
	callback    handler
}

type commands map[string]salvareCmd

func loadCmdRegistry() commands {
	return commands{
		"backup": {
			callback: DriverMiddleware(CommandBackup),
		},
		"config init": {
			callback: CommandConfigInit,
		},
	}
}
