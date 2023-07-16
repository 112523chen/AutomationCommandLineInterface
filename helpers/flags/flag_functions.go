package flags

import (
	"cli/helpers/commands"
	print "cli/helpers/prints"
	"fmt"
	"os"
)

func PerformFlags(flag *string, arguments ...string) {
	home, _ := os.UserHomeDir()
	executable_path := "./Execs"

	switch *flag {
	case "-h", "--help":
		print.PrintHelp()
	case "-v", "--version":
		fmt.Println("Version 1.0.0")
	case "-l", "--list":
		print.PrintExecutableNames()
	case "-e", "--executable":
		commands.RunExecutableOptionCommand(home, executable_path, arguments...)
	default:
		fmt.Println("Unknown flag")
		fmt.Println("Try -h or --help for help")
	}
}
