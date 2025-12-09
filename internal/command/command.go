package command

import "github.com/duanechan/salvare/internal/models"

type handler func(*State, []string) (*models.Metrics, error)

type salvareCmd struct {
	name        string
	description string
	usage       string
	callback    handler
}

type commands map[string]salvareCmd

type registry map[string]commands

func loadCmdRegistry() registry {
	return registry{
		"backup": {
			"": {
				name:        "backup",
				description: "Creates a dump (backup) file in the configuration's specified backup directory.",
				usage:       "salvare backup",
				callback:    DriverMiddleware(CommandBackup),
			},
		},
		"config": {
			"init": {
				name:        "config init",
				description: "Creates an empty Salvare configuration file.",
				usage:       "salvare config init",
				callback:    CommandConfig,
			},
		},
	}
}
