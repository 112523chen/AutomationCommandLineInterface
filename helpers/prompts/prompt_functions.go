package prompts

import (
	"cli/helpers/functions"
	"fmt"
	"strings"
)

func GetResponseFromQuestion(question string, responseList []string) string {
	var result string

	fmt.Printf("%s (%s) ", question, strings.Join(responseList, "/"))
	result = strings.TrimSpace(functions.GetText())

	for result != "y" && result != "n" {
		fmt.Println("")
		fmt.Printf("Please enter %s\n", strings.Join(responseList, " or "))
		fmt.Printf("%s (%s) ", question, strings.Join(responseList, "/"))
		result = strings.TrimSpace(functions.GetText())
	}

	return result
}

func GetSpotifyFlagList(prompt1 string, prompt2 string) []string {
	var flagList []string
	if prompt1 == "y" && prompt2 == "y" {
		flagList = []string{"-d"}
	} else if prompt1 == "y" && prompt2 == "n" {
		flagList = []string{"-d", "-f"}
	} else if prompt1 == "n" && prompt2 == "n" {
		flagList = []string{"-f"}
	} else {
		flagList = []string{}
	}

	return flagList
}

func GetBrewFlagList(prompt1 string, prompt2 string) []string {
	var flagList []string
	if prompt1 == "y" && prompt2 == "y" {
		flagList = []string{"cleanup", "--dry-run", "--prune=all"}
	} else if prompt1 == "y" && prompt2 == "n" {
		flagList = []string{"cleanup", "--dry-run"}
	} else if prompt1 == "n" && prompt2 == "y" {
		flagList = []string{"cleanup", "--prune=all"}
	} else {
		flagList = []string{"cleanup"}
	}

	return flagList
}
