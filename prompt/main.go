package prompt

import (
	"bufio"
	"fmt"
	"os"
)

func RunPrompt(msg string, delim byte) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg + "(" + string(delim) + ") for next")
	return reader.ReadString(delim)
}
