package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Prompt command line for username and password and returns them in that order
func UserPassPrompt() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	password := string(bytePassword)
	fmt.Print("\n\n\n")

	return strings.TrimSpace(username), strings.TrimSpace(password)
}
