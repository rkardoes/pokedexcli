package main

import (
	"fmt"
	"slices"
)

func commandHelp(cfg *config) error {
	cmds := cfg.commands

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	keys := make([]string, 0, len(cmds))
	for k := range cmds {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for i := 0; i < len(cmds); i++ {
		c := cmds[keys[i]]
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
