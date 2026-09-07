package utils

import (
	"bufio"
	"os"
	"strings"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
