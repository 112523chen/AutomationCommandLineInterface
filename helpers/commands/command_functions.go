package commands

import (
	"bufio"
	"cli/helpers/functions"
	print "cli/helpers/prints"
	"cli/helpers/prompts"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func RunCLI(home string, executable_path string) {
loop:
	for {
		print.PrintConsole()
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {

		case "1", "web":
			fmt.Println("Please enter the code error")
			codeError := functions.GetText()
			_, error := RunExecutable(home, executable_path, "search", codeError)
			if error != nil {
				fmt.Println("")
				fmt.Println(error)
				fmt.Println("Press enter to continue")
				functions.GetText()
			}
			fmt.Print("\033c")

		case "2", "job":
			fmt.Println("Please enter a job query")
			prompt := functions.GetText()
			_, err := RunExecutable(home, executable_path, "searchRole", prompt)
			if err != nil {
				fmt.Println("")
				fmt.Println(err)
				fmt.Println("Press enter to continue")
				functions.GetText()
			}
			fmt.Print("\033c")

		case "3", "recruiter":
			fmt.Println("Please enter a company name")
			prompt := functions.GetText()
			_, err := RunExecutable(home, executable_path, "searchRecruiter", prompt)
			if err != nil {
				fmt.Println("")
				fmt.Println(err)
				fmt.Println("Press enter to continue")
				functions.GetText()
			}
			fmt.Print("\033c")

		case "4", "spotify":
			RunSpotifyCleaner(home, executable_path)

		case "5", "brew":
			RunBrewCleaner()

		case "q", "quit":
			fmt.Print("\033c")
			break loop

		default:
			fmt.Println("Unknown option")
		}
	}
}

func RunExecutable(home string, executable_path string, command string, arguments string) (string, error) {
	os.Chdir(home)
	result, err := exec.Command(fmt.Sprintf("%s/%s", executable_path, command), arguments).Output()
	return string(result), err
}

func RunSpotifyCleaner(home string, executable_path string) string {
	var prompt, compress string
	var flagList []string
	os.Chdir(home)
	command := fmt.Sprintf("%s/%s", executable_path, "cleanSpotifyCache")

	fmt.Print("\033c")
	fmt.Println("Cleaning Spotify Cache")

	var question = "Do you want to dry run?"
	var options = []string{"y", "n"}
	prompt = prompts.GetResponseFromQuestion(question, options)

	question = "Do you want the result to be simplified?"
	compress = prompts.GetResponseFromQuestion(question, options)

	fmt.Println("")
	fmt.Println("Running Spotify Cleaner")
	flagList = prompts.GetSpotifyFlagList(prompt, compress)

	result, _ := exec.Command(command, flagList...).Output()
	fmt.Println(string(result))
	fmt.Println("Press y to complete the process not dry run or press enter to continue")
	prompt = functions.GetText()
	if prompt == "y" {
		fmt.Println("Press y to complete the process not dry run or press enter to continue")
		prompt = functions.GetText()
		if prompt == "y" {
			prompt = "n"
			flagList := prompts.GetBrewFlagList(prompt, compress)
			exec.Command("brew", flagList...).Output()
		}
	}

	return string(result)
}

func RunBrewCleaner() string {
	var prompt, length string
	var flagList []string
	length = "all"

	fmt.Print("\033c")
	fmt.Println("Cleaning Brew Cache")

	var question = "Do you want to run a dry run?"
	var options = []string{"y", "n"}
	prompt = prompts.GetResponseFromQuestion(question, options)

	question = "Do you want to clean all brew cache?"
	length = prompts.GetResponseFromQuestion(question, options)

	fmt.Println("")
	fmt.Println("Running Brew Cleaner")
	flagList = prompts.GetBrewFlagList(prompt, length)

	result, _ := exec.Command("brew", flagList...).Output()
	fmt.Println(string(result))
	fmt.Println("Press y to complete the process not dry run or press enter to continue")
	prompt = functions.GetText()
	if prompt == "y" {
		fmt.Println("Press y to complete the process not dry run or press enter to continue")
		prompt = functions.GetText()
		if prompt == "y" {
			prompt = "n"
			flagList := prompts.GetBrewFlagList(prompt, length)
			exec.Command("brew", flagList...).Output()
		}
	}
	fmt.Print("\033c")

	return string(result)
}

func RunExecutableOptionCommand(home string, executable_path string, arguments ...string) {
	if len(arguments) > 1 && (arguments[0] == "search" || arguments[0] == "searchRole" || arguments[0] == "searchRecruiter" || arguments[0] == "cleanSpotifyCache") {
		result, _ := RunExecutable(home, executable_path, arguments[0], arguments[1])
		fmt.Println(result)
	} else {
		print.PrintExecutableNames()
	}
}
