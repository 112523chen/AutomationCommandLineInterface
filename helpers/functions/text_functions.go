package functions

import (
	"bufio"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func GetText() string {
	var text string

	text, _ = in.ReadString('\n')
	text = strings.Replace(text, "\n", " ", -1)

	return text
}
