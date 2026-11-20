package main


func main() {
	cfg := config{
        commands:       map[string]cliCommand {
                "exit": {
                        name:           "exit",
                        description:    "Exit the Pokedex",
                        callback:       commandExit,
                	},
                "help": {
                        name:           "help",
                        description:    "Displays a help message",
                        callback:       commandHelp,
                	},
		"map": {
			name:		"map",
			description:	"Get location areas",
			callback:	commandMap,
			},
		"mapb": {
			name:		"mapb",
			description:	"Get prior page location areas",
			callback:	commandMapb,
			},
        	},
	page:		0,

	}

	startRepl(&cfg)
}
