package print

import "fmt"

func PrintConsole() {
	fmt.Println("")
	fmt.Println("Please select an option:")
	fmt.Println("1. Search for Solution to Code Error")
	fmt.Println("2. Search for Jobs")
	fmt.Println("3. Search for Recruiters' LinkedIn")
	fmt.Println("4. Clean Spotify Cache")
	fmt.Println("5. Clean Brew Cache")
	fmt.Println("q. Exit")
	fmt.Println("")
}

func PrintExecutableNames() {
	fmt.Println("Executables:")
	fmt.Println("  search\t\t\tSearch for code solution")
	fmt.Println("  searchRole\t\t\tSearch for job")
	fmt.Println("  searchRecruiter\t\tSearch for recruiters' LinkedIn")
	fmt.Println("  cleanSpotifyCache\t\tClean Spotify Cache")
	fmt.Println("")

	fmt.Println("Usage: cli -e [executable] [flag]")
	fmt.Println("Flags:")
	fmt.Println("  -h, --help\t\t\tShow help message for executable")
	fmt.Println("  -v, --version\t\t\tShow executable's version number")
}

func PrintHelp() {
	fmt.Println("Usage: cli [flag] [options]")
	fmt.Println("Search for code solution, job postings or recruiter on LinkedIn from the command line")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -h, --help\t\t\tShow this help message and exit")
	fmt.Println("  -v, --version\t\t\tShow program's version number and exit")
	fmt.Println("  -l, --list\t\t\tShow all commands available")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -e, --executable\t\tSpecifies a certain executable")
}
