package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

type cliCommand struct {
		name		string
		description	string
		callback	func(*config) error
	}


type config struct {
	commands	map[string]cliCommand;
	mapPage		int;
}

func (c *config) changeMapPage(d int) {
	c.mapPage += d
}

func cleanInput(text string) []string {
	// make lowercase
	// remove leading/trailing whitespace
	cleanText := strings.TrimSpace(strings.ToLower(text))
	return strings.Split(cleanText, " ")
}

func startRepl(cfg *config) {
	cmds := cfg.commands
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		cmd_txt := input[0]
		cmd, ok := cmds[cmd_txt]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback(cfg)
	}
}
