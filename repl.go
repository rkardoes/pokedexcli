package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// make lowercase
	// remove leading/trailing whitespace
	cleanText := strings.TrimSpace(strings.ToLower(text))
	return strings.Split(cleanText, " ")
}
