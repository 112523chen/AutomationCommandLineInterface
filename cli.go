package main

import (
	"cli/helpers/commands"
	"cli/helpers/flags"
	"fmt"
	"os"
)

// TODO: Refactor command_functions.go line 104 to 114
// TODO: Add functionality to add or remove websites from search and searchRole
// TODO: Add restart functionally

func main() {
	home, _ := os.UserHomeDir()
	executable_path := "./Execs"
	var flag string

	if len(os.Args[1:]) > 0 {
		if os.Args[1] == "-e" || os.Args[1] == "--executable" {
			flag = os.Args[1]
			flags.PerformFlags(&flag, os.Args[2:]...)
		} else {
			flag = os.Args[1]
			flags.PerformFlags(&flag)
		}
	} else {
		fmt.Print("\033c")
		fmt.Println("Welcome to the CLI")
		commands.RunCLI(home, executable_path)
	}

}
