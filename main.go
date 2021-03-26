package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	fmt.Println("Go Shell starting")

	// read input from keyboard
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(createPrompt())
		input, err := reader.ReadString('\n')
		if err != nil {
			// print to stderr
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func createPrompt() string {
	hostname, _ := os.Hostname()
	cwd, _ := os.Getwd()
	currentUser, _ := user.Current()
	username := currentUser.Username

	return fmt.Sprintf("%s@%s[%s]> ", username, hostname, cwd)
}
func execInput(input string) error {
	// remove \n character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
			return errors.New("path required")
		}
		// change directory to path provided
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// execute the command
	// arg[0] is ls
	// arg[1:] is -l -t -r
	command := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	// Execute the command and return the error.
	return command.Run()
}
